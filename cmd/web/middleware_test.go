package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoServe(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}
