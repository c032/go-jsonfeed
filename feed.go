package jsonfeed

const (
	Version1  = "https://jsonfeed.org/version/1"
	Version11 = "https://jsonfeed.org/version/1.1"

	VersionLatest = Version11
)

type Feed struct {
	Version     string   `json:"version"`
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	IconURL     string   `json:"icon,omitempty"`
	FaviconURL  string   `json:"favicon,omitempty"`
	HomePageURL string   `json:"home_page_url,omitempty"`
	FeedURL     string   `json:"feed_url,omitempty"`
	UserComment string   `json:"user_comment,omitempty"`
	NextURL     string   `json:"next_url,omitempty"`
	Authors     []Author `json:"authors,omitempty"`
	Language    string   `json:"language,omitempty"`
	RawExpired  *bool    `json:"expired,omitempty"`
	Items       []Item   `json:"items"`

	// Author has been deprecated in favor of Authors. It's included here only
	// for compatibility.
	Author *Author `json:"author,omitempty"`
}

func (f *Feed) IsExpired() bool {
	return f.RawExpired != nil && *f.RawExpired == true
}
