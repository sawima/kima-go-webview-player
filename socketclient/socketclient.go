package socketclient

import (
	"encoding/json"
	"log"

	"golang.org/x/net/websocket"
)

type recieveURL struct {
	Url    string `json:"url"`
	Reload bool   `json:"reload"`
}

func ReadSocket(input string) (string, bool) {
	socketData := &recieveURL{}
	origin := "http://127.0.0.1/"
	url := "ws://127.0.0.1:8899"
	ws, err := websocket.Dial(url, "echo-protocolx", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte(input + "\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(msg[:n], socketData)
	return socketData.Url, socketData.Reload
}
