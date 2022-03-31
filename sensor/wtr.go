package wtr10e

import (
	"fmt"
	"log"
	"time"
	"wtr10e/utils"

	"github.com/goburrow/modbus"
)

func WTR10E(port string, baudrate int, slaveid byte, timeout time.Duration) (results [2]string, err error) {
	handler := modbus.NewRTUClientHandler(port)
	handler.BaudRate = baudrate
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = slaveid
	handler.Timeout = timeout * time.Millisecond
	//defer handler.Close()

	err = handler.Connect()
	if err != nil {
		log.Println("port sensor busy")
		return results, err
	} else {
		//log.Println("port sensor conn")

		client := modbus.NewClient(handler)
		data, err := client.ReadHoldingRegisters(0, 2)
		if err != nil {
			log.Println("error timeout")
			return results, err
		} else {
			//hex := fmt.Sprintf("%x", data)
			//log.Printf("hex : %s", hex)

			buff := utils.BytesToUint16s(utils.BIG_ENDIAN, data)
			//log.Printf("buff : %v", buff)

			var buffert [2]float32
			buffert[0] = float32(buff[0]) / 10
			buffert[1] = float32(buff[1]) / 10

			results[0] = fmt.Sprintf("%.2f", buffert[0])
			results[1] = fmt.Sprintf("%.2f", buffert[1])

		}
	}
	handler.Close()

	return results, err
}
