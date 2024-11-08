package appstate

import (
	"encoding/gob"
	"errors"
	"fmt"

	"os"
	"path"

	"shazin.me/arlin/cmd/arlin-server/models"
	"shazin.me/arlin/cmd/arlin-server/utils"
)

var appStateDir = ".local/arlin"
var appStateFileName = "appstate.dat"

var homeDir, _ = os.UserHomeDir()

var appStateDirAbs = path.Join(homeDir, appStateDir)
var appStatFileAbs = path.Join(appStateDirAbs, appStateFileName)

func InitAppStats() error {

	fmt.Println("Abs path ", appStatFileAbs)

	// Check whether path exists
	_, error := os.Stat(appStateDir)

	// create file if not exist
	if errors.Is(error, os.ErrNotExist) {
		// Path doesn't exist
		error = os.MkdirAll(appStateDirAbs, 0755)

		if error != nil {
			return error
		}

	}

	// directory is ensured
	// ensure path
	_, error = os.Stat(appStatFileAbs)

	if errors.Is(error, os.ErrNotExist) {
		// file path doesn't exist
		appStateFile, error := os.Create(appStatFileAbs)
		if error != nil {
			return error
		}
		defer appStateFile.Close()
		// since  file is being created for the first time we generate a device id and save the default state
		deviceID, error := utils.GenerateDeviceID(10)
		if error != nil {
			return error
		}
		appState_ := models.AppState{
			DeviceID: deviceID,
		}
		gobEnocder := gob.NewEncoder(appStateFile)
		if error := gobEnocder.Encode(appState_); error != nil {
			return error
		}

	}

	// appStatePath is ensured

	// try to decode file
	var appState models.AppState
	if error = utils.LoadFromFile(appStatFileAbs, &appState); error != nil {
		fmt.Println("Error loading:", error)
		return error
	}

	fmt.Println("app state initialized")

	return nil

}

// function to save the app state
func saveAppState(appState models.AppState) error {
	error := utils.SaveToFile(appStatFileAbs, appState)
	if error != nil {
		return error
	}
	return nil
}

func GetAppState() (models.AppState, error) {
	var appState models.AppState
	error := utils.LoadFromFile(appStatFileAbs, &appState)

	if error != nil {
		return models.AppState{}, error
	}

	return appState, nil

}

// an utility function to check whether device with that id exists
func IsDevicePaired(deviceID string) (bool, int) {
	appState, error := GetAppState()
	if error != nil {
		return false, -1
	}

	for index, device := range appState.PairedDevicesInfo {
		if device.DeviceID == deviceID {
			return true, index
		}
	}
	return false, -1
}

func AddPairedDevice(device models.ArlinPairedDeviceInfo) error {

	devicePaired, _ := IsDevicePaired(device.DeviceID)

	// check whether device is already paired
	if devicePaired {
		return errors.New("Cannot pair device, Device is already paired")
	}

	appState, error := GetAppState()
	if error != nil {
		return error
	}

	appState.PairedDevicesInfo = append(appState.PairedDevicesInfo, device)

	// fmt.Println("Updated app state: ", oldAppState)

	error = utils.SaveToFile(appStatFileAbs, appState)

	// if error != nil {
	// 	return error
	// }
	return nil
}

func UnpairDevice(deviceID string) error {

	devicePaired, pairingIndex := IsDevicePaired(deviceID)

	// check whether device is already paired
	if !devicePaired {
		return errors.New("Couldn't unpair device. Device doesn;t exists in the paired list!")
	}

	appState, error := GetAppState()
	if error != nil {
		return error
	}

	// delete the element from the slice
	updatedPairedList := append(
		appState.PairedDevicesInfo[:pairingIndex], appState.PairedDevicesInfo[pairingIndex+1:]...,
	)

	updatedAppState := models.AppState{
		DeviceID:          appState.DeviceID,
		PairedDevicesInfo: updatedPairedList,
		LastConnected:     appState.LastConnected,
	}

	error = utils.SaveToFile(appStatFileAbs, updatedAppState)

	return nil
}
