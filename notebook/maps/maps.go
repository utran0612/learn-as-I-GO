//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

func displayServerInfo(servers map[string]string){
	numOfServers := len(servers)
	fmt.Println("Number of servers is: ",numOfServers)

	statuses := make(map[string]int)
	for _,status := range servers{
		switch status := status; {
		case status == "Online":
			statuses["online"] += 1
		case status == "Offline":
			statuses["offline"] += 1
		case status == "Maintenance":
			statuses["maint"] += 1
		case status == "Retired":
			statuses["ret"] += 1
		}
	}

	for status, num := range statuses{
		fmt.Println("Number of",status,"servers is",num)
	}
}

func main() {
	servers := map[string]string{
		"london":"Online",
		"nyc": "Online",
		"macau":"Online",
		"singapore":"Online",
		"madrid":"Online",
		"darkstar":"Online",
		"auir":"Online",
	}

	displayServerInfo(servers)

	servers["darkstar"] = "Retired"
	servers["auir"] = "Offline"

	displayServerInfo(servers)

	for _, status := range servers{
		if status != "Maintenance"{
			status = "Maintenace"
		}
	}

	displayServerInfo(servers)


}

