package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	localAddr, err := net.ResolveUDPAddr("udp", ":8989")
	if err != nil {
		panic(err)
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", "224.0.0.250:9981")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			data := make([]byte, 1024)
			n, addr, err := conn.ReadFrom(data)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(data[:n]), "from", addr)
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	for i := 0; i < 10000; i++ {
		select {
		case <-ticker.C:
			fmt.Println("ping")
			conn.WriteToUDP([]byte("ping"), remoteAddr)
		}
	}

}
