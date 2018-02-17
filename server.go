package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/*any", func(c *gin.Context) {
		ip := c.Request.RemoteAddr
		fmt.Printf("IP: %s\n", ip)
		c.String(200, "IP: %s", ip)
	})
	r.Run(":9000")
}