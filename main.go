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
	address := flag.Int("addr", 8, "modbus address to write")
	value := flag.Int("val", -1, "modbus value to write")
	flag.Parse()
	if *serverUri == "" {
		flag.Usage()
		return
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	if *value < 0 {
		*value = r1.Intn(255)
		fmt.Println("random value: ", *value)
	}

	client := modbus.TCPClient(*serverUri)

	results, e := client.ReadInputRegisters(uint16(*address), 1)
	if e != nil {
		panic(e.Error())
	}
	fmt.Println("original: ", hex.EncodeToString(results))

	results, e = client.WriteSingleRegister(uint16(*address), uint16(*value))
	if e != nil {
		panic(e.Error())
	}
	fmt.Println("written: ", hex.EncodeToString(results))

	results, e = client.ReadInputRegisters(uint16(*address), 1)
	if e != nil {
		panic(e.Error())
	}
	fmt.Println("updates: ", hex.EncodeToString(results))
}
