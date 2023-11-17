package utils

import (
	"encoding/json"
	"io/ioutil"
	"main/server/response"

	"github.com/gin-gonic/gin"
)

// func RequestDecoding(ctx *gin.Context, data interface{}) {

// 	reqBody, err := ioutil.ReadAll(ctx.Request.Body)
// 	if err != nil {
// 		response.ShowResponse(err.Error(), HTTP_BAD_REQUEST, "Failure", nil, ctx)
// 		return
// 	}
// 	err = json.Unmarshal(reqBody, &data)
// 	if err != nil {
// 		response.ShowResponse(err.Error(), HTTP_BAD_REQUEST, "Failure", nil, ctx)
// 		return
// 	}
// }

func RequestDecoding(context *gin.Context, data interface{}) {

	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
}

func SetHeader(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

}
