package pkg

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	Clients = make(map[string]net.Conn)
	Mu      sync.Mutex
	leaving = make(chan message)
	join    = make(chan message)
)

type message struct {
	text    string
	address string
}

func newMessages(msg string, username string) message {
	return message{
		text:    msg,
		address: username,
	}
}

var (
	messages = make(chan message)
	history  string
)

func Broadcaster() {
	for {
		select {
		case msg := <-messages:
			Mu.Lock()
			for username, conn := range Clients {
				time := time.Now().Format("2006-01-02 15:04:05")
				if msg.address != username {
					fmt.Fprintln(conn, ClearLine(msg.text)+msg.text)
				}

				fmt.Fprint(conn, "["+time+"]"+"["+username+"]"+":")

			}
			Mu.Unlock()
		case msg := <-leaving:
			Mu.Lock()
			for username, conn := range Clients {
				time := time.Now().Format("2006-01-02 15:04:05")
				fmt.Fprintln(conn, ClearLine(msg.address+msg.text)+msg.address+msg.text)
				fmt.Fprint(conn, "["+time+"]"+"["+username+"]"+":")
			}
			Mu.Unlock()
		case msg := <-join:
			Mu.Lock()
			for username, conn := range Clients {
				if msg.address != username {

					fmt.Fprintln(conn)
					time := time.Now().Format("2006-01-02 15:04:05")

					fmt.Fprintln(conn, ClearLine(msg.address+msg.text)+msg.address+msg.text)
					fmt.Fprint(conn, "["+time+"]"+"["+username+"]"+":")

				} else {
					continue
				}
			}
			Mu.Unlock()

		}
	}
}
