package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	var testCases = []struct {
		name       string
		expectBody string
		expectCode int
	}{
		{
			name:       "ok",
			expectBody: "pong",
			expectCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := setupRouter()
			w := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/ping", nil)
			r.ServeHTTP(w, req)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectCode, w.Code)
			assert.Equal(t, tc.expectBody, w.Body.String())
		})
	}
}
