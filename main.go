package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	listenAddr string
	wsAddr     string
	jsTemplate *template.Template
)

func init() {
	flag.StringVar(&listenAddr, "listen-addr", "", "Address to listen on")
	flag.StringVar(&wsAddr, "ws-addr", "", "Address for websocket connection")
	flag.Parse()
	var err error

	jsTemplate, err = template.ParseFiles("logger.js")
	if err != nil {
		panic(err)
	}
}


func serveIndex(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// http.Error(w, "", 500)
		log.Println(err)
		return
	}
	defer conn.Close()
	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			// http.Error(w, "", 500)
			log.Println(err)
			
			return
		}
		fmt.Printf("From %s: %s\n", conn.RemoteAddr().String(), string(msg))
	}

}

func serveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	// jsTemplate.Execute(w, nil)
	jsTemplate.Execute(w, struct{ WSAddr string }{ wsAddr })
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveIndex)
	r.HandleFunc("/ws", serveWS)
	r.HandleFunc("/k.js", serveFile)
	log.Fatal(http.ListenAndServe(":8080", r))
}
