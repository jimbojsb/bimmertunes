package bimmertunes

type Config struct {
	Airplay AirplayConfig
	Wifi    WifiConfig
	Ibus    IbusConfig
}

type AirplayConfig struct {
	Name string
}

type WifiConfig struct {
	Ssid     string
	Password string
}

type IbusConfig struct {
	Port         string
	Display      string
	IkeBroadcast bool
}
