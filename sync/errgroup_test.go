package sync_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"

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

	func() {
		defer tm.Track(time.Now(), "run without errgroup")
		for _, url := range urls {
			// Launch a goroutine to fetch the URL.
			url := url // https://golang.org/doc/faq#closures_and_goroutines
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			} else {
				fmt.Printf("url: %s, error: %v\n", url, err)
			}
		}
	}()

	func() {
		defer tm.Track(time.Now(), "run with errgroup")
		var g errgroup.Group
		for _, url := range urls {
			// Launch a goroutine to fetch the URL.
			url := url // https://golang.org/doc/faq#closures_and_goroutines
			g.Go(func() error {
				// Fetch the URL.
				resp, err := http.Get(url)
				if err == nil {
					resp.Body.Close()
				} else {
					fmt.Printf("url: %s, error: %v\n", url, err)
				}
				return err
			})
		}
		// Wait for all HTTP fetches to complete.
		if err := g.Wait(); err == nil {
			fmt.Println("Successfully fetched all URLs.")
		}
	}()
}

func TestParallel(t *testing.T) {
	//func() {
	//	defer tm.Track(time.Now(), "run parallel without timeout")
	//	g, _ := errgroup.WithContext(context.Background())
	//	results := make([]string, len(urls))
	//
	//	for idx, url := range urls {
	//		// Launch a goroutine to fetch the URL.
	//		idx, url := idx, url // https://golang.org/doc/faq#closures_and_goroutines
	//		g.Go(func() error {
	//			// Fetch the URL.
	//			resp, err := http.Get(url)
	//			if err == nil {
	//				resp.Body.Close()
	//				results[idx] = fmt.Sprintf("url: %s is OK", url)
	//			} else {
	//				results[idx] = fmt.Sprintf("url: %s is Failed", url)
	//			}
	//			return err
	//		})
	//	}
	//	// Wait for all HTTP fetches to complete.
	//	if err := g.Wait(); err == nil {
	//		fmt.Println("Successfully fetched all URLs.")
	//	}
	//
	//	for _, result := range results {
	//		fmt.Println(result)
	//	}
	//}()

	func() {
		defer tm.Track(time.Now(), "run parallel with timeout")

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)

		g, _ := errgroup.WithContext(timeoutCtx)

		// wait timeout
		go func() {
			select {
			case <-timeoutCtx.Done():
				fmt.Println("timeout")
				// FIXME: cancel() doesn't cancel rest of working goroutne
				cancel()
				return
			}
		}()

		for idx, url := range urls {
			// Launch a goroutine to fetch the URL.
			idx, url := idx, url // https://golang.org/doc/faq#closures_and_goroutines
			g.Go(func() error {
				// Fetch the URL.
				resp, err := http.Get(fmt.Sprintf("%s?123", url))
				if err == nil {
					resp.Body.Close()
					fmt.Printf("[%d]url: %s is OK\n", idx, url)
				} else {
					fmt.Printf("[%d]url: %s is Failed\n", idx, url)
				}
				return err
			})
		}
		// Wait for all HTTP fetches to complete.
		if err := g.Wait(); err == nil {
			fmt.Println("Successfully fetched all URLs.")
		}
	}()
}
