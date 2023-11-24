package request 


type RoomCreateRequest struct {

	RoomName string `json:"room_name"`
}

type RoomJoin struct{

	RoomId string `json:"room_id"`
}