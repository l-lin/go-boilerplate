package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_DoOnlyOk(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/success" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}))
	defer ts.Close()

	client := NewClient(false)

	request, err := http.NewRequest("GET", ts.URL+"/success", nil)
	if err != nil {
		t.Errorf("could not create request: %v", err)
		t.Fail()
	}
	response, err := client.DoOnlyOk(request)
	if err != nil {
		t.Errorf("should not have errors: %v", err)
		t.Fail()
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("expected response status %d, actual %d", http.StatusOK, response.StatusCode)
		t.Fail()
	}

	request, err = http.NewRequest("GET", ts.URL+"/fail", nil)
	if err != nil {
		t.Errorf("could not create request: %v", err)
		t.Fail()
	}
	response, err = client.DoOnlyOk(request)
	if err == nil {
		t.Errorf("should have error")
		t.Fail()
	}
}
