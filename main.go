package main

import (
	"encoding/hex"
	"flag"
	"fmt"
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

	v := r1.Intn(255)
	fmt.Println("random value: ", v)

	client := modbus.TCPClient(*serverUri)

	results, e := client.ReadInputRegisters(8, 1)
	if e != nil {
		panic(e.Error())
	}
	println(hex.EncodeToString(results))

	results, e = client.WriteSingleRegister(8, uint16(v))
	if e != nil {
		panic(e.Error())
	}
	println(hex.EncodeToString(results))

	results, e = client.ReadInputRegisters(8, 1)
	if e != nil {
		panic(e.Error())
	}
	println(hex.EncodeToString(results))
}
