package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "wtf://daily-wtf.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.co.in",
		"https://cricinfo.com",
		"wtf://daily-wtf.com",
	}

	want := map[string]bool{
		"https://google.co.in": true,
		"https://cricinfo.com": true,
		"wtf://daily-wtf.com":  false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}

}
