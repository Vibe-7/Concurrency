package checker

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type CheckResult struct {
	URL        string
	StatusCode int
	Duration   time.Duration
	Err        error
}

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func CheckURL(ctx context.Context, url string) (*CheckResult, error) {
	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return &CheckResult{
			URL:      url,
			Duration: time.Since(start),
		}, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return &CheckResult{
			URL:      url,
			Duration: time.Since(start),
			Err:      err,
		}, nil
	}
	defer resp.Body.Close()

	return &CheckResult{
		URL:        url,
		StatusCode: resp.StatusCode,
		Duration:   time.Since(start),
	}, nil

}
