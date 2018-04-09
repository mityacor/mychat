package main

import (
	"flag"
	"os"
	"regexp"

	"github.com/mityacor/mychat/lib"
)

func main() {
	var isHost bool
	flag.BoolVar(&isHost, "listen", false, "Listen on the specified IP")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		connIP := os.Args[2]
		checkIP(connIP)
		lib.RunHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		checkIP(connIP)
		lib.RunGuest(connIP)
	}
}

func checkIP(ip string) bool{
	match, matchErr := regexp.MatchString("^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}
	(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$", ip)
	if matchErr != nil {
		log.Fatal("Error: ", matchErr)
	}

	if match != nil {
		return true
	}else{
		return false
	}


}