package main

import (
	"encoding/hex"
	"flag"
	"github.com/goburrow/modbus"
	"math/rand"
	"time"
)

func main() {
	serverUri := flag.String("h", "localhost:502", "uri of modbus server")
	flag.Parse()
	if *serverUri == "" {
		flag.Usage()
		return
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	client := modbus.TCPClient(*serverUri)

	results, _ := client.ReadInputRegisters(8, 1)
	println(hex.EncodeToString(results))

	v := r1.Intn(255)
	results, _ = client.WriteSingleRegister(8, uint16(v))
	println(hex.EncodeToString(results))

	results, _ = client.ReadInputRegisters(8, 1)
	println(hex.EncodeToString(results))
}
