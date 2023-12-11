package nowthatssomespeed

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return true
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsited(t *testing.T) {
	t.Run("Check Websites", func(t *testing.T) {
		websites := []string{
			"https://google.com",
			"https://lostboy.com",
			"https://passwordmanagers.com",
		}

		got := CheckWebsites(mockWebsiteChecker, websites)
		want := map[string]bool{
			"https://google.com":           true,
			"https://lostboy.com":          true,
			"https://passwordmanagers.com": true,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got different results")
		}

	})
}

func BenchmarkCheckWebsites(b *testing.B) {
	defaultUrl := []string{"https://samaybhavan.in"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, defaultUrl)
	}
}
