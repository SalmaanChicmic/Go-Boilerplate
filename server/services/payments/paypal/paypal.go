package paypal

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"main/server/response"
	"main/server/utils"
	"net/http"

	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var baseURL = os.Getenv("PAYPAL_BASE_URL")
var accessToken string

var webHookURL = "https://model-bass-settling.loca.lt"

func GetAccessToken(ctx *gin.Context) {

	authString := base64.StdEncoding.EncodeToString([]byte(os.Getenv("PAYPAL_CLIENT_ID") + ":" + os.Getenv("PAYPAL_SECRET_ID")))
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", baseURL+"/v1/oauth2/token", strings.NewReader(data.Encode()))
	if err != nil {
		response.ShowResponse(err.Error(), int64(req.Response.StatusCode), utils.FAILURE, nil, ctx)
		return

	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+authString)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		response.ShowResponse(err.Error(), int64(resp.StatusCode), utils.FAILURE, nil, ctx)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	if resp.StatusCode != http.StatusOK {
		response.ShowResponse("error getting access token", int64(resp.StatusCode), utils.FAILURE, body, ctx)
		return
	}
	var mp map[string]interface{}
	if err := json.Unmarshal(body, &resp); err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	accessToken, ok := mp["access_token"].(string)
	if !ok {

		response.ShowResponse("error parsing access token from response", utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, body, ctx)

		return
	}

	fmt.Println("access_token: ", accessToken)
	// return accessToken, nil

	response.ShowResponse(utils.SUCCESS, utils.HTTP_OK, utils.SUCCESS, accessToken, ctx)

}

func CreateWebhook(ctx *gin.Context) {

	payload := map[string]interface{}{
		"url": webHookURL + "/order-approved",
		"event_types": []map[string]string{
			{"name": "CHECKOUT.ORDER.APPROVED"},
		},
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	req, err := http.NewRequest("POST", baseURL+"/v1/notifications/webhooks", bytes.NewBuffer(payloadJSON))
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body: ", string(body))
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		response.ShowResponse("error subscribing to webhook. ", int64(resp.StatusCode), utils.FAILURE, resp.Status, ctx)
		return
	}

	fmt.Println("Sucess in creating the webhook")
	response.ShowResponse("Successfully created webhook", utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}

func HandleWebhookNotification(ctx *gin.Context) {

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	var mp map[string]interface{}
	err = json.Unmarshal(body, &mp)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	res := mp["resource"].(map[string]interface{})
	fmt.Println(res["status"])
	paymentId := res["id"].(string)

	if res["status"].(string) == "APPROVED" {
		err = CapturePayment(paymentId)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
			return
		}
		response.ShowResponse("Payment Capture success", utils.HTTP_OK, utils.SUCCESS, paymentId, ctx)
	} else {
		fmt.Println("Status of order is:", res["status"].(string))

	}

}

func CapturePayment(id string) error {

	fmt.Println("Capture payment hit")

	req, err := http.NewRequest("POST", baseURL+"/v2/checkout/orders/"+id+"/capture", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("error executing payment. Status: %s, Body: %s", resp.Status, body)
	}

	return nil
}
