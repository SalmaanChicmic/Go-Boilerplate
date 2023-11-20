package awssns

import (
	"fmt"
	"main/server/request"
	"main/server/response"
	"main/server/utils"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gin-gonic/gin"
)

func SendSMS(ctx *gin.Context, reqBody request.AwsTextMessagingRequest) {

	// Create Session and assign AccessKeyID and SecretAccessKey
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACESSS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	},
	)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// Create SNS service
	svc := sns.New(sess)

	fmt.Println("svc", svc.Endpoint)

	// Pass the phone number and message.
	params := &sns.PublishInput{
		PhoneNumber: aws.String(reqBody.PhoneNumber),
		Message:     aws.String(reqBody.Message),
	}

	// sends a text message (SMS message) directly to a phone number.
	resp, err := svc.Publish(params)

	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	fmt.Println(resp) // print the response data.

	response.ShowResponse("Message sent sucessfully", utils.HTTP_OK, utils.SUCCESS, resp, ctx)

}
