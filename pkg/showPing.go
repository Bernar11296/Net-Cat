package pkg

import (
	"bufio"
	"log"
	"os"
)

func Showpinguin() string {
	f, err := os.Open("pinguin.txt")
	if err != nil {
		log.Fatal(err)
	}
	pingv := ""
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pingv += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return pingv
}
