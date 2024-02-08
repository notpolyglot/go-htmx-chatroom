package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", GetChatRoomPage)
	e.GET("/ws", Websocket)
	e.Logger.Fatal(e.Start(":1323"))
}
