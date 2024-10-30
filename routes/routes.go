package routes

import (
	"fall-cloud/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest()  {
	app := gin.Default()

	app.GET("/",controllers.Index)
	app.POST("/register",controllers.CreateUser)
	app.POST("/login",controllers.Login)

	//Services
	app.GET("/services/list",controllers.ListOfAvailableServices)
	app.GET("/services/new/:id",controllers.CreateNewService)

	app.Run(":3000")
}