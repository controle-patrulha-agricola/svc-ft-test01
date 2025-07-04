package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEchoEndpoint(t *testing.T) {
	resp, err := http.Get("http://localhost:8888/echo/hello")
	assert.NoError(t, err)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	result := strings.TrimSpace(string(body))

	assert.Equal(t, "hellocpa", result)
}
