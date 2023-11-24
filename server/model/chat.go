package model 





type Room struct{

	RoomId string `json:"sessionId" gorm:"default:uuid_generate_v4();unique;primaryKey"`
	Name string `json:"name"`
	Creator string `json:"creator"`
}

type RoomMembers struct{

	RoomId string `json:"roomId"`
	UserId string `json:"userId"`

}