package parking_lot

type vehicleInfo struct {
	VehicleRegNO string
	VehicleColor string
}

type slotInfo struct {
	SlotID int
	*vehicleInfo
}

type slotInterface interface {
	GetSlot() slotInfo
	Exit()
}

var slotMap map[int]slotInfo
var slotStorageMap map[string][]slotInfo
var slotCount int

func init() {
	slotMap = make(map[int]slotInfo)
	slotStorageMap = make(map[string][]slotInfo)
}

func (vehicle *vehicleInfo) GetSlot() slotInfo {
	slotCount++
	return slotInfo{vehicleInfo: vehicle, SlotID: slotCount}
}

func (slot slotInfo) Exit() bool {
	//remove the slot at the position slotID
	return true
}
