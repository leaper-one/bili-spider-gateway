package main

import (
	"ginForBH/handle"
	"ginForBH/initialize"
	"ginForBH/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	var err error
	initialize.InitMysql()
	r := gin.Default()

	r.POST("/bililive/api/getrank", func(c *gin.Context) {
		var get model.Getrank
		var back model.GetrankBack

		if err := c.ShouldBindJSON(&get); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if back, err = handle.PostGetrank(get); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, back)
	})

	r.POST("/bililive/api/hasnote", func(c *gin.Context) {
		var get model.Hasnote

		if err := c.ShouldBindJSON(&get); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err = handle.PostHasnote(get); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg": "ok", "timestamp": time.Now().Unix()})
	})

	r.POST("/bililive/api/denote", func(c *gin.Context) {
		var get model.Hasnote

		if err := c.ShouldBindJSON(&get); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err = handle.PostDenote(get); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg": "ok", "timestamp": time.Now().Unix()})
	})

	r.Run(":8000")
}
