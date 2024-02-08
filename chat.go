package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"websocket-htmx/components"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type HTMXMessage struct {
	ChatMessage string `json:"chat_message"`
	Headers     struct {
		HXRequest     string `json:"HX-Request"`
		HXTrigger     string `json:"HX-Trigger"`
		HXTriggerName string `json:"HX-Trigger-Name"`
		HXTarget      string `json:"HX-Target"`
		HXCurrentURL  string `json:"HX-Current-URL"`
	} `json:"HEADERS"`
}

func ChatLoop(ws *websocket.Conn, c echo.Context) {
	_, p, err := ws.ReadMessage()
	if err != nil {
		fmt.Println(err)
	}

	var msg HTMXMessage
	json.Unmarshal(p, &msg)

	var buf bytes.Buffer
	components.SentAndRecv(msg.ChatMessage, "You're an idiot").Render(context.Background(), &buf)
	err = ws.WriteMessage(websocket.TextMessage, buf.Bytes())
	if err != nil {
		c.Logger().Error(err)
	}

}
