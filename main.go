package main

import "mutant-project/services"

func main() {
	// create a service struct.
	app := &services.Service{}

	// initialize api.
	app.Initialize()

	// run the server with specific port.
	app.Run(":8080")
}
