package models

type AppState struct {
	DeviceID          string
	PairedDevicesInfo []ArlinPairedDeviceInfo
	LastConnected     uint64
}
