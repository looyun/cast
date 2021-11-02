package main

import (
	"fmt"
	"net"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", "224.0.0.250:9981")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listen on addr: %s\n", addr.String())

	data := make([]byte, 1024)
	for {
		n, udpAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			panic(err)
		}
		fmt.Println(udpAddr.String(), string(data[:n]))
		_, err = conn.WriteToUDP([]byte("pong"), udpAddr)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("done")
}
