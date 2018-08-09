package main

import (
	"io"
	"os"

	api "github.com/Todai88/faceIt/microservice/api"
	"github.com/gin-gonic/gin"
)

func main() {
	setupLogging()
	startServer()
}

/*
 */
func setupLogging() {
	// TODO: Check environment variables to decide where to log.
	// No need for color overhead when writing to console
	logFile := logFileLocation()

	if len(logFile) > 0 {
		gin.DisableConsoleColor()
		f, _ := os.Create("log/" + logFile)
		gin.DefaultWriter = io.MultiWriter(f)
	}
}

// Used to start the server and serve the endpoints
func startServer() {
	// TODO: start server, subscribe to endpoints.
	router := gin.Default()
	v1 := router.Group("/api/v1/users")
	{
		v1.GET("/", api.GetUsers)
		// v1.POST("/", createTodo)
		// v1.GET("/:id", fetchSingleTodo)
		// v1.PUT("/:id", updateTodo)
		// v1.DELETE("/:id", deleteTodo)
	}

	router.Run(":5050")
}

func logFileLocation() string {
	logFileLocation := os.Getenv("logFile")
	if len(logFileLocation) == 0 {
		logFileLocation = ""
	}
	return logFileLocation
}
