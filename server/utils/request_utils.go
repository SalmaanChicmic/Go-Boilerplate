package utils

import (
	"encoding/json"
	"io/ioutil"
	"main/server/response"

	"github.com/gin-gonic/gin"
)

func RequestDecoding(ctx *gin.Context, data interface{}) {

	reqBody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		response.ShowResponse(err.Error(), HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}
	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		response.ShowResponse(err.Error(), HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}
}
