package coresql

import "os"

const (
	// DefaultMySQLURL is the default url to a MYSQL database.
	DefaultMySQLURL = "tcp(127.0.0.1:3306)/service"
)

// MySQLURLFromEnv tries to retrieve the redis url from the environment.
func MySQLURLFromEnv() string {
	url := os.Getenv("MYSQL_URL")
	if url == "" {
		url = DefaultMySQLURL
	}
	return url
}
