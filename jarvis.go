package main

import (
	"fmt"

	"github.com/krakenlab/ternary"
	"github.com/magrathealabs/jarvis/infrastructure/controllers"
	"github.com/magrathealabs/jarvis/infrastructure/env"
)

func main() {
	err := controllers.NewAppController().Server.Run(fmt.Sprintf("%s:%s", env.Web().Interface(), env.Web().Port()))

	ternary.Func(err == nil, func() {}, func() { panic(err) })()
}
