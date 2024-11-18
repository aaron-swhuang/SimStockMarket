package client

import (
	"SimStockMarket/constants"
	"SimStockMarket/data"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RequestData struct {
	Code      string `json:"code"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Interval  string `json:"interval"`
}

func StartClient(code, startDate, endDate, interval string) {
	fmt.Printf("Client code: %s, startDate: %s, endDate: %s, interval: %s\n", code, startDate, endDate, interval)

	// Construct request data
	requestData := RequestData{
		Code:      code,
		StartDate: startDate,
		EndDate:   endDate,
		Interval:  interval,
	}

	// Convert request to JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalf("Error marshalling request data: %v", err)
	}

	// Send HTTP POST request
	url := fmt.Sprintf("http://%s/trading-data", constants.TRADING_SERVER)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	// Read respose from server
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Server returned error: %s", string(body))
	}

	log.Println("Received trading data:")
	log.Println(string(body))

	var tradingData []data.TradingData
	err = json.Unmarshal(body, &tradingData)
	if err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)
	}

	// Print trading data
	log.Println("Received trading data:")
	for _, data := range tradingData {
		log.Printf("Code: %s, Date: %s, Time: %s, Open: %.2f, High: %.2f, Low: %.2f, Close: %.2f, Volume: %d\n",
			data.Code, data.Date, data.Time, data.Open, data.High, data.Low, data.Close, data.Volume)
	}
}

func FetchTradingData(code, startDate, endDate, interval string) ([]data.TradingData, error) {
	// 構造請求數據
	requestData := RequestData{
		Code:      code,
		StartDate: startDate,
		EndDate:   endDate,
		Interval:  interval,
	}

	// Convert request data to JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request data: %w", err)
	}

	// Send HTTP POST request
	url := fmt.Sprintf("http://%s/trading-data", constants.TRADING_SERVER)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error sending POST request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned error: %s", string(body))
	}

	// Parse JSON response
	var tradingData []data.TradingData
	if err := json.Unmarshal(body, &tradingData); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return tradingData, nil
}
