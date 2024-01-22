package handler

import (
	"main/server/request"
	"main/server/response"
	authentication "main/server/services/Authentication"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func SignupHandler(ctx *gin.Context) {
	var input request.AuthRequest
	err := utils.RequestDecoding(ctx, &input)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//call the service with the inputRequest credentials
	authentication.Signup(ctx, &input)

}

func LoginHandler(ctx *gin.Context) {

	var input request.AuthRequest
	err := utils.RequestDecoding(ctx, &input)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	authentication.Login(ctx, &input)

}
