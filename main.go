package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	showIntroduction()

	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			showMonitoring()
		case 2:
			fmt.Println("Showing logs...")
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "Jose"
	version := 1.1
	fmt.Println("Hello, ", name)
	fmt.Println("This program version is", version)
}

func showMenu() {
	fmt.Println("\n1. Monitoring")
	fmt.Println("2. Logs")
	fmt.Println("0. Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func showMonitoring() {
	fmt.Println("Monitoring...")

	sites := []string{
		"https://lestetelecom.com.br",
		"https://google.com",
		"https://amazon.com",
		"https://netflix.com",
	}

	for _, site := range sites {
		response, _ := http.Get(site)

		fmt.Println("Site ->", site, "returned status:", response.StatusCode)
	}
}
