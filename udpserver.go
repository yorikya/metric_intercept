package main

import (
	"bytes"
	"log"
	"net"
)

const (
	UDP = "udp"
)

// func main() {
func _main() {
	address, _ := net.ResolveUDPAddr(UDP, ":8125")
	listener, err := net.ListenUDP(UDP, address)
	if err != nil {
		log.Fatalf("ListenAndServe: %s", err.Error())
	}
	log.Println("start UDP server")
	for {
		message := make([]byte, 512)
		n, remaddr, error := listener.ReadFrom(message)
		if error != nil {
			continue
		}
		buf := bytes.NewBuffer(message[0:n])
		log.Printf("addr:%s, get message: %s\n", remaddr.String(), buf.String())

	}
	listener.Close()
}
