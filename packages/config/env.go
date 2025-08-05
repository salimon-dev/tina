package config

import (
	"os"
)

func GetUsername() string {
	return os.Getenv("NEXUS_USERNAME")
}
