package groutine_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)
	go func(ctx context.Context) {
		i := 1
		for {
			if isCancelled(ctx) {
				break
			}
			i += 1
		}
		defer func() {
			ch <- i
			fmt.Println("Cancelled")
		}()
	}(ctx)
	time.Sleep(time.Second * 1)
	cancel()
	fmt.Println(<-ch)

}
