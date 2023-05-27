package datasource

import (
	"fmt"
	"net/http"
)

// NewSdnHTTPReader constructor
func NewSdnHTTPReader(path string) (ISdnReader, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("need response code 200 given: %d", resp.StatusCode)
	}
	return resp.Body, nil
}
