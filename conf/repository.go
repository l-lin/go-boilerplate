package conf

// Repository to persist the app config
type Repository interface {
	Init() error
	Get() *Conf
	Save(*Conf) error
}
