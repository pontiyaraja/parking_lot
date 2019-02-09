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

var slotMap *sync.Map
var slotStorageMap *sync.Map
var freeSlotMap *sync.Map
var slotCount int
var maxSlotCount int

func init() {
	slotMap = nil
	slotStorageMap = nil
	freeSlotMap = nil

	slotMap = new(sync.Map)
	slotStorageMap = new(sync.Map)
	freeSlotMap = new(sync.Map)
}

func (vehicle *vehicleInfo) GetSlot() (*slotInfo, error) {
	if vehicleParked(vehicle.VehicleRegNO) {
		return nil, errors.New("already parked")
	}
	freeSlots := GetFreeSlots()
	slot := slotInfo{vehicleInfo: vehicle}
	if freeSlots != nil {
		slot.SlotID = freeSlots[0]
		freeSlotMap.Delete(freeSlots[0])
	} else if slotCount < maxSlotCount {
		slotCount++
		slot.SlotID = slotCount
	} else {
		return nil, errors.New("slots are full")
	}
	slot.Ticket = "random uuid"
	slotMap.Store(slot.SlotID, slot)
	val, ok := slotStorageMap.Load(vehicle.VehicleColor)
	if !ok {
		slotStorageMap.Store(vehicle.VehicleColor, []slotInfo{slot})
	} else {
		slotStorageMap.Store(vehicle.VehicleColor, append(val.([]slotInfo), slot))
	}
	slotStorageMap.Store(vehicle.VehicleRegNO, []slotInfo{slot})
	return &slot, nil
}

func vehicleParked(regNo string) bool {
	_, ok := slotStorageMap.Load(regNo)
	return ok
}

func (slot *slotInfo) Exit() *bool {
	slotMap.Delete(slot.SlotID)
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

		slotStorageMap.Store(slot.VehicleColor, slotArray)
		freeSlotMap.Store(slot.SlotID, true)
		slotStorageMap.Delete(slot.VehicleRegNO)

		return func(resp bool) *bool {
			return &resp
		}(true)
	}
	return nil
}

func GetFreeSlots() []int {
	var freeSlot []int
	freeSlotMap.Range(func(key, value interface{}) bool {
		freeSlot = append(freeSlot, key.(int))
		return true
	})

	sort.Slice(freeSlot, func(i, j int) bool { return freeSlot[i] < freeSlot[j] })
	return freeSlot
}

func SetMAxSlot(mxSlot int) {
	maxSlotCount = mxSlot
	clearSlots()
}

func GetStorageMap() *sync.Map {
	return slotStorageMap
}

func GetVehicleByProps(prop string) []slotInfo {
	val, ok := slotStorageMap.Load(prop)
	if ok {
		return val.([]slotInfo)
	}
	return nil
}

func GetSlotStatus() []slotInfo {
	var slotStatus []slotInfo
	slotMap.Range(func(key, value interface{}) bool {
		slotStatus = append(slotStatus, value.(slotInfo))
		return true
	})
	sort.Slice(slotStatus, func(i, j int) bool { return slotStatus[i].SlotID < slotStatus[j].SlotID })
	return slotStatus
}

func GetSlotBySlotID(slotID int) *slotInfo {
	val, ok := slotMap.Load(slotID)
	if ok {
		return func(slot slotInfo) *slotInfo {
			return &slot
		}(val.(slotInfo))
	}
	return nil
}

func clearSlots() {
	slotMap.Range(func(key, value interface{}) bool {
		slotMap.Delete(key)
		return true
	})
	slotStorageMap.Range(func(key, value interface{}) bool {
		slotStorageMap.Delete(key)
		return true
	})
	freeSlotMap.Range(func(key, value interface{}) bool {
		freeSlotMap.Delete(key)
		return true
	})
	slotMap = nil
	slotStorageMap = nil
	freeSlotMap = nil
	slotCount = 0

	slotMap = new(sync.Map)
	slotStorageMap = new(sync.Map)
	freeSlotMap = new(sync.Map)
}
