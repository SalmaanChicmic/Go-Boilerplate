package authentication

import (
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AlreadyExists(data string) bool {

	return db.RecordExist("user", data, "email")

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignupService(ctx *gin.Context, input *request.AuthRequest) {

	//check the credentials if already exists

	if AlreadyExists(input.Email) {
		//return
		return
	}

	var user *model.User

	user.Email = input.Email
	//encrypt the password then store in db

	encryptedPassword,_:=HashPassword(input.Password)

	user.Password=encryptedPassword


	err:=db.CreateRecord(&user)
	if err != nil {

		response.ShowResponse(err.Error(),500,"server error","",ctx)
		return 
	}
}


func LoginService(ctx *gin.Context,input *request.AuthRequest){


	//check if the user exists in db or not 

	if !(db.RecordExist("user",input.Email,"email")){
		//return

		response.ShowResponse("user doesn't exist",400,"Bad request","",ctx)
		return
	}

	//get the encrypted password from the db


	var user *model.User
	db.FindById(user, input.Email,"email")






}