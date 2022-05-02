package main

import (
	"ginForBH/handle"
	"ginForBH/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("bililive", func(c *gin.Context) {
		var get model.PostFormGet
		if err := c.ShouldBindJSON(&get); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var err error
		if _, err = handle.Handle(get); err != nil {
			log.Printf("err: %v", err)
		}
		c.JSON(http.StatusOK, gin.H{})
	})

	r.Run(":8000")
}
