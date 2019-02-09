package parking

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetSlot(t *testing.T) {
	vehicle := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2839"}
	SetMAxSlot(5)
	slot, err := vehicle.GetSlot()
	if err != nil && strings.Compare(err.Error(), "slots are full") == 0 {

	}
	t.Log(slot)
	if slot.VehicleColor != vehicle.VehicleColor {
		t.Error("failed to create slot")
	}
	if slot.VehicleColor != vehicle.VehicleColor {
		t.Error("slot not created successfully")
	}
}
func TestSlotsFull(t *testing.T) {
	vehicle := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2839"}
	vehicle1 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2832"}
	SetMAxSlot(1)
	_, err := vehicle.GetSlot()
	if err != nil {
		t.Error("failed to allocate slot")
	}
	_, err = vehicle1.GetSlot()
	if err != nil && strings.Compare(err.Error(), "slots are full") == 0 {
		t.Log(err.Error())
	}
}

func TestExit(t *testing.T) {
	vehicle := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2839"}
	SetMAxSlot(5)
	slot, err := vehicle.GetSlot()
	if err != nil {
		t.Error("failed to allocate slot")
	}
	isLeft := slot.Exit()
	if isLeft != nil && *isLeft != true {
		t.Error("failed to clear slot")
	}
	fmt.Println(GetFreeSlots())
	slots := GetVehicleByProps(vehicle.VehicleColor)
	fmt.Println(fmt.Sprintf("%+v", slots))
}

func TestGetVehicleByProps(t *testing.T) {
	vehicle := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2839"}
	vehicle1 := &vehicleInfo{VehicleColor: "red", VehicleRegNO: "TN AH 2887"}
	vehicle2 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2898"}
	SetMAxSlot(5)
	vehicle.GetSlot()
	vehicle1.GetSlot()
	vehicle2.GetSlot()
	slots := GetVehicleByProps(vehicle2.VehicleColor)
	if slots == nil {
		t.Error("slots by color ", slots)
	}
	t.Log(GetStorageMap())
	fmt.Println(slots)
}

func TestMaxSlot(t *testing.T) {
	vehicle := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2839"}
	vehicle1 := &vehicleInfo{VehicleColor: "red", VehicleRegNO: "TN AH 2887"}
	vehicle2 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2898"}
	vehicle3 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AA 2898"}
	vehicle4 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AB 2898"}
	vehicle5 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AC 2898"}
	SetMAxSlot(5)
	vehicle.GetSlot()
	vehicle1.GetSlot()
	vehicle2.GetSlot()
	vehicle3.GetSlot()
	vehicle4.GetSlot()
	slot, err := vehicle5.GetSlot()
	if slot == nil && err != nil {
		t.Log("failed to allocate slot")
	}
	if slot != nil {
		t.Error("should not exced the maximum slot")
	}
	fmt.Println(GetSlotStatus())
	if len(GetSlotStatus()) > 5 || len(GetSlotStatus()) > 5 {
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
	SetMAxSlot(5)
	vehicle.GetSlot()
	slot1, err := vehicle1.GetSlot()
	if err != nil {
		t.Error("failed to allocate slot")
	}
	isLeft := slot1.Exit()
	if isLeft == nil && *isLeft != true {
		t.Error("failed to free the slot")
	}
	slot2, err := vehicle2.GetSlot()
	if err != nil {
		t.Error("failed to allocate slot")
	}
	if slot2.SlotID != slot1.SlotID {
		t.Error("allocating wrong slot")
	}
	vehicle3.GetSlot()
	vehicle4.GetSlot()
	slot, err := vehicle5.GetSlot()
	if slot == nil && err != nil {
		t.Log("failed to allocate slot")
	}
	fmt.Println(GetSlotStatus())
	if len(GetSlotStatus()) > 5 || len(GetSlotStatus()) > 5 {
		t.Error("exceding max slot")
	}
}

func TestDuplicateAlocation(t *testing.T) {
	vehicle := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2839"}
	vehicle1 := &vehicleInfo{VehicleColor: "red", VehicleRegNO: "TN AH 2887"}
	vehicle2 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AH 2898"}
	vehicle3 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AA 2898"}
	vehicle4 := &vehicleInfo{VehicleColor: "white", VehicleRegNO: "TN AA 2898"}
	SetMAxSlot(5)
	vehicle.GetSlot()
	vehicle1.GetSlot()
	vehicle2.GetSlot()
	vehicle3.GetSlot()
	vehicle4.GetSlot()
	fmt.Println(GetSlotStatus())
	if len(GetSlotStatus()) == 5 {
		t.Error("exceding max slot")
	}
}
