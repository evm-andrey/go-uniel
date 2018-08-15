package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jacobsa/go-serial/serial"
)

func main() {
	logger := log.New(os.Stdout, "", 0) // log.Logger - потоково-безопасный тип для вывода
	logger.Println("API server for UNIEL")
	txdata := string("FFFF0B010000000C")
	reley := map[int]int{
		1: 12,
	}
	logger.Println(reley)

	options := serial.OpenOptions{
		PortName:               "/dev/ttyUSB0",
		BaudRate:               9600,
		DataBits:               8,
		StopBits:               1,
		MinimumReadSize:        0,
		InterCharacterTimeout:  100, // ms
		ParityMode:             serial.PARITY_NONE,
		Rs485Enable:            false,
		Rs485RtsHighDuringSend: false,
		Rs485RtsHighAfterSend:  false,
	}

	f, err := serial.Open(options)

	if err != nil {
		logger.Println("Error opening serial port: ", err)
		os.Exit(-1)
	} else {
		defer f.Close()
	}

	txdata1, err := hex.DecodeString(txdata)

	if err != nil {
		logger.Println("Error decoding hex data: ", err)
		os.Exit(-1)
	}

	logger.Println("Sending: ", hex.EncodeToString(txdata1))

	count, err := f.Write(txdata1)

	if err != nil {
		logger.Println("Error writing to serial port: ", err)
	} else {
		logger.Printf("Wrote %v bytes\n", count)
	}

	for {
		buf := make([]byte, 32)
		n, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from serial port: ", err)
			}
		} else {
			buf = buf[:n]
			fmt.Println("Rx: ", hex.EncodeToString(buf))
		}
	}
}
