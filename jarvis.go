package main

import "github.com/magrathealabs/jarvis/infrastructure/controllers"

func main() {
	err := controllers.NewAppController().Engine.Run()
	panic(err)
}
