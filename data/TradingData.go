package data

type TradingData struct {
	Code   string  `json:"code"`
	Date   string  `json:"date"`
	Time   string  `json:"time"`
	Open   float32 `json:"open"`
	High   float32 `json:"high"`
	Low    float32 `json:"low"`
	Close  float32 `json:"close"`
	Volume int     `json:"volume"`
}
