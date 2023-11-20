package server

import (
	_ "main/docs"
	"main/server/handler"
	gomail "main/server/services/alert_service/Gomail"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.engine.POST("/send-email", gomail.SendEmailOtpService)
	server.engine.POST("/send-sms", handler.TwilioServiceHnadler)

	server.engine.GET("/ping", Pong)

}
//For server Testing(acknowledgement)
func Pong(ctx *gin.Context) {

	msg := "Pong"
	ctx.Writer.Write([]byte(msg))
}
