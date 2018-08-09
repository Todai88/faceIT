package main

import (
	"io"
	"os"

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
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":5050")
}

func logFileLocation() string {
	logFileLocation := os.Getenv("logFile")
	if len(logFileLocation) == 0 {
		logFileLocation = ""
	}
	return logFileLocation
}
