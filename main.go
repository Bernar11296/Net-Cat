package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"a/pkg"
)

func main() {
	port := ""
	arg := os.Args[1:]
	if len(arg) > 1 {
		fmt.Printf("[USAGE]: ./TCPChat $port")
		return
	}
	if len(arg) == 0 {
		port = "8989"
	} else if len(arg) == 1 {
		_, err := strconv.Atoi(arg[0])
		if err != nil {
			fmt.Println("ENTER NUMBER!")
			return
		}
		port = arg[0]
	}
	fmt.Printf("Listening on the port" + " : " + port)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	go pkg.Broadcaster()
	defer l.Close()
	for {

		conn, err := l.Accept()
		if err != nil {
			return
		}
		pkg.Mu.Lock()

		if len(pkg.Clients) > 9 {
			conn.Write([]byte("Chat is full"))
			conn.Close()
		}
		pkg.Mu.Unlock()
		go pkg.UserConnection(conn)

	}
}
