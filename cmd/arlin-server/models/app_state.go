package models

type AppState struct {
	DeviceID         string
	PairedDeviceInfo ArlinDeviceInfo
	LastConnected    uint64
}
