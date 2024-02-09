package jsonfeed

import (
	"encoding/json"
	"fmt"
	"io"
)

func Parse(r io.Reader) (*Feed, error) {
	var err error

	feed := &Feed{}

	d := json.NewDecoder(r)
	err = d.Decode(&feed)
	if err != nil {
		return nil, fmt.Errorf("could not decode json: %w", err)
	}

	return feed, nil
}
