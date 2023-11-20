package server

import (
	_ "main/docs"
<<<<<<< Updated upstream
	"main/server/gateway"
=======
	alertservices "main/server/handler/alert_services"
>>>>>>> Stashed changes

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {

<<<<<<< Updated upstream
	//Allowing CORS
	server.engine.Use(gateway.CORSMiddleware())

	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

=======
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.engine.POST("/send-sms", alertservices.AwsTextMessaging)

>>>>>>> Stashed changes
}
