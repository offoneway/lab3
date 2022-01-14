package main

import (
	"github.com/gin-gonic/gin"
	"github.com/offoneway/SE2/entity"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()

	// Run the server
	r.Run()
}
