package main

import (
	"flag"
	"fmt"
	"goserver/websocketlib"
	"log"
	"net/http"
	"os"
)

func init() {
	websocketlib.GenPem()
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Llongfile)
	port := flag.String("p", "8081", "https port")
	http.HandleFunc("/ws", websocketlib.WebsocketHandler)
	fmt.Println("Web listening :" + *port)
	panic(http.ListenAndServeTLS(":"+*port, "cert.pem", "key.pem", nil))
}
