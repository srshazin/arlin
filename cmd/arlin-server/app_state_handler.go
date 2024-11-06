package arlinserver

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

func GetAppState() (models.AppState, error) {
	var appState models.AppState
	error := utils.LoadFromFile(appStatFileAbs, &appState)

	if error != nil {
		return models.AppState{}, error
	}

	return appState, nil

}

// func AddPairedDevice(device models.ArlinPairedDeviceInfo) error {
// 	error := utils.SaveToFile(appStatFileAbs, appState)

// 	if error != nil {
// 		return error
// 	}
// 	return nil
// }
