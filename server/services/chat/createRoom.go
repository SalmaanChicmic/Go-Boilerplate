package chat

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/services/token"

	socketio "github.com/googollee/go-socket.io"
)

func RoomCreate(s socketio.Conn, data map[string]interface{}) {
	// fmt.Println("inside room creation process...")
	// Get the user ID from the query params
	var room model.Room
	var participant model.RoomMembers

	headerToken := s.RemoteHeader().Get("authToken")
	claims, err := token.DecodeToken(headerToken)
	if err != nil {

		fmt.Println("error in token decoding", err)
		return
	}

	// s.Emit("hello","hi there")
	// utils.SocketServerInstance.BroadcastToNamespace("/","reply",data["message"])

	room.Creator = claims.Id
	roomName, ok := data["name"].(string)
	if !ok {
		fmt.Println("not ok")
		response.SocketResponse("error in type assertion", 400, "failure", nil, "createRoom", s)
		return
	}
	room.Name = roomName
	db.CreateRecord(&room)

	s.Join(room.RoomId)
	s.Emit("createRoom", claims.Id+" :is connected to the room")

	//update the participant table with roomid and user id

	participant.RoomId = room.RoomId
	participant.UserId = claims.Id

	db.CreateRecord(&participant)

	// s.Emit("createRoom","room created successfully")
	// response.SocketResponse("Success", "room created Successfully", s)

}

func RoomJoin(s socketio.Conn, req request.RoomJoin) {

	fmt.Println("Room join process...")

}
