package portscanner

/*
	Author: Mellowz, Version 1.1
	Creation Date: Nov 16, 2026
	This is a port scanner class for implementing a threaded port scanner using goroutines-
	and the standard go net library.
	Created with - GO VERSION: 1.27.4

	*WARNING* This port scanner is to be used only on networks with given permission and educational purposes!
*/

import (
	"fmt"
	"net"
	"time"
)

// Struct for port scanner object
type PortScanner struct {
	Host    string
	Timeout time.Duration
	Workers int
}

// This function initializes a new port scanner object with the given parameters
func NewScanner(host string, timeout time.Duration, workers int) *PortScanner {
	return &PortScanner{
		Host:    host,
		Timeout: timeout,
		Workers: workers,
	}
}

// Function to scan a single port via TCP
func (ps *PortScanner) ScanPort(port int) bool {
	ipAddr := fmt.Sprintf("%s:%d", ps.Host, port)

	// Attempt to connect by dialing
	conn, err := net.DialTimeout("tcp", ipAddr, ps.Timeout)
	if err != nil {
		return false
	}

	conn.Close()
	return true

}

// TO-DO add scan range

// TO-DO add multiple scan types

// TO-DO add version/software detection
