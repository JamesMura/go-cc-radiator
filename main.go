package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"code.google.com/p/gcfg"
	"github.com/gorilla/handlers"
	"github.com/jamesmura/go-cc-radiator/ccrad"
	"golang.org/x/net/websocket"
)

type Config struct {
	Connection struct {
		Url            string
		Auth           bool
		Username       string
		Password       string
		UpdateInterval int
	}
}

type WebSocketEvent struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func getData(cfg Config) ccrad.Projects {
	response, err := http.Get(cfg.Connection.Url)
	if err != nil {
		log.Fatal("Request:", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Response:", err)
		os.Exit(1)
	}
	return ccrad.ParseXML(contents)

}

func Check(ws *websocket.Conn) {
	var err error
	var cfg Config
	err = gcfg.ReadFileInto(&cfg, "conf/app.ini")
	if err != nil {
		log.Fatal("Config:", err)
		return
	}

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		count := 1
		for {
			query := getData(cfg)
			event := WebSocketEvent{Event: "UPDATE", Data: query}
			b, err := json.Marshal(event)
			msg := fmt.Sprintf("message:  %v", string(b))
			fmt.Println("Sending to client: " + msg)
			if err = websocket.Message.Send(ws, string(b)); err != nil {
				log.Fatal("Cant send:", err)
				break
			}

			count += 1
			time.Sleep(time.Duration(cfg.Connection.UpdateInterval) * time.Second)
		}

	}
}

func main() {

	fmt.Println("starting the server")
	http.Handle("/", http.FileServer(http.Dir("./public/")))
	http.Handle("/check", websocket.Handler(Check))

	if err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
