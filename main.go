package main

import (
	"task-company/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/twosum", controllers.Towsum)
	r.POST("/threesum", controllers.ThreeSum)
	r.POST("/findsubstring", controllers.FindSubstring)

	r.Run()
}