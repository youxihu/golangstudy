package main

import (
	"gogo/day1-day10/day10/netprogram/client"
	"gogo/day1-day10/day10/netprogram/server"
	"time"
)

func main() {
	server.CreateServer()
	time.Sleep(1 * time.Second)
	client.CreateClient()
}
