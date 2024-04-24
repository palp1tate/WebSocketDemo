package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChatHandler(c *gin.Context) {
	user := c.Query("user")
	if user == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名参数不能为空"})
		return
	}
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket 升级失败：", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket 升级失败"})
		return
	}
	defer closeAndDeleteConnection(user, ws)

	webSocketMap.Store(user, ws)
	addOnlineCount()
	broadcastUserEvent(user, "加入了")
	log.Printf("用户连接：%s，当前在线人数：%d\n", user, getOnlineCount())

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			break
		}
		var msg Message
		if err := json.Unmarshal(p, &msg); err != nil {
			sendMessage(user, "消息格式错误："+err.Error())
			continue
		}

		if msg.ToUser == "" || msg.Content == "" {
			sendMessage(user, "接收人或消息内容不能为空")
			continue
		} else if _, ok := webSocketMap.Load(msg.ToUser); !ok {
			sendMessage(user, "接收人不在线")
			continue
		}

		msgFrom := fmt.Sprintf("来自%s的消息：%s", user, msg.Content)
		sendMessage(msg.ToUser, msgFrom)
	}
}
