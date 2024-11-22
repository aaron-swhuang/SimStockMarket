package indicator

import (
	"SimStockMarket/data"
	"SimStockMarket/utils"
	"log"
)

// Moving Average
func MA(data []data.TradingData, period int) []float32 {
	ma := make([]float32, len(data))

	for i := 0; i < len(data); i++ {
		if i < period-1 {
			ma[i] = 0
		} else {
			sum := float32(0.0)
			for j := i; i-period < j; j-- {
				sum += data[j].Close
			}
			ma[i] = sum / float32(period)
		}
	}
	return ma
}

// KD, Stochastic Oscillator
func KDLine(data []data.TradingData, n int, alpha, beta float32) ([]float32, []float32) {
	rsv := RSV(data, n)
	kline := make([]float32, len(rsv))
	dline := make([]float32, len(rsv))

	kline[0] = 50
	dline[0] = 50

	for i := 1; i < len(rsv); i++ {
		kline[i] = alpha*rsv[i] + (1-alpha)*kline[i-1]
		dline[i] = beta*kline[i] + (1-beta)*dline[i-1]
	}

	return kline, dline
}

// RSV, Raw Stochastic Value
// (Current.Close - Lowest(n)) / (Highest(n) - Lowest(n)) * 100
func RSV(data []data.TradingData, n int) []float32 {
	log.Println("len: ", len(data), ", n: ", n)
	// Sample is not enough
	if len(data) < n {
		return nil
	}

	rsvResult := make([]float32, len(data))

	// There's no RSV if samples are not enough
	for i := n - 1; i < len(data); i++ {
		lowest, highest := utils.FindMinMax(data, i-n+1, i)
		log.Println("min:", lowest, ", max:", highest)

		if lowest != highest {
			rsvResult[i] = (data[i].Close - lowest) / (highest - lowest) * 100
		} else {
			rsvResult[i] = 0
		}

	}
	return rsvResult
}

func BollingerBands(data []data.TradingData, n int, k float32) ([]float32, []float32, []float32) {
	// k is the standard deviation multiplier
	sma := MA(data, n)
	stdDev := utils.StandardDeviation(data, sma, n)

	var upperBand, lowerBand []float32
	for i := 0; i < len(sma); i++ {
		upperBand = append(upperBand, sma[i]+k*stdDev[i])
		lowerBand = append(lowerBand, sma[i]-k*stdDev[i])
	}
	return upperBand, sma, lowerBand
}
