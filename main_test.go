package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	assert.Equal(t, http.StatusOK, status)

	body := responseRecorder.Body
	assert.NotEmpty(t, body)
}

func TestMainHandleWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=samara", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	assert.Equal(t, http.StatusBadRequest, status)

	expected := "wrong city value"
	assert.Equal(t, expected, responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	cafes := responseRecorder.Body.String()
	cafeList := strings.Split(cafes, ", ")
	assert.Len(t, cafeList, totalCount)
}
