package models

// Channel is one configured channel on the connected radio.
type Channel struct {
	Index uint32 `json:"index"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}
