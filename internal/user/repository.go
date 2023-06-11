package user

import "go-boilerplate/pkg/user"

type Repository interface {
	Get(userId string) (*user.User, error)
}
