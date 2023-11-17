package server

import (
	_ "main/docs"
	gomail "main/server/services/alert_service/Gomail"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.engine.POST("/send-email", gomail.SendEmailOtpService)
}
