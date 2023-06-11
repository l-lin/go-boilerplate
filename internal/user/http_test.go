package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go-boilerplate/internal/config"
	"go-boilerplate/pkg/user"
)

func TestHttpRepository_Get(t *testing.T) {
	type given struct {
		tsFn func() *httptest.Server
	}

	var tests = map[string]struct {
		given given
		test  func(actual *user.User, err error)
	}{
		"happy path": {
			given: given{
				tsFn: func() *httptest.Server {
					return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.WriteHeader(http.StatusOK)
						w.Write([]byte("{\"uuid\": \"454e6ff2-3473-425a-91ac-1a518a92f6a0\"}"))
					}))
				},
			},
			test: func(actual *user.User, err error) {
				if err != nil {
					t.Errorf("expected no error, got: %v", err)
					t.Fail()
				}
				if actual == nil {
					t.Error("expected not nil user")
					t.Fail()
				}
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ts := tt.given.tsFn()
			repository := NewHttpRepository(config.Config{
				URL: ts.URL,
			})
			tt.test(repository.Get("uid"))
		})
	}

}
