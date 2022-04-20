package main

import (
	"time"

	"github.com/poolpOrg/maxworkers"
)

func main() {
	wg := maxworkers.NewGroup(10)
	for i := 0; i < 10; i++ {
		wg.Run(func() {
			time.Sleep(1 * time.Second)
		})
	}
	wg.Wait()
}
