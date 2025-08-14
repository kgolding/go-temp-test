package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kgolding/go-temp-test/scope"
)

func main() {
	if len(os.Args) < 2 {
		panic("missing arg")
	}

	port := os.Args[1]

	f, err := os.Open(port)
	if err != nil {
		log.Printf("Error opening port '%s': %s", port, err)
		return
	}
	log.Printf("Opened port '%s'", port)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Bytes()
		log.Printf("RX: '%s'\n", line)
		m, err := scope.Decode(append(line, 0x0d))

		if err != nil {
			log.Printf("Error: %s\n", err)
			continue
		}

		log.Println("Decoded: ", m.String())

		// Send OK
		f.Write([]byte{0x4F, 0x4B, 0x0D, 0x0A, 0x3E}) // "OK\r\n>"
	}
	log.Println("Port closed")
}
