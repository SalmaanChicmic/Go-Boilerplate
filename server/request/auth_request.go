package request

type AuthRequest struct {
	Email       string `json:"email" binding:"required,email" example:"user@example.com"`
	NewPassword string `json:"newPassword" binding:"required" example:"11111111"`
	Password    string `json:"password" binding:"required" example:"11111111"`
}
