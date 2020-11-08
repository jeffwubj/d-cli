package docker

import (
	"github.com/jeffwubj/d-cli/pkg/kubectl"
	"github.com/jeffwubj/d-cli/pkg/printconsole"
	"github.com/jeffwubj/d-cli/pkg/utils"
	"os"
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
		hostport, err := utils.GetFreePort()
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
