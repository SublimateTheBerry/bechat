package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Color    string `json:"color"`
	Message  string `json:"message"`
	Platform string `json:"platform"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket Upgrade error:", err)
		return
	}
	defer conn.Close()
	clients[conn] = true

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			delete(clients, conn)
			break
		}
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		msgJSON, _ := json.Marshal(msg)

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msgJSON)
			if err != nil {
				log.Printf("websocket error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func connectTwitchChat(channel string) {
	client := twitch.NewAnonymousClient()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		msg := Message{
			Username: message.User.DisplayName,
			Color:    message.User.Color,
			Message:  message.Message,
			Platform: "twitch",
		}
		broadcast <- msg
	})

	client.Join(channel)
	err := client.Connect()
	if err != nil {
		fmt.Println("Error connecting to Twitch chat:", err)
	}
}

func connectGoodgameChat(channelID string) {
	url := "wss://chat.goodgame.ru/chat/websocket"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Error connecting to Goodgame chat:", err)
	}
	defer conn.Close()

	message := fmt.Sprintf(`{"type":"join","data":{"channel_id":%s,"hidden":false}}`, channelID)
	conn.WriteMessage(websocket.TextMessage, []byte(message))

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		parsedMessage := parseGoodgameMessage(msg)
		parsedMessage.Platform = "gg"
		broadcast <- parsedMessage
	}
}

func parseGoodgameMessage(msg []byte) Message {
	var rawMessage map[string]interface{}
	err := json.Unmarshal(msg, &rawMessage)
	if err != nil {
		log.Println("Error parsing Goodgame message:", err)
		return Message{Username: "Unknown", Color: "#FFFFFF", Message: "Error parsing message"}
	}

	if rawMessage["type"] == "message" {
		data := rawMessage["data"].(map[string]interface{})
		username := data["user_name"].(string)
		messageText := data["text"].(string)
		userColor := data["user_color"].(string)

		return Message{
			Username: username,
			Color:    userColor,
			Message:  messageText,
		}
	}

	return Message{}
}

func main() {
	twitchUsername := flag.String("tu", "", "Twitch username to connect to")
	goodgameUsername := flag.String("ggu", "", "Goodgame username to connect to")

	flag.Parse()

	if *twitchUsername == "" && *goodgameUsername == "" {
		fmt.Printf("Please provide at least one username using --tu (Twitch) or --ggu (Goodgame) flags.")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/ws", wsHandler)

	go handleMessages()

	if *twitchUsername != "" {
		go connectTwitchChat(*twitchUsername)
	}

	if *goodgameUsername != "" {
		go connectGoodgameChat(*goodgameUsername)
	}

	port := ":12597"
	fmt.Printf("Server running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
