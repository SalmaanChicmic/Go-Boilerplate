package request

type AuthRequest struct {
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	NewPassword string `json:"newPassword" binding:"required" example:"11111111"`
	Password    string `json:"password" binding:"required" example:"11111111"`
}
