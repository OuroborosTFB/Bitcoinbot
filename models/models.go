package models

type Price struct {
	Last   float32 `json:"last"`
	Buy    float32 `json:"buy"`
	Sell   float32 `json:"sell"`
	Symbol string  `json:"$"`
}
