package jsonfeed_test

import (
	"testing"

	"github.com/c032/go-jsonfeed"
)

func TestFeed_IsExpired_Nil(t *testing.T) {
	feed := &jsonfeed.Feed{}

	if got, want := feed.IsExpired(), false; got != want {
		t.Errorf("(&jsonfeed.Feed{}).IsExpired() = %#v; want %#v", got, want)
	}
}

func TestFeed_IsExpired_False(t *testing.T) {
	isExpired := false
	feed := &jsonfeed.Feed{
		RawExpired: &isExpired,
	}

	if got, want := feed.IsExpired(), false; got != want {
		t.Errorf("feed.IsExpired() = %#v; want %#v", got, want)
	}
}

func TestFeed_IsExpired_True(t *testing.T) {
	isExpired := true
	feed := &jsonfeed.Feed{
		RawExpired: &isExpired,
	}

	if got, want := feed.IsExpired(), true; got != want {
		t.Errorf("feed.IsExpired() = %#v; want %#v", got, want)
	}
}
