package jsonfeed

import (
	"time"
)

type Item struct {
	ID               string       `json:"id"`
	URL              string       `json:"url,omitempty"`
	ExternalURL      string       `json:"external_url,omitempty"`
	Title            string       `json:"title,omitempty"`
	ContentText      string       `json:"content_text,omitempty"`
	ContentHTML      string       `json:"content_html,omitempty"`
	Summary          string       `json:"summary,omitempty"`
	ImageURL         string       `json:"image,omitempty"`
	BannerImageURL   string       `json:"banner_image,omitempty"`
	RawDatePublished string       `json:"date_published,omitempty"`
	RawDateModified  string       `json:"date_modified,omitempty"`
	Authors          []Author     `json:"authors,omitempty"`
	Tags             []string     `json:"tags,omitempty"`
	Language         string       `json:"language,omitempty"`
	Attachments      []Attachment `json:"attachments,omitempty"`
}

func (fi *Item) PublishedAt() (time.Time, error) {
	return time.Parse(time.RFC3339, fi.RawDatePublished)
}

func (fi *Item) SetPublishedAt(t time.Time) {
	fi.RawDatePublished = t.Format(time.RFC3339)
}

func (fi *Item) ModifiedAt() (time.Time, error) {
	return time.Parse(time.RFC3339, fi.RawDateModified)
}

func (fi *Item) SetModifiedAt(t time.Time) {
	fi.RawDateModified = t.Format(time.RFC3339)
}

func (fi *Item) IsValid() bool {
	if fi.ID == "" {
		return false
	}
	if fi.ContentText == "" && fi.ContentHTML == "" {
		return false
	}

	return true
}
