package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"SimStockMarket/data"
)

/*
"Date","Time","Open","High","Low","Close","TotalVolume"
1999/7/21,09:01:00,345.00,345.00,345.00,345.00,10
*/

// Rule:
// lowest | highest | lifting
//	0.01  |    5    |  0.01
//	  5   |   10    |  0.01
//	 10   |   50    |  0.05
//	 50   |  100    |  0.10
//	100   |  150    |  0.50
//	150   |  500    |  0.50
//	500   | 1000    |  1.00
// 1000   |    ~    |  5.00

func GenerateTradingData(code string, currentTime time.Time) data.TradingData {
	date := currentTime.Format("2006/01/02")

	tradingTime := currentTime.Format("15:04:05")

	// Generate price and volume randomly
	openPrice := float32(rand.Intn(200) + 100) // （range: 100 to 300）
	highPrice := openPrice + float32(rand.Intn(20))
	lowPrice := openPrice - float32(rand.Intn(20))
	closePrice := openPrice + float32(rand.Intn(10))

	volume := rand.Intn(5000)

	// Return trading data
	return data.TradingData{
		Code:   code,
		Date:   date,
		Time:   tradingTime,
		Open:   openPrice,
		High:   highPrice,
		Low:    lowPrice,
		Close:  closePrice,
		Volume: volume,
	}
}

func IsValidTradingTime(current time.Time) bool {

	return current.Weekday() != time.Saturday && current.Weekday() != time.Sunday
}

func GenerateDataSeries(
	code string, startDate time.Time, endDate time.Time,
	interval time.Duration) []data.TradingData {

	var dataSeries []data.TradingData
	startTime := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 9, 0, 0, 0, startDate.Location())
	endTime := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 13, 30, 0, 0, startDate.Location())

	for !startTime.After(endTime) {
		if IsValidTradingTime(startTime) {
			dataSeries = append(dataSeries, GenerateTradingData(code, startTime))

		}
		startTime = startTime.Add(interval)
	}

	return dataSeries
}

func ParseInterval(intv string) (time.Duration, error) {
	if len(intv) < 2 {
		return 0, errors.New("Invalid interval")
	}

	valueStr := intv[:len(intv)-1]
	unit := intv[len(intv)-1:]

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, errors.New("interval value is not a valid number")
	}

	switch unit {
	case "s":
		return time.Duration(value) * time.Second, nil
	case "m":
		return time.Duration(value) * time.Minute, nil
	case "h":
		return time.Duration(value) * time.Hour, nil
	case "d":
		return time.Duration(value) * 24 * time.Hour, nil
	default:
		return 0, errors.New("unknown interval unit")
	}
}

func HandleTradingData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqData struct {
		Code      string `json:"code"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		Interval  string `json:"interval"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	// Parse request
	startDate, err := time.Parse("2006-01-02", reqData.StartDate)
	if err != nil {
		http.Error(w, "Invalid start date format", http.StatusBadRequest)
		return
	}
	endDate, err := time.Parse("2006-01-02", reqData.EndDate)
	if err != nil {
		http.Error(w, "Invalid end date format", http.StatusBadRequest)
		return
	}
	interval, err := ParseInterval(reqData.Interval)
	if err != nil {
		http.Error(w, "Invalid interval format", http.StatusBadRequest)
		return
	}

	// Generate trading data
	dataSeries := GenerateDataSeries(reqData.Code, startDate, endDate, interval)

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dataSeries); err != nil {
		http.Error(w, "Error encoding response data", http.StatusInternalServerError)
		return
	}
	log.Printf("Sent %d trading data entries for code %s", len(dataSeries), reqData.Code)
}

func StartServer() {
	http.HandleFunc("/trading-data", HandleTradingData)
	fmt.Println("HTTP server started on http://localhost:8080/trading-data")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
