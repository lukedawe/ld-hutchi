package util

import (
	"fmt"
)

func GetDsn(host string, port uint, user string, password string, database string) string {
	return fmt.Sprintf(
		// TODO: sslmode needs to be configurable
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		database,
	)
}
