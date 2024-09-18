package httpclient

import (
	"net/http"
	"sync"
	"time"
)

var (
	once       sync.Once
	httpClient *http.Client
)

func New() *http.Client {
	once.Do(
		func() {
			httpClient = &http.Client{
				Timeout: time.Duration(10) * time.Second,
			}
		},
	)
	return httpClient
}
