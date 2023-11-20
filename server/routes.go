package server

import (
	_ "main/docs"

	"main/server/gateway"
	alertservices "main/server/handler/alert_services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	//Allowing CORS
	server.engine.Use(gateway.CORSMiddleware())

	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.engine.POST("/send-sms", alertservices.AwsTextMessaging)
}
