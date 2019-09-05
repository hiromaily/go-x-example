package sync_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"
)

const (
	Limit  = 3 // 同時実行数の上限
	Weight = 1 // 1処理あたりの実行コスト
)

func doSomething(idx int) {
	fmt.Printf("[%d] sleep 2 second\n", idx)
	time.Sleep(2 * time.Second)
}

func TestSemaphore(t *testing.T) {
	s := semaphore.NewWeighted(Limit)
	var w sync.WaitGroup
	for i := 0; i < 100; i++ {
		w.Add(1)
		s.Acquire(context.Background(), Weight)
		go func(idx int) {
			doSomething(idx)
			s.Release(Weight)
			w.Done()
		}(i)
	}
	w.Wait()
}
