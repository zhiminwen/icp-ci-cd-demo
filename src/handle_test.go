package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnvHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/env", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleListEnv)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("env handler return non ok status. got %v", status)
	}
}

func TestWelcomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/welcome", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleWelcome)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("welcome handler return non ok status. got %v", status)
	}

	if rr.Body.String() != "Hello world!!" {
		t.Errorf("welcome handler return not expected result. got %v", rr.Body.String())
	}
}
