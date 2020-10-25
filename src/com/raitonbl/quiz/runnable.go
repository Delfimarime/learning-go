package main

import (
	"./impl"
	"fmt"
	"strconv"
)

func main() {

	client := impl.TimelessClient{}
	client.SetRepository(impl.XmlRepository{Path: "/Users/delfimarime/Documents/git/github/learning-go/quiz.xml"})
	client.AfterPropertiesSet()

	for ok := true; ok; ok = client.HasNext() {
		options := client.GetOptions()

		fmt.Println("Question:", client.GetQuestion()+"?")

		if len(options) == 0 {
			doAnswer(&client)
		} else {
			doOptions(&client, options)
		}

	}

	fmt.Printf("Correct answers: %d of %d",client.GetScore(), client.GetMaximumScore())

}

func doAnswer(client *impl.TimelessClient) {
	var input string
	fmt.Print("Answer:")
	fmt.Scan(&input)
	fmt.Println("Is answer correct:", client.Answer(input))
}

func doOptions(client *impl.TimelessClient, options []string) {
	fmt.Println("Available Options")

	for i, p := range options {
		fmt.Println(strconv.Itoa(i+1) + "." + p)
	}

	var input int
	fmt.Print("Answer:")
	fmt.Scan(&input)

	if input > 0 && input <= len(options) {
		fmt.Println("Is answer correct:", client.Answer(options[input-1]))
	}

}
