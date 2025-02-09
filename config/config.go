package config

import "os"

var JwtSecret = "supersecretkey"

func LoadConfig() {
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		JwtSecret = secret
	}
}
