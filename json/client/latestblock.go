package client

// LatestBlock LatestBlock
type LatestBlock struct {
	Hash       string  `json:"hash"`
	Time       float64 `json:"time"`
	BlockIndex float64 `json:"block_index"`
	Height     float64 `json:"height"`
}
