package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatalf("Error resolving address: %s\n", err.Error())
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Error connecting: %s\n", err.Error())
	}
	defer conn.Close()
	for {
		fmt.Printf(">")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		_, err := conn.Write([]byte(text))
		if err != nil {
			log.Fatalf("Error writing to conn: %s\n", err.Error())
		}
	}
}
