package utils

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func WifiData(platform string) (WifiInterface, error) {
	switch platform {
	case "windows":
		return ForWindows()
	default:
		return WifiInterface{}, fmt.Errorf("unsuported platform")
	}
}

// ForWindows process wifi details for te windows platform
func ForWindows() (WifiInterface, error) {
	cmd := exec.Command("Netsh", "WLAN", "show", "interfaces")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return WifiInterface{}, err
	}

	cmdOutput := string(stdout[:])

	return WifiInterface{
		Name:           GetWinKeyValue("Name", cmdOutput),
		Description:    GetWinKeyValue("Description", cmdOutput),
		GUID:           GetWinKeyValue("GUID", cmdOutput),
		MACAddress:     GetWinKeyValue("Physical address", cmdOutput),
		Type:           GetWinKeyValue("Interface type", cmdOutput),
		State:          GetWinKeyValue("State", cmdOutput),
		SSID:           GetWinKeyValue("SSID", cmdOutput),
		BSSID:          GetWinKeyValue("BSSID", cmdOutput),
		NetworkType:    GetWinKeyValue("Network type", cmdOutput),
		RadioType:      GetWinKeyValue("Radio type", cmdOutput),
		Authentication: GetWinKeyValue("Authentication", cmdOutput),
		Cipher:         GetWinKeyValue("Cipher", cmdOutput),
		ConnectionMode: GetWinKeyValue("Connection mode", cmdOutput),
		Band:           GetWinKeyValue("Band", cmdOutput),
		Channel:        GetWinKeyValue("Channel", cmdOutput),
		ReceiveRate:    GetWinKeyValue("Receive rate (Mbps)", cmdOutput),
		TransmitRate:   GetWinKeyValue("Transmit rate (Mbps)", cmdOutput),
		Signal:         GetWinKeyValue("Signal", cmdOutput),
	}, nil

}

//func forOSX() (string, error) {}

// GetValueFromCmdResponse given a cmd response match some string and replace a portion of the matched string to get te desired value
// inspired by this cmd in powershell used to get signal strength `(netsh wlan show interfaces) -Match '^\s+Signal' -Replace '^\s+Signal\s+:\s+',''`
func GetValueFromCmdResponse(matchRegex *regexp.Regexp, replaceRegex *regexp.Regexp, input string) string {
	result := replaceRegex.ReplaceAllString(matchRegex.FindString(input), "")
	return result
}

// GetWinKeyValue A simple wrapper for GetValueFromCmdResponse function
func GetWinKeyValue(key string, cmd string) string {
	value := GetValueFromCmdResponse(MustGetWinMatchExpForKey(key), MustGetWinReplaceExpForKey(key), cmd)
	return strings.TrimSpace(value)
}
