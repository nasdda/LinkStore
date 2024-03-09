package main

import (
	routes "github.com/nasdda/linkstore/routes"
)

func main() {
	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
