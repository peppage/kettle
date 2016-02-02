package errors

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	PrivateUser  = 401
	UserNotFound = 500
)

type ApiError struct {
	StatusCode int
	Header     http.Header
	Body       string
	URL        *url.URL
}

func (aerr ApiError) Error() string {
	switch aerr.StatusCode {
	case PrivateUser:
		return fmt.Sprintf("User is private")
	case UserNotFound:
		return fmt.Sprintf("User not found")
	}
	return fmt.Sprintf("Steam returned status %d, %s", aerr.URL, aerr.StatusCode, aerr.Body)
}

func (aerr ApiError) RateLimitCheck() (bool, time.Time) {
	if aerr.StatusCode == 500 || aerr.StatusCode == 503 {
		return true, time.Now().Add(30 * time.Second)
	}
	return false, time.Time{}
}
