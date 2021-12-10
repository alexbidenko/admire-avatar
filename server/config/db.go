package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

func getDatabaseHost() string {
	if os.Getenv("MODE") == "production" {
		return "postgres"
	}
	return "localhost"
}

var connStr = "host=" + getDatabaseHost() + " user=postgres password=" + PostgresPassword + " dbname=postgres sslmode=disable"
var DB, DBError = gorm.Open(postgres.New(postgres.Config{
	DSN: connStr,
}), &gorm.Config{})

type Model struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
