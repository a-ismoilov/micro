package main

import "net"

func main() {
	net.Listen("tcp", ":3333")
}
