package tservlib

import (
	"encoding/json"
	"io"
)

func NewLeague(r io.Reader) (League, error) {
	var plrs League
	err := json.NewDecoder(r).Decode(&plrs)
	if err != nil {
		return nil, err
	}
	return plrs, nil
}
