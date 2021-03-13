package main

import (
	"log"
	"wirepoc/config/providers"
	"wirepoc/internal/controllers"
)

func main() {
	greetingMessage := controllers.GreetingMessage("Hello")
	farewellMessage := controllers.FarewellMessage("Goodbye")

	msgController, err := providers.InitializeMessages(greetingMessage, farewellMessage)
	if err != nil {
		log.Panic(err)
	}

	msgController.PrintGreetingMessage()
	print("\n")
	msgController.PrintFarewellMessage()
}
