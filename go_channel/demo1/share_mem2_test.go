package share_mem

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestCounterOne(t *testing.T) {
	counter := 0
	for i := 0; i < 20000; i++ {
		counter += i
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// counter = 12497500
// ok  	gplang_tutorial/go_channel/demo1	1.322s

func TestCounterTwo(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func(i int) {
			counter += i
		}(i)
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

// counter = 12418014
// ok  	gplang_tutorial/go_channel/demo1	1.291s

func TestCounterWaitChan(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	channel := make(chan int, 10000)
	counter := 0
	maxnum := 20000
	go func(maxnum int) {
		for i := 0; i < maxnum; i++ {
			channel <- i
		}
		close(channel)
	}(maxnum)
	maxcpu := runtime.NumCPU()
	for i := 0; i < maxcpu; i++ {
		wg.Add(1)
		go func(channel chan int) {
			for ii := range channel {
				mut.Lock()
				counter += ii
				mut.Unlock()
			}
			wg.Done()
		}(channel)
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}

// counter = 12497500
// ok  	gplang_tutorial/go_channel/demo1	0.358s
