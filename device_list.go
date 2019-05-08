package unifi_api

import (
	"fmt"
)

var deviceCodes = map[string]string{
	"BZ2":      "UniFi AP",
	"BZ2LR":    "UniFi AP-LR",
	"U2HSR":    "UniFi AP-Outdoor+",
	"U2IW":     "UniFi AP-In Wall",
	"U2L48":    "UniFi AP-LR",
	"U2Lv2":    "UniFi AP-LR v2",
	"U2M":      "UniFi AP-Mini",
	"U2O":      "UniFi AP-Outdoor",
	"U2S48":    "UniFi AP",
	"U2Sv2":    "UniFi AP v2",
	"U5O":      "UniFi AP-Outdoor 5G",
	"U7E":      "UniFi AP-AC",
	"U7EDU":    "UniFi AP-AC-EDU",
	"U7Ev2":    "UniFi AP-AC v2",
	"U7HD":     "UniFi AP-HD",
	"U7SHD":    "UniFi AP-SHD",
	"U7NHD":    "UniFi AP-nanoHD",
	"UCXG":     "UniFi AP-XG",
	"UXSDM":    "UniFi AP-BaseStationXG",
	"UCMSH":    "UniFi AP-MeshXG",
	"U7IW":     "UniFi AP-AC-In Wall",
	"U7IWP":    "UniFi AP-AC-In Wall Pro",
	"U7MP":     "UniFi AP-AC-Mesh-Pro",
	"U7LR":     "UniFi AP-AC-LR",
	"U7LT":     "UniFi AP-AC-Lite",
	"U7O":      "UniFi AP-AC Outdoor",
	"U7P":      "UniFi AP-Pro",
	"U7MSH":    "UniFi AP-AC-Mesh",
	"U7PG2":    "UniFi AP-AC-Pro",
	"p2N":      "PicoStation M2",
	"US8":      "UniFi Switch 8",
	"US8P60":   "UniFi Switch 8 POE-60W",
	"US8P150":  "UniFi Switch 8 POE-150W",
	"S28150":   "UniFi Switch 8 AT-150W",
	"USC8":     "UniFi Switch 8",
	"US16P150": "UniFi Switch 16 POE-150W",
	"S216150":  "UniFi Switch 16 AT-150W",
	"US24":     "UniFi Switch 24",
	"US24P250": "UniFi Switch 24 POE-250W",
	"US24PL2":  "UniFi Switch 24 L2 POE",
	"US24P500": "UniFi Switch 24 POE-500W",
	"S224250":  "UniFi Switch 24 AT-250W",
	"S224500":  "UniFi Switch 24 AT-500W",
	"US48":     "UniFi Switch 48",
	"US48P500": "UniFi Switch 48 POE-500W",
	"US48PL2":  "UniFi Switch 48 L2 POE",
	"US48P750": "UniFi Switch 48 POE-750W",
	"S248500":  "UniFi Switch 48 AT-500W",
	"S248750":  "UniFi Switch 48 AT-750W",
	"US6XG150": "UniFi Switch 6XG POE-150W",
	"USXG":     "UniFi Switch 16XG",
	"UGW3":     "UniFi Security Gateway 3P",
	"UGW4":     "UniFi Security Gateway 4P",
	"UGWHD4":   "UniFi Security Gateway HD",
	"UGWXG":    "UniFi Security Gateway XG-8",
	"UP4":      "UniFi Phone-X",
	"UP5":      "UniFi Phone",
	"UP5t":     "UniFi Phone-Pro",
	"UP7":      "UniFi Phone-Executive",
	"UP5c":     "UniFi Phone",
	"UP5tc":    "UniFi Phone-Pro",
	"UP7c":     "UniFi Phone-Executive",
}

func DeviceCodeToName(code string) (string, error) {
	if name, ok := deviceCodes[code]; ok {
		return name, nil
	}

	return "", fmt.Errorf("no device with code %s found", code)
}
