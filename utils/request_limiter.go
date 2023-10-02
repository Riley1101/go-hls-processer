package utils

import (
	"net/http"
)

type RequestValidator struct {
	Request  *http.Request
	Response http.ResponseWriter
	Limit    int64 `default:"5"`
	Interval int64 `default:"10"`
}

func NewRequestValidator(w http.ResponseWriter, r *http.Request) *RequestValidator {
	return &RequestValidator{
		Response: w,
		Request:  r,
		Limit:    5,
		Interval: 5,
	}
}

func (r *RequestValidator) ValidateRequest() bool {

	request_id := r.Request.Header.Get("X-Request-Id")

	if request_id == "" {
		return false
	}
	limit := TokenBucket(request_id, r.Limit, r.Interval)
	if !limit {
		return false
	}
	return true
}
