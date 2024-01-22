package authentication

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/services/token"
	"main/server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//checks whether the user present in the database with the same email address
func AlreadyExists(data string) bool {

	return db.RecordExist("users", data, "email")

}

func Signup(ctx *gin.Context, input *request.AuthRequest) {

	var user model.User

	//check the credentials if already exists
	if AlreadyExists(input.Email) {

		fmt.Println("email already exists", input.Email)
		response.ShowResponse(utils.EMAIL_EXISTS, utils.HTTP_BAD_REQUEST, utils.FAILURE, "", ctx)
		return
	}

	user.Email = input.Email
	user.FullName = input.FullName
	//encrypt the password then store in db

	encryptedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	user.Password = *encryptedPassword

	err = db.CreateRecord(&user)
	if err != nil {

		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, "", ctx)
		return
	} else {

		response.ShowResponse(utils.SIGNUP_SUCCESS, utils.HTTP_OK, utils.SUCCESS, "", ctx)
	}
}

func Login(ctx *gin.Context, input *request.AuthRequest) {

	var user *model.User
	var userClaims model.Claims

	//check if the user exists in db or not
	if !(db.RecordExist("users", input.Email, "email")) {
		//return

		response.ShowResponse(utils.USER_NOT_FOUND, utils.HTTP_BAD_REQUEST, utils.FAILURE, "", ctx)
		return
	}

	//get the encrypted password from the db and then compare
	db.FindById(&user, input.Email, "email")
	fmt.Println("user:", user)

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		//RETURN

		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, "", ctx)
		return
	}

	//if password is correct ,provide a token to the user

	userClaims.Id = user.UserId
	Token, err := token.GenerateToken(userClaims)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	//create a cookie, store the value of the token in the http cookie
	cookie := &http.Cookie{Name: "Auth", Value: *Token}

	http.SetCookie(ctx.Writer, cookie)

	//show a success response to login attempt

	response.ShowResponse(utils.LOGIN_SUCCESS, utils.HTTP_OK, utils.SUCCESS, "", ctx)

}
