package parking_lot

type slotStorage interface {
	GetVehicleByColor(string) slotInfo
}

func GetVehicleByColor(color string) slotInfo {

}
