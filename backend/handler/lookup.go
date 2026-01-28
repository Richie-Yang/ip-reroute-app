package handler

import (
	"ip-reroute-app/service"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.POST("/lookup", lookup)
}

func lookup(c *gin.Context) {
	clientIP := c.ClientIP()

	ipInfo, err := service.GetIPInfo(clientIP)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"ip":       ipInfo.IP,
		"city":     ipInfo.City,
		"region":   ipInfo.Region,
		"country":  ipInfo.Country,
		"location": ipInfo.Location,
	})
}

func fakeLookup(c *gin.Context) {
	c.JSON(200, gin.H{
		"ip":       "127.0.0.1",
		"city":     "Taipei",
		"region":   "Taiwan",
		"country":  "TW",
		"location": "Taiwan",
	})
}
