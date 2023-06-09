package router

import (
	"hello-run/controller"
	"hello-run/middleware/googleauth"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(googleauth.FakeAuth())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "ok",
		})
	})

	router.GET("/user/", controller.GetUserInfo)

	router.GET("/record/list/", controller.ListRecordByRoomId)
	router.POST("/record/create/", controller.CreateRecord)
	router.POST("/record/update/", controller.UpdateRecord)
	router.POST("/record/delete/", controller.DeleteRecord)

	router.GET("/room/", controller.GetRoomInfo)
	router.POST("/room/update/", controller.UpdateRoom)
	router.POST("/room/create/", controller.CreateRoom)
	router.POST("/room/delete/", controller.DeleteRoom)

	router.POST("/household/create/", controller.CreateHousehold)
	router.POST("/household/update/", controller.UpdateHousehold)
	router.POST("/household/delete/", controller.DeleteHousehold)

	return router
}
