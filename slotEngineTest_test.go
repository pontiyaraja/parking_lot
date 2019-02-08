package parking_lot

import (
	"fmt"
	"testing"
)

func TestGetSlot(t *testing.T) {
	vehicle := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2839"}
	setMAxSlot(5)
	slot := vehicle.GetSlot()
	fmt.Println(slot)
	if slot.VehicleColor != vehicle.VehicleColor {
		t.Error("failed to create slot")
	}
	if slot.VehicleColor != vehicle.VehicleColor {
		t.Error("slot not created successfully")
	}
}

func TestExit(t *testing.T) {
	vehicle := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2839"}
	setMAxSlot(5)
	slot := vehicle.GetSlot()
	isLeft := slot.Exit()
	if isLeft != nil && *isLeft != true {
		t.Error("failed to clear slot")
	}
	fmt.Println(getFreeSlots())
	slots := getVehicleByProps(vehicle.VehicleColor)
	fmt.Println(fmt.Sprintf("%+v", slots))
}

func TestGetVehicleByProps(t *testing.T) {
	vehicle := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2839"}
	vehicle1 := &vehicleInfo{VehicleColor: "red", VehicleRegNO: "TN AH 2887"}
	vehicle2 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2898"}
	setMAxSlot(5)
	vehicle.GetSlot()
	vehicle1.GetSlot()
	vehicle2.GetSlot()
	slots := getVehicleByProps(vehicle2.VehicleColor)
	if slots == nil {
		t.Error("slots by color ", slots)
	}
	t.Log(getStorageMap())
	fmt.Println(slots)
}

func TestMaxSlot(t *testing.T) {
	vehicle := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2839"}
	vehicle1 := &vehicleInfo{VehicleColor: "red", VehicleRegNO: "TN AH 2887"}
	vehicle2 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2898"}
	vehicle3 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AA 2898"}
	vehicle4 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AB 2898"}
	vehicle5 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AC 2898"}
	setMAxSlot(5)
	vehicle.GetSlot()
	vehicle1.GetSlot()
	vehicle2.GetSlot()
	vehicle3.GetSlot()
	vehicle4.GetSlot()
	slot := vehicle5.GetSlot()
	if slot != nil {
		t.Error("should not exced the maximum slot")
	}
	fmt.Println(getSlotStatus())
	if len(getSlotStatus()) > 5 || len(getSlotStatus()) > 5 {
		t.Error("exceding max slot")
	}
}

func TestFreeSlotAllocation(t *testing.T) {
	vehicle := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2839"}
	vehicle1 := &vehicleInfo{VehicleColor: "red", VehicleRegNO: "TN AH 2887"}
	vehicle2 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2898"}
	vehicle3 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AA 2898"}
	vehicle4 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AB 2898"}
	vehicle5 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AC 2898"}
	setMAxSlot(5)
	vehicle.GetSlot()
	slot1 := vehicle1.GetSlot()
	isLeft := slot1.Exit()
	if isLeft == nil && *isLeft != true {
		t.Error("failed to free the slot")
	}
	slot2 := vehicle2.GetSlot()
	if slot2.SlotID != slot1.SlotID {
		t.Error("allocating wrong slot")
	}
	vehicle3.GetSlot()
	vehicle4.GetSlot()
	slot := vehicle5.GetSlot()
	if slot == nil {
		t.Error("failed to allocate slot")
	}
	fmt.Println(getSlotStatus())
	if len(getSlotStatus()) > 5 || len(getSlotStatus()) > 5 {
		t.Error("exceding max slot")
	}
}

func TestDuplicateAlocation(t *testing.T) {
	vehicle := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2839"}
	vehicle1 := &vehicleInfo{VehicleColor: "red", VehicleRegNO: "TN AH 2887"}
	vehicle2 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2898"}
	vehicle3 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AA 2898"}
	vehicle4 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AA 2898"}
	setMAxSlot(5)
	vehicle.GetSlot()
	vehicle1.GetSlot()
	vehicle2.GetSlot()
	vehicle3.GetSlot()
	vehicle4.GetSlot()
	fmt.Println(getSlotStatus())
	if len(getSlotStatus()) == 5 {
		t.Error("exceding max slot")
	}
}
