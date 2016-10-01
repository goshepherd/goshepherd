package main

import (
	"testing"
	"net/http"
)



func TestBaseURLNotFound(t *testing.T) {
	e := serve().Tester(t)

	e.GET("/").Expect().Status(http.StatusNotFound)
}

func TestPostEmpty(t *testing.T) {
	e := serve().Tester(t)

	e.GET("/api/posts").Expect().Status(http.StatusUnauthorized)
}
