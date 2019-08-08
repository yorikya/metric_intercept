package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	hostsFilePath        = "/etc/hosts"
	localhostCarbonRelay = "127.0.0.1 carbonrelay-consul\n"
)

func main() {
	fmt.Println("hello world")

	err := editHostFile()
	fmt.Println("the error:", err)

	l, err := net.Listen("tcp", "localhost:2001")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// Make a buffer to hold incoming data.
		buf := make([]byte, 1024)
		// Read the incoming connection into the buffer.
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		fmt.Println("Message", string(buf))
		// Close the connection when you're done with it.
		conn.Close()
	}

}

func editHostFile() error {
	hostsFile, err := os.OpenFile(hostsFilePath, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		return err
	}
	defer hostsFile.Close()

	scanner := bufio.NewScanner(hostsFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "carbonrelay-consul") {
			if strings.HasPrefix(line, "#") {
				return fmt.Errorf("please remove comment tag from: %s", line)
			}
			fmt.Println("found row for metrics redirection, skipping editing...")
			return nil
		}
	}
	_, err = hostsFile.WriteString(localhostCarbonRelay)
	if err != nil {
		return err
	}
	return nil
}
