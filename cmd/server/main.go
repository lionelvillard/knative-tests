package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received.")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func ip() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr)
}

func main() {
	ip()
	http.HandleFunc("/", handler)
	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
