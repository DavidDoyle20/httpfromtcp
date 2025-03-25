package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main (){
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println("Unable to create listener", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error with accepting listener connection: ", err)
			return
		}
		fmt.Println("Connection has been accepted!")

		lines := getLinesChannel(conn)
		for line := range lines  {
			fmt.Printf("%s\n", line)
		}
		fmt.Println("Channel has been closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func () {
		defer close(lines)
		defer f.Close()
		bytes := make([]byte, 8)
		currentLine := ""
		for {
			numBytes, err := f.Read(bytes)
			if err != nil {
				if err == io.EOF {
					lines <- currentLine
					break
				}
				fmt.Println("Error reading file: ", err)
				return
			}
	
			parts := strings.Split(string(bytes[:numBytes]), "\n")
	
			if len(parts) > 1 {
				for i := 0; i < len(parts)-1; i++ {
					currentLine += parts[i]
					lines <- currentLine
					currentLine = ""
				}
			}
			currentLine += parts[len(parts)-1]
		}
	}()
	return lines
}