package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/gorilla/websocket"
)

var (
	webSocketMap sync.Map
	onlineCount  int32
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func sendMessage(toUser, message string) {
	if conn, ok := webSocketMap.Load(toUser); ok {
		if err := conn.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Printf("Error sending message to %s: %s\n", toUser, err)
		}
	}
}

func broadcastMessage(message string) {
	webSocketMap.Range(func(key, value interface{}) bool {
		go func(conn *websocket.Conn) {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Printf("Error broadcasting message: %s\n", err)
			}
		}(value.(*websocket.Conn))
		return true
	})
}

func broadcastUserEvent(user string, event string) {
	currentUsers := getCurrentUserList()
	message := fmt.Sprintf("%s%s聊天室，当前在线人数：%d。\n在线用户列表：%s\n", user, event, getOnlineCount(), currentUsers)
	broadcastMessage(message)
}

func closeAndDeleteConnection(user string, conn *websocket.Conn) {
	err := conn.Close()
	if err != nil {
		log.Println("Error closing connection: ", err)
	}
	webSocketMap.Delete(user)
	reduceOnlineCount()
	broadcastUserEvent(user, "离开了")
}

func addOnlineCount() int32 {
	return atomic.AddInt32(&onlineCount, 1)
}

func getOnlineCount() int32 {
	return atomic.LoadInt32(&onlineCount)
}

func reduceOnlineCount() int32 {
	return atomic.AddInt32(&onlineCount, -1)
}

func getCurrentUserList() string {
	users := make([]string, 0, 1000)
	webSocketMap.Range(func(key, value interface{}) bool {
		user := key.(string)
		users = append(users, user)
		return true
	})
	userList, _ := json.Marshal(users)
	return string(userList)
}
