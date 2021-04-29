package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"hello":"world"})
	})
	r.POST("/user", func(c *gin.Context){
		c.Status(http.StatusNotImplemented)
	})
	log.Println("Server listening")
	err := r.Run()

	if err != nil {
		log.Fatal(err)
	}
}
