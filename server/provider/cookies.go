package provider

import (
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// Set cookie handler
func SetCookie(context *gin.Context, tokenString string) {

	context.SetCookie(
		"cookie",
		tokenString,
		7200,
		"/",
		"localhost",
		false,
		true,
	)

	response.ShowResponse(
		"Success",
		utils.HTTP_OK,
		"Cookies saved successfully",
		"",
		context,
	)
}

// Delete cookie handler
func DeleteCookie(context *gin.Context) {
	context.SetCookie(
		"cookie",
		"",
		-1,
		"",
		"",
		false,
		false,
	)

	response.ShowResponse(
		"Success",
		utils.HTTP_OK,
		"Cookie deleted successfully",
		"",
		context,
	)
}
