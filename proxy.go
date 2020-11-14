package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 3 {
		log.Fatal("Please provide a port number, server host and port number")
		return
	}
	PORT := ":" + arguments[1]
	serverAddress := arguments[2]
	l, err := net.Listen("tcp4", PORT)
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
	for {
		netData, err := bufio.NewReader(client).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		temp := strings.TrimSpace(netData)
		if temp == "stop" {
			fmt.Printf("Goodbye, %s!\n", clientAddress)
			break
		}

		result := strconv.Itoa(1) + "\n"
		client.Write([]byte(string(result)))
	}
}