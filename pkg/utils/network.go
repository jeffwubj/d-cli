package utils

import (
	"net"
	"strconv"
)

func GetFreePort() (port int, err error) {
	// Create a new server without specifying a port
	// which will result in an open port being chosen
	server, err := net.Listen("tcp", ":0")

	// If there's an error it likely means no ports
	// are available or something else prevented finding
	// an open port
	if err != nil {
		return 0, err
	}

	// Defer the closing of the server so it closes
	defer server.Close()

	// Get the host string in the format "127.0.0.1:4444"
	hostString := server.Addr().String()

	// Split the host from the port
	_, portString, err := net.SplitHostPort(hostString)
	if err != nil {
		return 0, err
	}

	// Return the port as an int
	return strconv.Atoi(portString)
}
