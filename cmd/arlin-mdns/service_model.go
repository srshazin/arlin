package arlinmdns

type AdvertisedServiceModel struct {
	HostName    string `json:"hostName"`
	HostAddress string `json:"hostAddress"`
	Port        int    `json:"port"`
}
