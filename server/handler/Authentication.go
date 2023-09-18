package handler

import (
	"main/server/request"
	"main/server/response"
	authentication "main/server/services/Authentication"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)






func SignupHandler(ctx *gin.Context){

	
	var input request.AuthRequest
	utils.RequestDecoding(ctx, &input)


	//check if the newPassword and password are the same
	if input.NewPassword!=input.Password{
		
		response.ShowResponse("Password are not same", 400, "Bad Request", "", ctx)
		return
	}
	
	
	//call the service with the inputRequest credentials 
	authentication.SignupService(ctx, &input)


}

func LoginHandler(ctx *gin.Context){

	var input request.AuthRequest
	utils.RequestDecoding(ctx, &input)

	authentication.LoginService(ctx, &input)



}