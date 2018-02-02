package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	udp, err := net.ResolveUDPAddr("udp", "127.0.0.1:43123")
	checkError(err)
	conn, err := net.ListenUDP("udp", udp)
	defer conn.Close()
	checkError(err)
	var buf [20]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	checkError(err)
	println(addr)
	println(n)
	println(string(buf[0:n]))

}
