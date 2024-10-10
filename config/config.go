package config

import "os"

type Config struct {
	BackendURL string
	CertFile   string
	KeyFile    string
	Address    string
}

func LoadConfig() *Config {
	backendUrl := getEnv("BACKEND_URL", "http://localhost:8080")
	certFile := getEnv("CERT_FILE", "cert.pem")
	keyFile := getEnv("KEY_FILE", "key.pem")
	address := getEnv("ADDRESS", ":8443")

	return &Config{
		BackendURL: backendUrl,
		CertFile:   certFile,
		KeyFile:    keyFile,
		Address:    address,
	}
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
