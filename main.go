package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/kgolding/go-temp-test/rutos"
	"github.com/kgolding/go-temp-test/scope"
)

func main() {
	var port, tel string

	flag.StringVar(&port, "port", "/dev/rs232", "serial port device")
	flag.StringVar(&tel, "tel", "", "telephone number to send SMS too")

	flag.Parse()

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

		if tel != "" {
			log.Printf("Sending SMS to %s: '%s'\n", tel, m.Message)
			err = rutos.SendSMS(tel, m.Message)
			if err != nil {
				log.Printf("Error sending SMS to %s: '%s': %s\n", tel, m.Message, err.Error())
			}
		}
	}
	log.Println("Port closed")
}
