package main

import (
	"SimStockMarket/client"
	"SimStockMarket/constants"
	"SimStockMarket/server"
	"SimStockMarket/web"
	"log"

	"flag"
	"fmt"
	"os"
)

func GetVersion() {
	fmt.Printf("Generator version - %s\n", constants.VERSION)
}

var (
	version   = flag.Bool("version", false, "Print generator version")
	interval  = flag.String("interval", "1m", "Interval to request trading data")
	code      = flag.String("code", "2330", "stock code to request data for")
	startDate = flag.String("startDate", "", "start date for data (YYYY-MM-DD)")
	endDate   = flag.String("endDate", "", "end date for data (YYYY-MM-DD)")
	saveFile  = flag.Bool("saveFile", false, "Generate output jason file")
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
		log.Printf("%s, %s, %s, %s\n", *code, *startDate, *endDate, *interval)
		client.StartClient(*code, *startDate, *endDate, *interval, *saveFile)
	case "web":
		web.RunWebServer()
	case "gen":
		client.Generate("trading_data.json")
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
