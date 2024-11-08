package command

// command checks if device is paired then replies with a message that contains device information

type inqData struct {
	DeviceID    string `json:"deviceID"`
	HostName    string `json:"hostName"`
	HostAddress string `json:"hostAddress"`
	Port        string `json:"port"`
}
