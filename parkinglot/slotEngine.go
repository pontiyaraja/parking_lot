package parking

import (
	"errors"
	"sort"
	"sync"
)

type vehicleInfo struct {
	VehicleRegNO string
	VehicleColor string
}

type Framework struct {
	SlotInfo    slotInfo
	VehicleInfo vehicleInfo
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

var slotMap map[int]slotInfo
var slotStorageMap map[string][]slotInfo
var freeSlotMap map[int]bool
var slotCount int
var maxSlotCount int

func init() {
	slotMap = make(map[int]slotInfo)
	slotStorageMap = make(map[string][]slotInfo)
	freeSlotMap = make(map[int]bool)
}

func (vehicle *vehicleInfo) GetSlot() (*slotInfo, error) {
	if vehicleParked(vehicle.VehicleRegNO) {
		return nil, errors.New("already parked")
	}
	freeSlots := GetFreeSlots()
	slot := slotInfo{vehicleInfo: vehicle}
	if freeSlots != nil {
		slot.SlotID = freeSlots[0]
		delete(freeSlotMap, freeSlots[0])
	} else if slotCount < maxSlotCount {
		slotCount++
		slot.SlotID = slotCount
	} else {
		return nil, errors.New("slots are full")
	}
	slot.Ticket = "random uuid"
	slotMap[slot.SlotID] = slot
	slotStorageMap[vehicle.VehicleColor] = append(slotStorageMap[vehicle.VehicleColor], slot)
	slotStorageMap[vehicle.VehicleRegNO] = []slotInfo{slot}
	return &slot, nil
}

func vehicleParked(regNo string) bool {
	_, ok := slotStorageMap[regNo]
	return ok
}

func (slot *slotInfo) Exit() *bool {
	//remove the slot at the position slotID
	delete(slotMap, slot.SlotID)
	isRemoved := slot.RemoveVehicle()
	slot = nil
	return isRemoved
}

func (slot *slotInfo) RemoveVehicle() *bool {
	slots := GetVehicleByProps(slot.VehicleColor)
	var slotArray []slotInfo
	if slots != nil {
		for _, slotData := range slots {
			if slotData.SlotID != slot.SlotID {
				slotArray = append(slotArray, slotData)
			}
		}
		slotStorageMap[slot.VehicleColor] = slotArray
		freeSlotMap[slot.SlotID] = true
		delete(slotStorageMap, slot.VehicleRegNO)
		return func(resp bool) *bool {
			return &resp
		}(true)
	}
	return nil
}

func GetFreeSlots() []int {
	var freeSlot []int
	for key := range freeSlotMap {
		// if freeSlot == 0 {
		// 	freeSlot = key
		// } else if freeSlot > key {
		// 	freeSlot = key
		// }
		freeSlot = append(freeSlot, key)
	}
	sort.Slice(freeSlot, func(i, j int) bool { return freeSlot[i] < freeSlot[j] })
	return freeSlot
}

func SetMAxSlot(mxSlot int) {
	maxSlotCount = mxSlot
	clearSlots()
}

func GetStorageMap() map[string][]slotInfo {
	return slotStorageMap
}

func GetVehicleByProps(prop string) []slotInfo {
	val, ok := slotStorageMap[prop]
	if ok {
		return val
	}
	return nil
}

func GetSlotStatus() []slotInfo {
	var slotStatus []slotInfo
	for _, val := range slotMap {
		slotStatus = append(slotStatus, val)
	}
	sort.Slice(slotStatus, func(i, j int) bool { return slotStatus[i].SlotID < slotStatus[j].SlotID })
	return slotStatus
}

func GetSlotBySlotID(slotID int) *slotInfo {
	val, ok := slotMap[slotID]
	if ok {
		return &val
	}
	return nil
}

func clearSlots() {
	slotMap = nil
	slotStorageMap = nil
	freeSlotMap = nil
	slotCount = 0
	slotMap = make(map[int]slotInfo)
	slotStorageMap = make(map[string][]slotInfo)
	freeSlotMap = make(map[int]bool)
}
