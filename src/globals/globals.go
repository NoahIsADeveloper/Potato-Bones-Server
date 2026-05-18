package globals

import "flag"

// Flags
var Port *int
var Host *string

var Tickrate *int
var SessionLength *int

var OnlyReadTCP *bool
var OnlySendTCP *bool

var DebugShowOutgoing *bool
var DebugShowIncoming *bool
var DebugLobbyInfo *bool
var DebugHideMovePacket *bool

var MaxEntities *int
var GameSpeed *float64
var MaxClients *int
var MaxLobbies *int

var MaxPacketSize *int

func ParseFlags() {
	// Potatos are 4th most grown crop, and there's 206 bones in a human body!
	Port = flag.Int("port", 4206, "Port to run the server on")

	Host = flag.String("host", "0.0.0.0", "Host address to bind to")
	Tickrate = flag.Int("tickrate", 50, "Server update rate (in miliseconds)")
	MaxEntities = flag.Int("max-entities", 255, "Max Entities per world")
	MaxClients = flag.Int("max-clients", 255, "Max Clients per world")
	MaxLobbies = flag.Int("max-lobbies", 3, "Max Lobbies")
	MaxPacketSize = flag.Int("max-packet-size", 1024, "Max incoming packet size in bytes")
	GameSpeed = flag.Float64("gamespeed", 1, "Game speed multiplier")
	DebugShowOutgoing = flag.Bool("debug-outgoing", false, "Print outgoing packets")
	DebugShowIncoming = flag.Bool("debug-incoming", false, "Print incoming packets")
	DebugHideMovePacket = flag.Bool("debug-hide-move", false, "Hides the move packet when printing incoming/outgoing packets")
	DebugLobbyInfo = flag.Bool("debug-lobby", false, "Print lobby updates")
	SessionLength = flag.Int("session-length", 1440, "How long before a session expires (in minutes)")
	OnlySendTCP = flag.Bool("only-send-tcp", false, "Always use TCP over UDP for outgoing packets")
	OnlyReadTCP = flag.Bool("only-read-tcp", false, "Disables the UDP server")

	flag.Parse()
}