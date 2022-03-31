package main

import (
	"log"
	sensor "wtr10e/sensor"
)

func main() {
	data, err := sensor.WTR10E("/dev/ttyUSB0", 19200, 1, 300)
	if err != nil {
		log.Println("error get data")
	} else {
		log.Println("succ get data")
		log.Printf("temp : %v C", data[0])
		log.Printf("rh   : %v %%", data[1])
	}
}
