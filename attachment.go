package jsonfeed

type Attachment struct {
	URL               string  `json:"url"`
	MimeType          string  `json:"mime_type"`
	Title             string  `json:"title,omitempty"`
	SizeInBytes       uint64  `json:"size_in_bytes,omitempty"`
	DurationInSeconds float64 `json:"duration_in_seconds,omitempty"`
}

func (fia *Attachment) IsValid() bool {
	if fia.URL == "" {
		return false
	}
	if fia.MimeType == "" {
		return false
	}

	return true
}
