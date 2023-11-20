package alertservices

import (
	"main/server/request"
	"main/server/response"

	awssns "main/server/services/alert_service/aws_sns"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func AwsTextMessaging(ctx *gin.Context) {
	var reqBody request.AwsTextMessagingRequest

	err := utils.RequestDecoding(ctx, &reqBody)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	awssns.SendSMS(ctx, reqBody)
}
