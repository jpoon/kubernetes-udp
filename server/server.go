package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, msg []byte) {
	_, err := conn.WriteToUDP(msg, addr)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	port := "10001"

	fmt.Println("Listening on", port)

	/* Hostname */
	hostname, err := os.Hostname()
	fmt.Println("Hostname=" + hostname)
	checkError(err)

	/* Lets prepare a address at any address at port 10001*/
	serverAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	checkError(err)

	/* Now listen at selected port */
	conn, err := net.ListenUDP("udp", serverAddr)
	checkError(err)
	defer conn.Close()

	msg := []byte(hostname)
	buf := make([]byte, 2048)

	for {
		_, remoteaddr, err := conn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf), " from ", remoteaddr)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		go sendResponse(conn, remoteaddr, msg)
	}
}
