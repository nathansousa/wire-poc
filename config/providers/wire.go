//+build wireinject

package providers

import (
	"github.com/google/wire"
	"wirepoc/internal/controllers"
)

func InitializeMessages(greetingMessage controllers.GreetingMessage,
	farewellMessage controllers.FarewellMessage) (controllers.IMessages, error) {
	wire.Build(controllers.NewMessageController)
	return &controllers.Messages{}, nil
}
