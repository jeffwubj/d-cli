package vm

import (
	"github.com/jeffwubj/d-cli/pkg/restapi"
	"sync"
)

func RunRestApiServer(wg *sync.WaitGroup, port int) error {
	err := restapi.RunRestApi(wg, port)
	if err != nil {
		panic(err)
	}
	return nil
}


func Run() error {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		RunRestApiServer(&wg, RestAPIPort)
	}()
	wg.Wait()
	return nil
}