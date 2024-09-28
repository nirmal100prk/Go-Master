package main

func main() {
NewConfig(WithAddress(""),WithPort(""))
}

type Config struct {
	Address string
	Port    string
	Host    string
}

func NewConfig(opt ...func(*Config)) *Config {
	cfg := Config{}

	for _, o := range opt {
		o(&cfg)
	}
	return &cfg
}

func WithAddress(address string) func(*Config) {
	return func(cfg *Config) {
		cfg.Address = address
	}
}

func WithPort(port string) func(*Config) {
	return func(cfg *Config) {
		cfg.Port = port
	}
}

func WithHost(host string) func(*Config) {
	return func(cfg *Config) {
		cfg.Host = host
	}
}
