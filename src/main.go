package main

import (
	"fmt"
	"net"
	"potato-bones/src/globals"
	"potato-bones/src/networking"
)

func HandleTCPServer() {
	ln, err := net.Listen("tcp", *globals.Host + ":" + fmt.Sprint(*globals.Port))
	if (err != nil) { panic(err) }
	fmt.Println("TCP Server running on " + *globals.Host + ":" + fmt.Sprint(*globals.Port))

	for {
		conn, err := ln.Accept()
		if err != nil { continue }
		go networking.HandleTCPClient(conn)
	}
}

func HandleUDPServer() {
	addr, err := net.ResolveUDPAddr("udp", *globals.Host + ":" + fmt.Sprint(*globals.Port))
	if err != nil { panic(err) }

	udpConn, err := net.ListenUDP("udp", addr)
	if err != nil { panic(err) }
	defer udpConn.Close()
	networking.SetUDPConn(udpConn)

	fmt.Println("UDP Server running on " + *globals.Host + ":" + fmt.Sprint(*globals.Port))

	buffer := make([]byte, *globals.MaxPacketSize)
	for {
		length, clientAddr, err := udpConn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		data := buffer[:length]
		go networking.HandleUDPPacket(clientAddr, data)
	}
}

func StartServers() {
    networking.InitNetworking()
    networking.StartUpdateLoop(*globals.Tickrate)

    if !*globals.OnlyReadTCP {
        go HandleUDPServer()
    }

    go HandleTCPServer()
}


func main() {
    globals.ParseFlags()
    fmt.Println("Starting servers...")
    StartServers()

    select {}
}
