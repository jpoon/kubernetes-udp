package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	Server := os.Getenv("SERVER_ADDRESS")
	fmt.Println("SERVER_ADDRESS=" + Server)

	ServerAddr, err := net.ResolveUDPAddr("udp", Server+":10001")
	fmt.Println("ServerAddr=" + ServerAddr.String())
	CheckError(err)

	Conn, err := net.DialUDP("udp", nil, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	i := 0

	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := Conn.Write(buf)

		if err != nil {
			fmt.Println(msg, err)
		}

		time.Sleep(time.Second * 1)
	}
}
