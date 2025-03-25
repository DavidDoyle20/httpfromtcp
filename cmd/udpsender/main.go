package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = ":42069"
func main() {
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatal("Could not resolve UDP address: ", err)
	}
	
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Could not dial UDP: ", err)
	}
	defer conn.Close()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := r.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading line: ", err)
			break
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatal("Unable to write to connection: ", err)
			break
		}
	}
}