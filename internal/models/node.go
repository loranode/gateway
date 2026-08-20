package models

import "time"

// Node is one entry in the mesh node database as the connected radio currently
// knows it. Optional fields are pointers so the API reports null when a value
// was never observed, and so partial updates can leave unknown fields untouched.
type Node struct {
	Num        uint32     `json:"num"`
	LongName   string     `json:"longName"`
	ShortName  string     `json:"shortName"`
	HwModel    string     `json:"hwModel"`
	Role       string     `json:"role"`
	Snr        float32    `json:"snr"`
	Rssi       int32      `json:"rssi"`
	HopsAway   *uint32    `json:"hopsAway"`
	Latitude   *float64   `json:"latitude"`
	Longitude  *float64   `json:"longitude"`
	Altitude   *int32     `json:"altitude"`
	Battery    *uint32    `json:"battery"`
	Voltage    *float32   `json:"voltage"`
	ViaMqtt    bool       `json:"viaMqtt"`
	IsFavorite bool       `json:"isFavorite"`
	LastHeard  *time.Time `json:"lastHeard"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}
