package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Say(c *gin.Context) {
	fmt.Println("heelo ,world")
	c.JSON(200, gin.H{
		"msg": "hello world",
	})

}
