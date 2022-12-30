package utils

import (
	"fmt"
	"os"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	var url string
	switch n {
	case "postgres": //postgres connection
		url = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Jakarta",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"))
	case "mysql":
		url = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"))
	case "redis":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"))
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	return url, nil
}
