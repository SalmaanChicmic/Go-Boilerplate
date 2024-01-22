package model

type User struct {
	// Id       string ` gorm:"default:uuid_generate_v4();unique" json:"id"`
	UserId       string `gorm:"type:uuid;default:uuid_generate_v4();unique" json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}
