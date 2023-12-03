package utils

import "os"

func GetEnv() string {
	return os.Getenv("APP_ENV")
}