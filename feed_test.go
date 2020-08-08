package jsonfeed_test

import (
	"os"
	"testing"

	"git.wark.io/lib/go/jsonfeed"
)

const (
	testFileSmall = "testdata/feed_small.json"
)

func TestParse(t *testing.T) {
	var (
		err error
		f   *os.File
	)

	f, err = os.Open(testFileSmall)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var feed *jsonfeed.Feed

	feed, err = jsonfeed.Parse(f)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := feed.Version, "https://jsonfeed.org/version/1.1"; got != want {
		t.Errorf("feed.Version = %q; want %q", got, want)
	}

	if got, want := feed.Title, "My Example Feed"; got != want {
		t.Errorf("feed.Title = %q; want %q", got, want)
	}

	if got, want := feed.HomePageURL, "https://example.org/"; got != want {
		t.Errorf("feed.HomePageURL = %q; want %q", got, want)
	}

	if got, want := feed.FeedURL, "https://example.org/feed.json"; got != want {
		t.Errorf("feed.FeedURL = %q; want %q", got, want)
	}

	expectedItems := []jsonfeed.Item{
		jsonfeed.Item{
			ID:          "2",
			ContentText: "This is a second item.",
			URL:         "https://example.org/second-item",
		},
		jsonfeed.Item{
			ID:          "1",
			ContentHTML: "<p>Hello, world!</p>",
			URL:         "https://example.org/initial-post",
		},
	}

	if len(feed.Items) == len(expectedItems) {
		for i, gotItem := range feed.Items {
			wantItem := expectedItems[i]

			if got, want := gotItem.ID, wantItem.ID; got != want {
				t.Errorf("feed.Items[%d].ID = %q; want %q", i, got, want)
			}
			if got, want := gotItem.ContentText, wantItem.ContentText; got != want {
				t.Errorf("feed.Items[%d].ContentText = %q; want %q", i, got, want)
			}
			if got, want := gotItem.ContentHTML, wantItem.ContentHTML; got != want {
				t.Errorf("feed.Items[%d].ContentHTML = %q; want %q", i, got, want)
			}
			if got, want := gotItem.URL, wantItem.URL; got != want {
				t.Errorf("feed.Items[%d].URL = %q; want %q", i, got, want)
			}
		}
	} else {
		t.Errorf("len(feed.Items) = %d; want %d", len(feed.Items), len(expectedItems))
	}
}
