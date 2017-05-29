package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	Server := os.Getenv("SERVER_ADDRESS")
	fmt.Println("SERVER_ADDRESS=" + Server)

	serverAddr, err := net.ResolveUDPAddr("udp", Server+":10001")
	fmt.Println("ServerAddr=" + serverAddr.String())
	checkError(err)

	conn, err := net.DialUDP("udp", nil, serverAddr)
	checkError(err)
	defer conn.Close()

	buf := make([]byte, 1024)
	i := 0

	for {
		msg := strconv.Itoa(i)
		fmt.Fprintf(conn, msg)
		i++

		n, addr, err := conn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		time.Sleep(time.Second * 1)
	}
}
