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
	var scanner *bufio.Scanner
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("failed to open file")
		}
		scanner = bufio.NewScanner(file)
	}

	reader := bufio.NewReader(os.Stdin)

	var dataString string
	var err error
	isScannerEmpty := false
	for {
		if scanner != nil && scanner.Scan() {
			dataString = scanner.Text()
		} else {
			if !isScannerEmpty {
				scanner = nil
			}
			dataString, err = reader.ReadString('\n')
			if err != nil {
				fmt.Println("wrong input please follow the instructions")
			}
			dataString = strings.TrimSuffix(dataString, "\n")
		}
		dataArray := strings.Split(dataString, " ")

		switch dataArray[0] {
		case "create_parking_lot":
			allocateSlot(dataArray[1])
			break

		case "park":
			parkTheVehicle(dataArray[1], dataArray[2])
			break

		case "leave":
			clearTheslot(dataArray[1])
			break

		case "status":
			showThelotStatus()
			break

		case "registration_numbers_for_cars_with_colour":
			getRegNOByVehicleProps(dataArray[1], "regNOByColor")
			break

		case "slot_numbers_for_cars_with_colour":
			getRegNOByVehicleProps(dataArray[1], "slotNOByColor")
			break

		case "slot_number_for_registration_number":
			getRegNOByVehicleProps(dataArray[1], "slotNOByRegNO")
			break

		case "exit":
			os.Exit(0)
		}
	}
}

func allocateSlot(input string) {
	maxSlot, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("failed to parse convert parking lot", err, input)
		return
	}
	parking.SetMAxSlot(maxSlot)
	fmt.Println(fmt.Sprintf("Created a parking lot with %d slots", maxSlot))
	return
}

func parkTheVehicle(regNO, color string) {
	vehicle := parking.Framework{}.VehicleInfo
	vehicle.VehicleRegNO = regNO
	vehicle.VehicleColor = color
	slot, err := vehicle.GetSlot()
	if err != nil {
		if err.Error() == "already parked" {
			fmt.Println("Vehicle already parked")
		} else if err.Error() == "slots are full" {
			fmt.Println("Sorry, parking lot is full")
		}
		return
	}
	fmt.Println("Allocated slot number:", slot.SlotID)
	slot = nil
	return
}

func clearTheslot(parkedSlotID string) {
	slotID, err := strconv.Atoi(parkedSlotID)
	if err != nil {
		fmt.Println("failed to parse convert parking lot", err, parkedSlotID)
		return
	}
	slotInfo := parking.GetSlotBySlotID(slotID)
	if slotInfo == nil {
		fmt.Println("vehicle not found")
		return
	}
	if slotInfo.SlotID == slotID {
		exited := slotInfo.Exit()
		if exited != nil && *exited == false {
			fmt.Println("failed to exit")
			return
		}
		fmt.Println(fmt.Sprintf("Slot number %d is free", slotInfo.SlotID))
		return
	}
	fmt.Println("failed to get vehicle")
	return
}

func showThelotStatus() {
	slots := parking.GetSlotStatus()
	fmt.Println(fmt.Sprintf("%-12s%-19s%-7s", "Slot No.", "Registration No", "Colour"))
	for _, slot := range slots {
		fmt.Println(fmt.Sprintf("%-12d%-19s%-7s", slot.SlotID, slot.VehicleRegNO, slot.VehicleColor))
	}
}

func getRegNOByVehicleProps(color, props string) {
	var resString string
	slots := parking.GetVehicleByProps(color)
	if len(slots) > 0 {
		if strings.Compare(props, "slotNOByRegNO") == 0 {
			fmt.Println(slots[0].SlotID)
			return
		}
		for index, slot := range slots {
			if index > 0 {
				resString = resString + ", "
			}
			if strings.Compare(props, "regNOByColor") == 0 {
				resString = resString + slot.VehicleRegNO
			} else if strings.Compare(props, "slotNOByColor") == 0 {
				resString = resString + strconv.Itoa(slot.SlotID)
			}
		}
		fmt.Print(resString + "\n")
	} else {
		fmt.Println("Not found")
	}
}
