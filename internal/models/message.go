package models

import "time"

// Message is one text message the connected radio delivered to us, together
// with the link metadata of the packet that carried it.
type Message struct {
	ID       uint32    `json:"id"`
	From     uint32    `json:"from"`
	To       uint32    `json:"to"`
	Channel  uint32    `json:"channel"`
	Text     string    `json:"text"`
	Snr      float32   `json:"snr"`
	Rssi     int32     `json:"rssi"`
	HopsAway *uint32   `json:"hopsAway"`
	RxTime   time.Time `json:"rxTime"`
}
