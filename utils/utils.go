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

func StandardDeviation(data []data.TradingData, sma []float32, n int) []float32 {
	var stdDev []float32
	var sum = float32(0)
	for i := 0; i < len(data)-n; i++ {
		for j := i; j < i+n; j++ {
			diff := data[j].Close - sma[i]
			sum += diff * diff
		}
		stdDev = append(stdDev, float32(math.Sqrt(float64(sum)/float64(n))))
	}
	return stdDev
}
