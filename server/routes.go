package server

import (
	_ "main/docs"
	"main/server/gateway"
	"main/server/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

	//Allowing CORS
	server.engine.Use(gateway.CORSMiddleware())

	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Paypal routes
	server.engine.POST("/generate-token", handler.GeneratePayPalAuthToken)
	server.engine.POST("/create-webhook", handler.CreateWebhook)
	server.engine.POST("/capture", handler.HandleWebHookNotification)

}
