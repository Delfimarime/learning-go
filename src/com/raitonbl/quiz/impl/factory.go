package impl

import (
	"../engine"
	"errors"
)

type TimelessClient struct {
	score      int
	index      int
	init       bool
	quiz       []engine.Question
	repository engine.Repository
}

func (client *TimelessClient) getHint() string {

	if !client.init {
		errors.New("client requires initialization")
	}

	if !(client.index > 0 && client.index < len(client.quiz)) {
		errors.New("no hint available since no question is available")
	}

	return client.quiz[client.index].Hint
}

func (client *TimelessClient) GetQuestion() string {

	if !client.init {
		errors.New("client requires initialization")
	}

	if !(client.index > 0 && client.index < len(client.quiz)) {
		errors.New("no question available")
	}

	return client.quiz[client.index].Text
}

func (client *TimelessClient) GetOptions() []string {

	if !client.init {
		errors.New("client requires initialization")
	}

	if !(client.index > 0 && client.index < len(client.quiz)) {
		errors.New("no options available since no question is available")
	}

	return client.quiz[client.index].Options
}

func (client *TimelessClient) Answer(answer string) bool {

	if !client.init {
		errors.New("client requires initialization")
	}

	if !(client.index > 0 && client.index < len(client.quiz)) {
		errors.New("no question is available")
	}

	index := client.index

	success := client.quiz[index].Answer == answer

	if success {
		client.score = client.score + 1
	}

	client.index = index + 1

	return success
}

func (client *TimelessClient) GetScore() int {

	if !client.init {
		errors.New("client requires initialization")
	}

	return client.score
}

func (client *TimelessClient) HasNext() bool {

	if !client.init {
		errors.New("client requires initialization")
	}

	return client.index <= len(client.quiz)-1
}

func (client *TimelessClient) IsLast() bool {

	if !client.init {
		errors.New("client requires initialization")
	}

	return client.index == len(client.quiz)-1
}

func (client *TimelessClient) IsFirst() bool {

	if !client.init {
		errors.New("client requires initialization")
	}

	return client.index == 0
}

func (client *TimelessClient) AfterPropertiesSet() {
	if !client.init {
		client.quiz = client.repository.Load()
		client.init = true
	}
}

func (client *TimelessClient) SetRepository(repository engine.Repository) {
	if repository != nil {
		client.repository = repository
		client.index = 0
		client.init = false
	}
}
