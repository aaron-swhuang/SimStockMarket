package utils

import (
	"SimStockMarket/data"
	"math"
)

func FindMinMax(data []data.TradingData, start int, end int) (float32, float32) {
	max := float32(-math.MaxFloat32)
	min := float32(math.MaxFloat32)

	for i := start; i <= end; i++ {
		if data[i].High > max {
			max = data[i].High
		}
		if data[i].Low < min {
			min = data[i].Low
		}
	}
	return min, max
}
