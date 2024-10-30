package controllers

import (
	"fall-cloud/database"
	"fall-cloud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context)  {
	ctx.JSON(200,gin.H{
		"message":"Hellow Docker",
	})
}

func CreateUser(ctx *gin.Context){
	var User models.User 
	if err := ctx.ShouldBindJSON(&User); err != nil{
		ctx.JSON(http.StatusNotAcceptable,err)
	}

	database.DB.Create(&User)

	ctx.JSON(http.StatusOK,gin.H{
		"status":"Ok",
	})
}

func Login(ctx *gin.Context){
	var User models.User 
	if err := ctx.ShouldBindJSON(&User); err != nil{
		ctx.JSON(http.StatusNotAcceptable,err)
	}

	database.DB.Where(&models.User{Name: User.Name,Password: User.Password}).First(&User)

	if User.ID == 0{
		ctx.JSON(http.StatusUnauthorized,gin.H{
			"NÃO AUTORIZADO!":"Senha Incorreta!",
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"SUCESSO!":"AUTORIZADO!",
	})

}