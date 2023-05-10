package socket

import (
	"main/server/handler"

	socketio "github.com/googollee/go-socket.io"
)

func SocketInit() *socketio.Server {
	server := socketio.NewServer(nil)
	handler.SocketHandler(server)
	go server.Serve()
	return server
}
