package main

import (
	"websocket-htmx/components"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

func Websocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		ChatLoop(ws, c)
	}

}

func GetChatRoomPage(c echo.Context) error {
	return render(c, components.ChatPage())
}
