package concurrency

import (
	"testing"
	"time"
)

func BenchmarkCheckWebistes(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(SlowWebsiteChecker, urls)
	}

}

func SlowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
