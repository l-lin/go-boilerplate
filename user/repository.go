package user

type Repository interface {
	Get(userId string) (*User, error)
}
