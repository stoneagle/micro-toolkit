package main

import (
	"toolkit/backend/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	bootstrap.Boot(app)
	app.Run(":8079")
}
