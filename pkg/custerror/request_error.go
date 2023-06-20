package custerror

type RequestError struct {
	StatusCode string
	Err        error
}
