package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "app.tmpl", gin.H{
		"version": version,
	})
}

func healthCheckHandler(c *gin.Context) {
	fmt.Fprint(c.Writer, "ok")
}

func startServer(port int) error {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// statics
	router.Static("/css", "./static/css")
	router.Static("/img", "./static/img")

	// templates
	router.LoadHTMLGlob("templates/*")

	// handlers
	router.GET("/", indexHandler)
	router.GET("/_ah/health", healthCheckHandler)

	// server
	httpserver := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handlers.CombinedLoggingHandler(os.Stdout, router),
		ReadTimeout:  40 * time.Second,
		WriteTimeout: 40 * time.Second,
	}

	// run server
	logger.Fatal(httpserver.ListenAndServe())
	return nil
}
