package handler

import (
	"main/server/services/payments/paypal"

	"github.com/gin-gonic/gin"
)

func GeneratePayPalAuthToken(ctx *gin.Context) {

	paypal.GetAccessToken(ctx)
}

func CreateWebhook(ctx *gin.Context) {

	paypal.CreateWebhook(ctx)
}

func HandleWebHookNotification(ctx *gin.Context) {
	paypal.HandleWebhookNotification(ctx)
}
