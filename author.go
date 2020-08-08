package jsonfeed

type Author struct {
	Name      string `json:"name,omitempty"`
	URL       string `json:"url,omitempty"`
	AvatarURL string `json:"avatar,omitempty"`
}

func (a *Author) IsValid() bool {
	return a.Name != "" || a.URL != "" || a.AvatarURL != ""
}
