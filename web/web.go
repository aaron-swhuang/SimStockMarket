package web

import (
	"SimStockMarket/client"
	"encoding/json"
	"log"
	"net/http"
)

// RequestData
type RequestData struct {
	Code      string `json:"code"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Interval  string `json:"interval"`
}

// RunWebServer starts web server
func RunWebServer() {
	http.Handle("/", http.FileServer(http.Dir("./www"))) // put index.html under ./www directory

	// handle /get-data route
	http.HandleFunc("/get-data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse request
		var requestData RequestData
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Fetch data through client
		tradingData, err := client.FetchTradingData("localhost:8080", requestData.Code, requestData.StartDate, requestData.EndDate, requestData.Interval)
		if err != nil {
			http.Error(w, "Error fetching trading data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Return result
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tradingData)
	})

	log.Println("Web server started on port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
