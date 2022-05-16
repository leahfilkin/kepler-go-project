package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// T - This type contains useful methods that you will need to output results, log errors to the screen, and signal failures, like the t. Errorf() method
func TestHello(t *testing.T) {
	// test for correct content type
	tests := map[string]struct {
		url        string
		statusCode int
		body       string
	}{
		"correctUrl": {
			url:        "/hello?msg=my%20test%20message",
			statusCode: http.StatusNoContent,
			body:       "my test message",
		},
		"incorrectUrl": {
			url:        "/hello?msg=1",
			statusCode: http.StatusBadRequest,
			body:       "You cannot use integers in your message string\n",
		},
		"mixtureOfLettersAndNumbers": {
			url:        "/hello?msg=abc%20def%20gh1%20jkl",
			statusCode: http.StatusBadRequest,
			body:       "You cannot use integers in your message string\n",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodGet, test.url, nil)
			writer := httptest.NewRecorder()
			hello(writer, request)
			resp := writer.Result()
			body, _ := ioutil.ReadAll(resp.Body)
			if string(body) != test.body {
				t.Errorf("got: %q want: %q", string(body), test.body)
			}
			if err != nil {
				t.Errorf("Error: %s", err)
			}

			if resp.StatusCode != test.statusCode {
				t.Errorf("got %d, want %d", resp.StatusCode, test.statusCode)
			}
		})
	}
}

func TestHealth(t *testing.T) {
	//incase i wanna write more tests
	tests := map[string]struct {
		url        string
		statusCode int
		body       string
	}{
		"correctUrl": {
			url:        "/health",
			statusCode: http.StatusNoContent,
			body:       fmt.Sprintf("Status code: %s", http.StatusText(http.StatusNoContent)),
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodGet, test.url, nil)
			writer := httptest.NewRecorder()
			health(writer, request)
			resp := writer.Result()
			body, _ := ioutil.ReadAll(resp.Body)
			if string(body) != test.body {
				t.Errorf("got: %q want: %q", string(body), test.body)
			}
			if err != nil {
				t.Errorf("Error: %s", err)
			}

			if resp.StatusCode != test.statusCode {
				t.Errorf("got %d, want %d", resp.StatusCode, test.statusCode)
			}
		})
	}
}

func TestMetadata(t *testing.T) {
	tests := map[string]struct {
		url        string
		statusCode int
		expected   map[string][]Metadata
	}{
		"correctUrl": {
			url:        "/metadata",
			statusCode: http.StatusNoContent,
			expected: map[string][]Metadata{
				"myapplication": {
					{
						Version:       "1.0",
						Description:   "pre-interview technical test",
						LastCommitSha: "abc57858585",
					},
				},
			},
		},
	}
	for testName, tt := range tests {
		// test for correct content type
		// counter
		// mutex
		t.Run(testName, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, tt.url, nil)
			writer := httptest.NewRecorder()
			metadata(writer, request)
			resp := writer.Result()
			decoder := json.NewDecoder(resp.Body)
			var actual map[string][]Metadata
			err := decoder.Decode(&actual)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("got: %v want: %v", actual, tt.expected)
			}
			if err != nil {
				t.Errorf("Error: %s", err)
			}
			if resp.StatusCode != tt.statusCode {
				t.Errorf("got %d, want %d", resp.StatusCode, tt.statusCode)
			}
		})
	}
}
