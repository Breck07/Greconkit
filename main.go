package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Breck07/Greconkit/portscanner"
)

func main() {
	var port int
	var err error
	// Initialize scanner
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a port to scan: ")
	for scanner.Scan() {
		input := scanner.Text()
		port, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error with converting string to int.")
			scanner.Err()
			continue
		}

		fmt.Printf("Starting scan on port: %d\n", port)
		break
	}

	portScanner := portscanner.NewScanner("localhost", 5*time.Millisecond, 75)
	isOpen := portScanner.ScanPort(port)

	if isOpen {
		fmt.Printf("Port: %d is open\n", port)
	}

}
