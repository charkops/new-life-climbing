package main

import (
	"climbing/ginengine"
	"climbing/middleware"
	"climbing/routers"
	"fmt"
)

func main() {
	// What do we wanna do here ?
	// Initialize DB
	// Initialize routers/routes
	// Serve

	// Initialize gin.Engine
	// NOTE (@Charkops): Maybe use sync.Once to only do this once ?
	ginengine.Setup()
	middleware.Setup()
	routers.Setup()
	fmt.Println("New gin engine created")

	ginengine.Engine.Run()
}
