package utils

import (
	"database/sql"
	"encoding/json"
)

type WifiInterface struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	GUID           string `json:"guid"`
	MACAddress     string `json:"mac_address"`
	Type           string `json:"type"`
	State          string `json:"state"`
	SSID           string `json:"ssid"`
	BSSID          string `json:"bssid"`
	NetworkType    string `json:"network_type"`
	RadioType      string `json:"radio-type"`
	Authentication string `json:"authentication"`
	Cipher         string `json:"cipher"`
	ConnectionMode string `json:"connection_mode"`
	Band           string `json:"band"`
	Channel        string `json:"channel"`
	ReceiveRate    string `json:"receive-rate"`  // Mbps
	TransmitRate   string `json:"transmit-rate"` // Mbps
	Signal         string `json:"signal"`        // percentage
}

func (wifi *WifiInterface) PrettyPrint() string {
	s, _ := json.MarshalIndent(wifi, "", "\t")
	return string(s)
}

// Save commit the current instance to the d
func (wifi *WifiInterface) Save(db *sql.DB) (bool, error) {
	setup, err := wifi.CheckDbIsSetup(db)
	if err != nil {
		return false, err
	}

	if !setup {
		_, err := wifi.setUpDb(db)
		if err != nil {
			return false, err
		}
	}

	// TODO insert values into DB

	return true, nil
}

func (wifi *WifiInterface) SaveError(db *sql.DB) (bool, error) {
	return true, nil
}

// PrepareForSave return the appropriate sql insert statement for insert into sqlLiteDb
func (wifi *WifiInterface) PrepareForSave() string {
	return ""
}

//CheckDbIsSetup check to see if the required relatons have been set up in the target db
func (wifi *WifiInterface) CheckDbIsSetup(db *sql.DB) (bool, error) {
	return true, nil
}

func (wifi *WifiInterface) setUpDb(db *sql.DB) (bool, error) {
	return true, nil
}
