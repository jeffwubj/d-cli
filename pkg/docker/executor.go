package docker

import (
	"github.com/jeffwubj/d-cli/pkg/printconsole"
	"net"
	"os"
	"strconv"

	"github.com/jeffwubj/d-cli/pkg/kubectl"
	"sync"
)

func Run() {
	var wg sync.WaitGroup

	err := EnsureRuntime()
	if err != nil {
		printconsole.PrintWarning(err.Error())
		os.Exit(1)
	}

	wg.Add(1)
	go func() {
		hostport, err := getFreePort()
		if err != nil {
			wg.Done()
			printconsole.PrintError(err.Error())
			os.Exit(1)
		}
		SetDockerHost(hostport)
		podname, err := GetRunningRuntimePod()
		if err != nil {
			wg.Done()
			printconsole.PrintError(err.Error())
			os.Exit(1)
		}
		podport := DefaultDockerdPort
		kubectl.RunPortForward(&wg, podname, DefaultManagementNamespace, hostport, podport, true)
	}()
	wg.Wait()
	RunDockerCommand()
}

func getFreePort() (port int, err error) {
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
