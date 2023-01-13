package config

import "fmt"

const (
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "veronica"
	DBHost     = "localhost"
	DBPort     = "5432"
)

func GetSqlConnectionString() string {
	database := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
	return database
}
