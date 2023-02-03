package pkg

import (
	"bufio"
	"net"
	"strings"
)

func GetName(conn net.Conn) (string, bool) {
	pingv := Showpinguin()
	conn.Write([]byte(pingv))
	conn.Write([]byte("[ENTER YOUR NAME:] "))
	input := bufio.NewScanner(conn)
	var name string
	if input.Scan() {
		text := strings.TrimSpace(input.Text())
		name = text
	}
	return name, CheckName(name, conn)
}

func CheckName(username string, conn net.Conn) bool {
	if strings.TrimSpace(username) == "" {
		conn.Write([]byte("You cannot enter empty text"))
		return true
	}
	Mu.Lock()
	for name := range Clients {
		if name == username {
			conn.Write([]byte("You cannot enter empty text'"))
			Mu.Unlock()
			return true
		}
	}
	Mu.Unlock()
	return false
}

func ClearLine(s string) string {
	return "\r" + strings.Repeat(" ", len(s)+len(s)+10) + "\r"
}

func Checktext(s string) string {
	return strings.TrimSpace(s)
}

func LastCheckText(s string) bool {
	if s == "" {
		return false
	}
	for _, w := range s {
		if (w >= 65 && w <= 90) || (w >= 97 && w <= 122) {
			return true
		} else {
			return false
		}
	}

	return true
}
