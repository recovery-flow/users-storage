package cli

import (
	"context"
	"sync"

	"github.com/cifra-city/users-storage/internal/service"
)

func runServices(ctx context.Context, wg *sync.WaitGroup) {
	var ()

	run := func(f func()) {
		wg.Add(1)
		go func() {
			f()
			wg.Done()
		}()
	}

	run(func() { service.Run(ctx) })
}
