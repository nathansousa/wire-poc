package controllers

import (
	"fmt"
	"log"
)

type IMessages interface {
	PrintGreetingMessage()
	PrintFarewellMessage()
}

type GreetingMessage string
type FarewellMessage string

type Messages struct {
	Greeting GreetingMessage
	Farewell FarewellMessage
}

func NewMessageController(greetingMessage GreetingMessage, farewellMessage FarewellMessage) IMessages {
	return &Messages{
		Greeting: greetingMessage,
		Farewell: farewellMessage,
	}
}

func (m *Messages) PrintGreetingMessage() {
	if m.Greeting == "" {
		log.Panic("empty greeting message")
	}

	fmt.Print(m.Greeting)
}

func (m *Messages) PrintFarewellMessage() {
	if m.Greeting == "" {
		log.Panic("empty farewell message")
	}

	fmt.Print(m.Farewell)
}
