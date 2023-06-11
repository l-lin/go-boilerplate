package config

// Repository to persist the app config
type Repository interface {
	Get() (*Config, error)
	Save(*Config) error
}
