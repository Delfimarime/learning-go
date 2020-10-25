package impl

import (
	"../engine"
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
)

type XmlQuiz struct {
	XMLName   xml.Name     `xml:"quiz"`
	Questions XmlQuestions `xml:"questions"`
}

type XmlQuestions struct {
	XMLName xml.Name      `xml:"questions"`
	Content []XmlQuestion `xml:"question"`
}

type XmlQuestion struct {
	XMLName     xml.Name    `xml:"question"`
	Text        string      `xml:"text,attr"`
	Hint        string      `xml:"hint,attr"`
	Answer      string      `xml:"answer,attr"`
	Options     []XmlOption `xml:"option"`
	OptionIndex string      `xml:"option-index,attr"`
}

type XmlOption struct {
	XMLName xml.Name `xml:"option"`
	Text    string   `xml:"text,attr"`
}

type XmlRepository struct {
	Path string
}

func (repository XmlRepository) Load() []engine.Question {

	// Open our xmlFile
	xmlFile, err := os.Open(repository.Path)

	// if we os.Open returns an error then handle it
	if err != nil {
		panic(err)
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	content, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var quiz XmlQuiz

	// we unmarshal our byteArray which contains our
	xml.Unmarshal(content, &quiz)

	if quiz.Questions.Content == nil || len(quiz.Questions.Content) == 0 {
		return make([]engine.Question, 0)
	}

	questions := make([]engine.Question, 0)

	for i := 0; i < len(quiz.Questions.Content); i++ {
		question := quiz.Questions.Content[i]

		if len(question.Text) > 0 {

			onlyAnswer := len(question.Answer) > 0 && (question.Options == nil || len(question.Options) == 0) && len(question.OptionIndex) == 0
			onlyOptions := len(question.Answer) == 0 && question.Options != nil && len(question.Options) > 0 && len(question.OptionIndex) > 0

			if onlyAnswer || onlyOptions {
				isValue := false
				options := make([]string, 0)
				answer := ""

				if onlyOptions {

					if len(question.Options) == 1 {
						question.Answer = question.Options[0].Text
						isValue = true
					} else {
						for j := 0; j < len(question.Options); j++ {
							options = append(options, question.Options[j].Text)
						}

						i, err := strconv.Atoi(question.OptionIndex)

						if err == nil && i >= 0 && i < len(question.Options) {
							answer = question.Options[i].Text
							isValue = true
						}
					}
				} else {
					answer = question.Answer
					isValue = true
				}

				if isValue {
					value := engine.Question{Text: question.Text, Hint: question.Hint, Answer: answer, Options: options}
					questions = append(questions, value)
				}

			}
		}

	}

	return questions
}
