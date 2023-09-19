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

func AlreadyExists(data string) bool {

	return db.RecordExist("users", data, "email")

}

func SignupService(ctx *gin.Context, input *request.AuthRequest) {

	var user model.User
	//check the credentials if already exists

	if AlreadyExists(input.Email) {

		fmt.Println("email already exists", input.Email)
		response.ShowResponse("email already exists", utils.HTTP_BAD_REQUEST, "Bad Request", "", ctx)
		return
	}


	user.Email = input.Email
	user.FullName = input.FullName
	//encrypt the password then store in db

	encryptedPassword, _ := utils.HashPassword(input.Password)

	user.Password = encryptedPassword

	err := db.CreateRecord(&user)
	if err != nil {

		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, "server error", "", ctx)
		return
	} else {

		response.ShowResponse("Signup success", utils.HTTP_OK, "Success", "", ctx)
	}
}

func LoginService(ctx *gin.Context, input *request.AuthRequest) {
	
	var user *model.User
	var userClaims token.Claims

	//check if the user exists in db or not
	if !(db.RecordExist("users", input.Email, "email")) {
		//return

		response.ShowResponse("user doesn't exist", 400, "Bad request", "", ctx)
		return
	}

	//get the encrypted password from the db and then compare
	db.FindById(&user, input.Email, "email")
	fmt.Println("user:", user)

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		//RETURN

		response.ShowResponse("Password Doesn't Match", 401, "Unauthorized", "", ctx)
		return
	}

	//if password is correct ,provide a token to the user


	userClaims.UserId = user.UserId
	Token := token.GenerateToken(userClaims, ctx)

	//create a cookie, store the value of the token in the http cookie
	cookie := &http.Cookie{Name: "Auth", Value: Token}

	http.SetCookie(ctx.Writer, cookie)

	//show a success response to login attempt

	response.ShowResponse("Login Success", utils.HTTP_OK, "Success", "", ctx)

}
