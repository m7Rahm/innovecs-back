package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetToDos(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(GetToDos))
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200 got %v", resp.StatusCode)
	}
}

func TestPostToDo(t *testing.T) {
	payload := struct {
		Todo string `json:"todo"`
	}{"monster"}
	payload_j, _ := json.Marshal(payload)
	server := httptest.NewServer(http.HandlerFunc(PostToDo))
	defer server.Close()
	resp, err := http.Post(server.URL, "application/json", bytes.NewReader(payload_j))
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	defer resp.Body.Close()
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if resp.StatusCode != 201 {
		t.Errorf("expected 201 got %v", resp.StatusCode)
	}
}
