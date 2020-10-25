package main

import (
	"./impl"
	"fmt"
)

func main() {

	client := impl.TimelessClient{}
	client.SetRepository(impl.XmlRepository{Path: "/Users/delfimarime/Documents/git/github/learning-go/quiz.xml"})
	client.AfterPropertiesSet()
	fmt.Println("Question", client.GetQuestion())
	fmt.Println("Answer", client.Answer("4"))
	fmt.Println("Question", client.GetQuestion())
	fmt.Println("Score",client.GetScore())

}
