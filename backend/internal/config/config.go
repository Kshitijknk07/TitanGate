type BackendConfig struct {
	URL    string `env:"BACKEND_URLS" envSeparator:","`
	Weight int    `env:"BACKEND_WEIGHTS" envSeparator:","`
}

func LoadBackendConfig() []loadbalancer.Backend {
	urls := strings.Split(os.Getenv("BACKEND_URLS"), ",")
	weights := strings.Split(os.Getenv("BACKEND_WEIGHTS"), ",")

	backends := make([]loadbalancer.Backend, len(urls))
	for i, url := range urls {
		weight := 1
		if i < len(weights) {
			weight, _ = strconv.Atoi(weights[i])
		}
		backends[i] = loadbalancer.Backend{
			URL:    strings.TrimSpace(url),
			Weight: weight,
			Active: true,
		}
	}
	return backends
}