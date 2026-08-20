package models

// NodeInfo is the compact node projection returned by the node list: only the
// identity fields, without the full telemetry and position record.
type NodeInfo struct {
	Num       uint32 `json:"num"`
	ShortName string `json:"shortName"`
	LongName  string `json:"longName"`
}
