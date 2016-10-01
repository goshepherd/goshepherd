package main

import (
	"net/http"
	"testing"
)

func TestBaseURLNotFound(t *testing.T) {
	e := serve().Tester(t)

	e.GET("/").Expect().Status(http.StatusNotFound)
}

func TestUnauthorized(t *testing.T) {
	e := serve().Tester(t)

	e.GET("/api/posts").Expect().Status(http.StatusUnauthorized)
}
