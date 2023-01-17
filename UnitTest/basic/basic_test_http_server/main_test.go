package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttp(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	}

	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	// Status code test
	if w.Code != 200 {
		t.Error("Http test fail")
	}

	// Return value test
	if w.Body.String() != "pong" {
		t.Error("content fail")
	}

}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Tell the client that the API version is 1.3
	w.Header().Add("API-VERSION", "1.3")
	w.Write([]byte("ok"))
}

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	w := httptest.NewRecorder()
	Handler(w, req)
	// We should get a good status code
	if want, got := http.StatusOK, w.Result().StatusCode; want != got {
		t.Fatalf("expected a %d, instead got: %d", want, got)
	}
	// Make sure that the version was 1.3
	if want, got := "1.3", w.Result().Header.Get("API-VERSION"); want != got {
		t.Fatalf("expected API-VERSION to be %s, instead got: %s", want, got)
	}
}

func TestUpperCaseHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/upper?word=abc", nil)
	w := httptest.NewRecorder()
	upperCaseHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ABC" {
		t.Errorf("expected ABC got %v", string(data))
	}
}

// fake server is same onther service
func TestClientUpperCase(t *testing.T) {
	expected := "dummy data"
	//create new server for test case:  client call to other service
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer svr.Close()
	c := NewClient(svr.URL)
	res, err := c.UpperCase("anything")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	// res: expected\r\n
	// due to the http protocol cleanup response
	res = strings.TrimSpace(res)
	if res != expected {
		t.Errorf("expected res to be %s got %s", expected, res)
	}
}
