package pkg

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func UserConnection(conn net.Conn) {
	defer conn.Close()
	username, errBool := GetName(conn)

	if len(username) < 3 {
		fmt.Fprint(conn, "ENTER A LONGER NAME,PLEASE!")
		return
	}
	if errBool {
		return
	}
	Mu.Lock()
	Clients[username] = conn
	Mu.Unlock()
	conn.Write([]byte(history))
	join <- newMessages("  joined.", username)
	history += fmt.Sprintf("%s joined\n", username)
	time := time.Now().Format("2006-01-02 15:04:05")
	output := fmt.Sprintf("[%s][%s]:", time, username)
	input := bufio.NewScanner(conn)
	fmt.Fprint(conn, output)
	for input.Scan() {
		text := Checktext(input.Text())
		if LastCheckText(text) {
			output := fmt.Sprintf("[%s][%s]:%s", time, username, input.Text())

			history += output + "\n"
			messages <- newMessages(output, username)
		} else {
			fmt.Fprintln(conn, "Enter normal text")
			fmt.Fprint(conn, "["+time+"]"+"["+username+"]"+":")
		}
	}
	Mu.Lock()
	delete(Clients, username)
	Mu.Unlock()
	leaving <- newMessages(" has left.", username)
	history += fmt.Sprintf("%s has left\n", username)
}
