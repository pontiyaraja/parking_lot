package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	parking "github.com/pontiyaraja/parking_lot/parkinglot"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		dataString, err := reader.ReadString('\n')
		fmt.Println(dataString, err)
		dataString = strings.TrimSuffix(dataString, "\n")
		if err != nil {
			os.Exit(1)
		}
		dataArray := strings.Split(dataString, " ")
		switch dataArray[0] {
		case "create_parking_lot":
			fmt.Println("allocating lot ............")
			maxSlot, err := strconv.Atoi(dataArray[1])
			if err != nil {
				fmt.Println("failed to parse convert parking lot", err, dataArray[1])
				break
			}
			parking.SetMAxSlot(maxSlot)
			break

		case "park":
			fmt.Println(" Parking ............")
			vehicle := parking.Framework{}.VehicleInfo
			vehicle.VehicleRegNO = dataArray[1]
			vehicle.VehicleColor = dataArray[2]
			slot := vehicle.GetSlot()
			fmt.Println("booked slot ", slot)
			slot = nil
			fmt.Println("Status   :    ", parking.GetSlotStatus())
			break

		case "leave":
			fmt.Println("leaving.......")
			slotID, err := strconv.Atoi(dataArray[1])
			if err != nil {
				fmt.Println("failed to parse convert parking lot", err, dataArray[1])
				break
			}
			slotInfo := parking.GetSlotBySlotID(slotID)
			if slotInfo == nil {
				fmt.Println("vehicle not found")
				break
			}
			if slotInfo.SlotID == slotID {
				exited := slotInfo.Exit()
				if exited != nil && *exited == false {
					fmt.Println("failed to exit")
					break
				}
				break
			} else {
				fmt.Println("failed to get vehicle")
				break
			}

		case "status":
			slots := parking.GetSlotStatus()
			fmt.Println("Slot No.    Registration No        Colour")
			for _, slot := range slots {
				fmt.Println(fmt.Sprintf("%d           %s          %s", slot.SlotID, slot.VehicleRegNO, slot.VehicleColor))
			}
			break

		case "exit":
			os.Exit(0)
		}
	}
}
