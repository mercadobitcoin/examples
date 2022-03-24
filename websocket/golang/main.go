package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	socketUrl := "wss://ws.mercadobitcoin.net/ws"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()

	fmt.Println("connected")

	// listen to messages
	go func() {
		for {
			var data json.RawMessage
			if err := conn.ReadJSON(&data); err != nil {
				return
			}
			fmt.Println(string(data))
		}
	}()

	// send subscription message
	if err := conn.WriteMessage(websocket.TextMessage, []byte("{\"type\": \"subscribe\",\"subscription\": {\"name\": \"orderbook\",\"id\": \"BRLBTC\", \"limit\": 10}}")); err != nil {
		return
	}

	<-done
}
