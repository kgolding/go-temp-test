package main

import (
	"bufio"
	"hello/scope"
	"log"
	"os"
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
	}
	log.Println("Port closed")
}
