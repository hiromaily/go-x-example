package sync_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"

	tm "github.com/hiromaily/golibs/time"
)

var urls = []string{
	"https://github.com",
	"https://www.cnet.com/",
	"https://thenextweb.com",
	"https://golang.org/",
	"https://www.google.com/",
	"https://www.oreilly.com/",
	"http://www.some-stupid-name.com/",
}

func TestErrorGroup(t *testing.T) {
	//common func
	fn := func(url string) error {
		// Fetch the URL.
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
		} else {
			fmt.Printf("url: %s, error: %v\n", url, err)
		}
		return err
	}
	//without errgroup.Group
	func() {
		defer tm.Track(time.Now(), "run without errgroup")
		for _, url := range urls {
			// Fetch the URL.
			fn(url)
		}
	}()

	//with errgroup.Group
	func() {
		defer tm.Track(time.Now(), "run with errgroup")
		var g errgroup.Group
		for _, url := range urls {
			// Launch a goroutine to fetch the URL.
			url := url // https://golang.org/doc/faq#closures_and_goroutines
			g.Go(func() error {
				return fn(url)
			})
		}
		// Wait for all HTTP fetches to complete.
		if err := g.Wait(); err == nil {
			fmt.Println("Successfully fetched all URLs.")
		}
	}()
}

func TestErrGroupWithSemaphone(t *testing.T){
	defer tm.Track(time.Now(), "run errgroup with Semaphone")

	const maxWorkers = 10
	sem := semaphore.NewWeighted(maxWorkers)
	g, ctx := errgroup.WithContext(context.Background())

	//common func
	fn := func(url string) error {
		// Fetch the URL.
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
		} else {
			fmt.Printf("url: %s, error: %v\n", url, err)
		}
		return err
	}

	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			err := sem.Acquire(ctx, 1)
			if err != nil {
				return err
			}
			defer sem.Release(1)

			return fn(url)
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}

func TestParallel(t *testing.T) {
	//common func
	fn := func(idx int, url string) error {
		// Fetch the URL.
		resp, err := http.Get(fmt.Sprintf("%s?123", url))
		if err == nil {
			resp.Body.Close()
			fmt.Printf("[%d]url: %s is OK\n", idx, url)
		} else {
			fmt.Printf("[%d]url: %s is Failed\n", idx, url)
		}
		return err
	}

	//with errgroup.Group with context
	func() {
		defer tm.Track(time.Now(), "run parallel without timeout")
		g, _ := errgroup.WithContext(context.Background())
		results := make([]string, len(urls))

		for idx, url := range urls {
			// Launch a goroutine to fetch the URL.
			idx, url := idx, url // https://golang.org/doc/faq#closures_and_goroutines
			g.Go(func() error {
				return fn(idx, url)
			})
		}
		// Wait for all HTTP fetches to complete.
		if err := g.Wait(); err == nil {
			fmt.Println("Successfully fetched all URLs.")
		}

		for _, result := range results {
			fmt.Println(result)
		}
	}()
}

//FIXME: timeout doesn't work yet
func TestParallelWithTimeout(t *testing.T) {
	defer tm.Track(time.Now(), "run parallel with timeout")

	fn := func(ctx context.Context, eg *errgroup.Group, idx int, url string) {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				// abort on context cancel
				fmt.Println("timeout")
				return nil
			default:
				// Fetch the URL.
				resp, err := http.Get(fmt.Sprintf("%s?123", url))
				if err == nil {
					resp.Body.Close()
					fmt.Printf("[%d]url: %s is OK\n", idx, url)
				} else {
					fmt.Printf("[%d]url: %s is Failed\n", idx, url)
				}
				return err
			}
		})
	}

	g, egCtx := errgroup.WithContext(context.Background())
	timeoutCtx, cancel := context.WithTimeout(egCtx, 500*time.Millisecond)
	defer cancel()

	// wait timeout
	// if timeout is used, g.Wait() may not be compatible
	//go func() {
	//	select {
	//	case <-timeoutCtx.Done():
	//		fmt.Println("timeout")
	//		// FIXME: cancel() doesn't cancel rest of working goroutne
	//		cancel()
	//		return
	//	}
	//}()

	for idx, url := range urls {
		// Launch a goroutine to fetch the URL.
		//idx, url := idx, url // https://golang.org/doc/faq#closures_and_goroutines
		//g.Go(func() error {
		//	return fn(idx, url)
		//})
		fn(timeoutCtx, g, idx, url)
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}
