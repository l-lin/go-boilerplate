package user

import (
	"github.com/l-lin/go-boilerplate/conf"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpRepository_Get(t *testing.T) {
	type given struct {
		tsFn func() *httptest.Server
	}

	var tests = map[string]struct {
		given given
		test  func(actual *User, err error)
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
			test: func(actual *User, err error) {
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
			repository := NewHttpRepository(conf.Conf{
				URL: ts.URL,
			}, false)
			tt.test(repository.Get("uid"))
		})
	}

}
