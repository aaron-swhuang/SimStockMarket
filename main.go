package main

import (
	"SimStockMarket/client"
	"SimStockMarket/server"
	"SimStockMarket/web"

	"flag"
	"fmt"
	"os"
)

const VERSION = "1.0"

func GetVersion() {
	fmt.Printf("Generator version - %s\n", VERSION)
}

var (
	version    = flag.Bool("version", false, "Print generator version")
	serverAddr = flag.String("addr", "localhost:8080", "WebSocket server address")
	interval   = flag.String("interval", "1m", "Interval to request trading data")
	code       = flag.String("code", "2330", "stock code to request data for")
	startDate  = flag.String("startDate", "", "start date for data (YYYY-MM-DD)")
	endDate    = flag.String("endDate", "", "end date for data (YYYY-MM-DD)")
)

func main() {
	if len(os.Args) < 2 {
		PrintUsage()
		return
	}

	// check the command is `server` or `client`
	command := os.Args[1]
	flag.CommandLine.Parse(os.Args[2:])

	if *version {
		GetVersion()
		os.Exit(0)
	}

	switch command {
	case "server":
		server.StartServer()
	case "client":
		fmt.Printf("%s, %s, %s, %s\n", *code, *startDate, *endDate, *interval)
		client.StartClient(*serverAddr, *code, *startDate, *endDate, *interval)
	case "web":
		web.RunWebServer()
	default:
		{
			fmt.Println("Invalid argument.")
			PrintUsage()
		}
	}
}

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("  server - start the server")
	fmt.Println("  client - start the client")
	fmt.Println("  web    - start the web server")
}
