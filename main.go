package main

import (
	"fmt"
	"net"
	"time"
)
func main() {
	var websiteName string
	fmt.Println("Enter website name:")
	fmt.Scan(&websiteName)

	// Resolve the IP addresses associated with the given website name
	ipAddresses, err := net.LookupIP(websiteName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the resolved IP addresses
	fmt.Println("IP Addresses for", websiteName)
	for _, ipAddress := range ipAddresses {
		fmt.Println(ipAddress)
	}
	time.Sleep(2 * time.Second)

	fmt.Println("existing")
}
