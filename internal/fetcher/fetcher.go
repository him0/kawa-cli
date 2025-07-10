package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type ImageFetcher struct {
	client *http.Client
}

func New() *ImageFetcher {
	return &ImageFetcher{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (f *ImageFetcher) Fetch(url string) ([]byte, error) {
	resp, err := f.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return data, nil
}