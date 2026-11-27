package portscanner

/*
	Author: Mellowz, Version 1.21
	Creation Date: Nov 16, 2026
	This file implements a threaded port scanner - 
	using goroutines and the standard go net library.
	Created with - GO VERSION: 1.27.4

	*WARNING* This port scanner is to be used only on networks - 
	with given permission and educational purposes!
*/

import (
	"net"
	"strconv"
	"sync"
	"time"
)

// Struct of port scanner
type PortScanner struct {
	Host    string
	Timeout time.Duration
	Workers int
}

// This function creates a new port scanner with passed in parameters
func NewScanner(host string, timeout time.Duration, workers int) *PortScanner {
	return &PortScanner{
		Host:    host,
		Timeout: timeout,
		Workers: workers,
	}
}

// Scans a single port and returns boolean value
func (ps *PortScanner) ScanPort(port int) bool {
	ipAddr := net.JoinHostPort(ps.Host, strconv.Itoa(port)) // Itoa to handle IPv6 & IPv4

	// Attempt to connect by dialing via TCP
	conn, err := net.DialTimeout("tcp", ipAddr, ps.Timeout)
	if err != nil {
		return false
	}

	conn.Close()
	return true

}

// This function scans a range of ports using goroutines
func (ps *PortScanner) ScanPortRange(startPort, endPort int) <-chan int {
	ports := make(chan int, ps.Workers) // Channel size is set to the # of workers so each worker has a port
	results := make(chan int)
	var wg sync.WaitGroup // Wait group tracks active workers in sync

	// Starts the workers
	for i := 0; i < ps.Workers; i++ {
		wg.Go(func() {
			for p := range ports {
				if ps.ScanPort(p) {
					results <- p // Send the port into results if it is open
				}
			}
		})
	}

	// Loads ports into channel
	go func() {
		for p := startPort; p <= endPort; p++ {
			ports <- p
		}
		close(ports)
	}()

	// Spins up closing function to wait for wg to finish and then closes results channel
	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// TO-DO add version/software detection
