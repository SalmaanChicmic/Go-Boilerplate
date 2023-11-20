package alertservices

import (
	"main/server/request"
	awssns "main/server/services/alert_service/aws_sns"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func AwsTextMessaging(ctx *gin.Context) {
	var reqBody request.AwsTextMessagingRequest

	utils.RequestDecoding(ctx, &reqBody)

	awssns.SendSMS(ctx, reqBody)
}
