package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	listner, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalf("Error setting up listner: %s\n", err.Error())
	}
	defer listner.Close()
	fmt.Println("Listening...")
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatalf("Error waiting for connection: %s\n", err.Error())
		}
		fmt.Printf("Connection recieved from: %s\n", conn.RemoteAddr().String())

		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Printf("%s\n", line)
		}
		fmt.Println("Connection closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		lineBuffer := ""
		defer close(lines)
		for {
			buffer := make([]byte, 8)
			n, err := f.Read(buffer)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
			}
			readBuffer := string(buffer[:n])
			lineBuffer = parseLines(strings.Split(readBuffer, "\n"), lineBuffer, lines)
		}
		// If there was no newline character in the message
		if lineBuffer != "" {
			lines <- lineBuffer
		}
	}()
	return lines
}

func parseLines(parts []string, lineBuffer string, lines chan string) string {
	if len(parts) == 1 {
		return lineBuffer + parts[0]
	}
	lines <- lineBuffer + parts[0]
	return parseLines(parts[1:], "", lines)
}
