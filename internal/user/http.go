package user

import (
	"context"

	"github.com/carlmjohnson/requests"
	"go-boilerplate/internal/config"
	"go-boilerplate/pkg/user"
)

func NewHttpRepository(c config.Config) Repository {
	return HttpRepository{Config: c}
}

type HttpRepository struct {
	config.Config
}

func (h HttpRepository) Get(_ string) (*user.User, error) {
	url := h.URL + "/uuid"

	var u *user.User
	if err := requests.
		URL(url).
		ToJSON(&u).
		Fetch(context.Background()); err != nil {
		return nil, err
	}

	return u, nil
}
