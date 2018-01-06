package bimmertunes

type Config struct {
	Airplay   AirplayConfig
	Bluetooth BluetoothConfig
	Wifi      WifiConfig
	Ibus      IbusConfig
	Web       WebConfig
}

type AirplayConfig struct {
	Enabled bool
	Name    string
}

type BluetoothConfig struct {
	Enabled  bool
	Name     string
	Passcode string
	Codec    string
}

type WifiConfig struct {
	Enabled  bool
	Ssid     string
	Password string
}

type IbusConfig struct {
	Port         string
	Display      string
	IkeBroadcast bool
}

type WebConfig struct {
	Enabled bool
}
