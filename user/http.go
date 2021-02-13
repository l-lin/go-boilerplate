package user

import (
	"encoding/json"
	"github.com/l-lin/go-boilerplate/conf"
	wrappedhttp "github.com/l-lin/go-boilerplate/http"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

func NewHttpRepository(c conf.Conf, debug bool) Repository {
	return HttpRepository{Conf: c, client: wrappedhttp.NewClient(debug)}
}

type HttpRepository struct {
	conf.Conf
	client *wrappedhttp.Client
}

func (h HttpRepository) Get(_ string) (*User, error) {
	url := h.URL + "/uuid"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	log.Debug().Str("url", url).Msg("getting user")
	response, err := h.client.DoOnlyOk(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	var u *User
	if err = json.Unmarshal(body, &u); err != nil {
		return nil, err
	}
	return u, nil
}
