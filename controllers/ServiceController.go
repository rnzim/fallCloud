package controllers

import (
	
	"fall-cloud/docker"
	
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ListOfAvailableServices(ctx *gin.Context){
	var list = []string{"Firefox","Minecraft"}
	ctx.JSON(http.StatusOK,gin.H{
		"list": list,
	})
}

func ListAllServices()  {
	
}

func CreateNewService(ctx *gin.Context)  {
	//var UserID
	//var service models.Service
	id := ctx.Params.ByName("id")

	id_out,err := strconv.Atoi(id)
	if err != nil{
		panic(err.Error())
	}
	result,idContainer := docker.CreateContainer(id_out)
	clean := strings.ReplaceAll(idContainer,"\n","-Port")
	idContainer = clean
	
	//database.DB.Create(&service)
	ctx.JSON(http.StatusOK,gin.H{
			"id":idContainer,
			"status":"OK",
			"addr":"localhost",
			"port":result,
		})
		return
	
	//ctx.JSON(http.StatusNotFound,gin.H{
	//	"serviço inexistente":"escolha um serviço valido",
	//})


}