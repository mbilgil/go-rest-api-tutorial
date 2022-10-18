package main

import "go-rest-api-tutorial/src/app"

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":9000")
}
