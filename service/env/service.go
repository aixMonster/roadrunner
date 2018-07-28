package env

// ID contains default svc name.
const ID = "env"

// Service provides ability to map _ENV values from config file.
type Service struct {
	// values is default set of values.
	values map[string]string
}

// NewService creates new env service instance for given rr version.
func NewService(version string) *Service {
	return &Service{values: map[string]string{"rr": version}}
}

// Init must return configure svc and return true if svc hasStatus enabled. Must return error in case of
// misconfiguration. Services must not be used without proper configuration pushed first.
func (s *Service) Init(cfg *Config) (bool, error) {
	for k, v := range cfg.Values {
		s.values[k] = v
	}

	return true, nil
}

// GetEnv must return list of env variables.
func (s *Service) GetEnv() (map[string]string, error) {
	return s.values, nil
}