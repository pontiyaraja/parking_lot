package parking_lot

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

type vehicleInfo struct {
	VehicleRegNO string
	VehicleColor string
}

type framework struct {
	SlotInterface slotInfo
}

type slotInfo struct {
	SlotID int
	Ticket string
	*vehicleInfo
}

type slotInterface interface {
	GetSlot() slotInfo
	Exit()
}
type slotStorage struct {
	lck            sync.RWMutex
	slotMap        map[int]slotInfo
	slotStorageMap map[string][]slotInfo
	slotCount      int
}

var slotMap map[string]slotInfo
var slotStorageMap map[string][]slotInfo
var freeSlotMap map[int]bool
var slotCount int
var maxSlotCount int

func init() {
	slotMap = make(map[string]slotInfo)
	slotStorageMap = make(map[string][]slotInfo)
	freeSlotMap = make(map[int]bool)
}

func (vehicle *vehicleInfo) GetSlot() *slotInfo {
	freeSlots := getFreeSlots()
	slot := slotInfo{vehicleInfo: vehicle}
	if freeSlots != nil {
		slot.SlotID = freeSlots[0]
		delete(freeSlotMap, freeSlots[0])
	} else if slotCount < maxSlotCount {
		slotCount++
		slot.SlotID = slotCount
	} else {
		fmt.Println(errors.New(" slots are full "))
		return nil
	}
	slot.Ticket = "random uuid"
	_, ok := slotMap[slot.VehicleRegNO]
	if !ok {
		slotMap[slot.VehicleRegNO] = slot
		slotStorageMap[vehicle.VehicleColor] = append(slotStorageMap[vehicle.VehicleColor], slot)
		return &slot
	}
	return nil
}

func (slot *slotInfo) Exit() *bool {
	//remove the slot at the position slotID
	delete(slotMap, slot.VehicleRegNO)
	return slot.RemoveVehicle()
}

func (slot *slotInfo) RemoveVehicle() *bool {
	slots := getVehicleByProps(slot.VehicleColor)
	var slotArray []slotInfo
	if slots != nil {
		for _, slotData := range slots {
			if slotData.SlotID != slot.SlotID {
				slotArray = append(slotArray, slotData)
			}
		}
		slotStorageMap[slot.VehicleColor] = slotArray
		freeSlotMap[slot.SlotID] = true
		return func(resp bool) *bool {
			return &resp
		}(true)
	}
	return nil
}

func getFreeSlots() []int {
	var freeSlot []int
	for key := range freeSlotMap {
		// if freeSlot == 0 {
		// 	freeSlot = key
		// } else if freeSlot > key {
		// 	freeSlot = key
		// }
		freeSlot = append(freeSlot, key)
	}
	sort.Slice(freeSlot, func(i, j int) bool { return i < j })
	return freeSlot
}

func setMAxSlot(mxSlot int) {
	maxSlotCount = mxSlot
}

func getStorageMap() map[string][]slotInfo {
	return slotStorageMap
}

func getVehicleByProps(prop string) []slotInfo {
	val, ok := slotStorageMap[prop]
	if ok {
		return val
	}
	return nil
}

func getSlotStatus() []slotInfo {
	var slotStatus []slotInfo
	for _, val := range slotMap {
		slotStatus = append(slotStatus, val)
	}
	sort.Slice(slotStatus, func(i, j int) bool { return slotStatus[i].SlotID < slotStatus[j].SlotID })
	return slotStatus
}
