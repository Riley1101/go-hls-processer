package utils

import (
	"fmt"
	"net/http"
)

type RequestValidator struct {
	Request  *http.Request
	Response http.ResponseWriter
	Limit    int64 `default:"5"`
	Interval int64 `default:"20"`
}

func NewRequestValidator(w http.ResponseWriter, r *http.Request) *RequestValidator {
	return &RequestValidator{
		Response: w,
		Request:  r,
		Limit:    5,
		Interval: 20,
	}
}

func (r *RequestValidator) ValidateRequest() bool {

	request_id := r.Request.Header.Get("X-Request-Id")

	if request_id == "" {
		fmt.Fprintf(r.Response, "Please provide a request id")
		return false
	}
	limit := TokenBucket(request_id, r.Limit, r.Interval)
	if !limit {
		fmt.Println("You have exceeded the rate limit wait for the two processes to finish")
		return false
	}
	return true
}
