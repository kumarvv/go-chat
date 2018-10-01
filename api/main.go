package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	fmt.Println("Starting api ...")

	http := gin.Default()

	http.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"PUT, GET, POST, DELETE, OPTIONS"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	http.GET("/", WelcomeHandler)
	http.GET("/echo", EchoHandler)
	http.GET("/echo/:name", EchoHandler)

	http.Run()
}

func WelcomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "success",
		"msg":    "Hello, World!",
	})
}

func EchoHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "success",
		"msg":    "Helloooo, " + c.Param("name"),
	})
}
