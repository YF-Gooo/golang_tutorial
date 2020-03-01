package cancel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func PringSelect(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for i := 0; i < 100; i++ {
			fmt.Println(PringSelect(ctx))
		}
	}(ctx)
	time.Sleep(time.Second * 0.001)
	cancel()
}
