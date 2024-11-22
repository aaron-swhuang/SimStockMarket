package test

import (
	"SimStockMarket/data"
	"SimStockMarket/utils"
	"testing"
)

var tradingData = []data.TradingData{
	{
		Code:   "2330",
		Date:   "2024/11/01",
		Time:   "09:00:00",
		Open:   230.5,
		High:   235.0,
		Low:    229.0,
		Close:  234.5,
		Volume: 1500,
	},
	{
		Code:   "2330",
		Date:   "2024/11/02",
		Time:   "09:00:00",
		Open:   234.5,
		High:   238.0,
		Low:    232.0,
		Close:  237.0,
		Volume: 1800,
	},
	{
		Code:   "2330",
		Date:   "2024/11/03",
		Time:   "09:00:00",
		Open:   237.0,
		High:   240.0,
		Low:    235.5,
		Close:  238.5,
		Volume: 2100,
	},
	{
		Code:   "2330",
		Date:   "2024/11/04",
		Time:   "09:00:00",
		Open:   238.5,
		High:   242.0,
		Low:    237.0,
		Close:  241.0,
		Volume: 2200,
	},
	{
		Code:   "2330",
		Date:   "2024/11/05",
		Time:   "09:00:00",
		Open:   241.0,
		High:   245.0,
		Low:    240.0,
		Close:  244.5,
		Volume: 2500,
	},
	{
		Code:   "2330",
		Date:   "2024/11/06",
		Time:   "09:00:00",
		Open:   244.5,
		High:   247.0,
		Low:    242.5,
		Close:  246.0,
		Volume: 2300,
	},
	{
		Code:   "2330",
		Date:   "2024/11/07",
		Time:   "09:00:00",
		Open:   246.0,
		High:   250.0,
		Low:    245.0,
		Close:  249.5,
		Volume: 2400,
	},
	{
		Code:   "2330",
		Date:   "2024/11/08",
		Time:   "09:00:00",
		Open:   249.5,
		High:   253.0,
		Low:    248.0,
		Close:  251.5,
		Volume: 2600,
	},
	{
		Code:   "2330",
		Date:   "2024/11/09",
		Time:   "09:00:00",
		Open:   251.5,
		High:   255.0,
		Low:    250.0,
		Close:  253.0,
		Volume: 2700,
	},
	{
		Code:   "2330",
		Date:   "2024/11/10",
		Time:   "09:00:00",
		Open:   253.0,
		High:   257.0,
		Low:    252.0,
		Close:  255.5,
		Volume: 2800,
	},
}

func TestFindMinMax(t *testing.T) {
	lowest, highest := utils.FindMinMax(tradingData, 0, len(tradingData)-1)
	if lowest != 229 || highest != 257 {
		t.Errorf("unexpected")
	}
}

/*
func TestStandardDeviation(t *testing.T) {
	var expResult = [...]float32 {236.67242, 336.24088, 336.28226, 336.32596, 336.3748, 336.42215, 336.45758}
	stdDev := utils.StandardDeviation(tradingData, indicator.MA(tradingData, 3), 3)
	if stdDev[:] != expResult[:] { t.Errorf("unexpected result array")}
	log.Println(stdDev)

}*/
