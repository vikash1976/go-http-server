package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"vikash@google.com", "gopher, vikash"},
		{"vikash", "dear, vikash"},
	}
	for _, c := range cases {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/"+c.in,
			nil,
		)
		if err != nil {
			t.Fatalf("Error creating request %v\n", err)
		}
		rec := httptest.NewRecorder()

		myHandleFunc(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected 200 received %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), c.out) {
			t.Errorf("unexpected body in response %q", rec.Body.String())
		}

	}
}

func BenchmarkHandler(b *testing.B) {

	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/vp@google.com",
			nil,
		)
		if err != nil {
			b.Fatalf("Error creating request %v\n", err)
		}
		rec := httptest.NewRecorder()

		myHandleFunc(rec, req)

		if rec.Code != http.StatusOK {
			b.Errorf("Expected 200 received %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "gopher, vp") {
			b.Errorf("unexpected body in response %q", rec.Body.String())
		}

	}
}
