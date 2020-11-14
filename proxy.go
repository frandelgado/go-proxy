package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 2 {
		log.Fatal("Please provide a port number, server host and port number")
		return
	}
	PORT := ":" + arguments[1]
	serverAddress := arguments[2]
	l, err := net.Listen("tcp", PORT)
	log.Printf("Proxy running at %s", PORT)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c, serverAddress)
	}

}

func handleConnection(client net.Conn, serverAddress string) {
	clientAddress := client.RemoteAddr().String()
	fmt.Printf("Serving %s\n", clientAddress)
	server, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Could not connect to remote server %s", serverAddress)
	}
	log.Printf("Connected to server %s", serverAddress)

	go func() { _, _ = io.Copy(client, server) }()
	go func() { _, _ = io.Copy(server, client) }()

}