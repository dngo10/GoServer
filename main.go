package main

import (
	ctrl "goapp2/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	var controller ctrl.DataController
	controller.Init()

	//result := controller.Db.Model(&data.Data{}).Create(&data.Data{Uuid: "abc", Data: "hahaha"})
	//controller.Db.Model(&data.Data{}).Create(&data.Data{Uuid: "abc1", Data: "hahaha1"})
	//if result.Error != nil{
	//	fmt.Printf("Error: %s\n", result.Error.Error())
	//	return
	//}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",		
		})
	})

	r.GET("/api/data/:uuid", controller.GetData)
	r.POST("/api/data", controller.AddData)
	r.POST("/api/data/edit", controller.EditData)
	r.DELETE("/api/data/:uuid",controller.DeleteData)

	r.Run("localhost:8080")
}
