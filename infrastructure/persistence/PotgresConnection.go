package persistence

import (
	"fmt"

	"github.com/Wilo/url-shortener-echo/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewPostgresConnection create a new connection to postgres
func NewPostgresConnection() (*gorm.DB, error) {
	user := utils.GetEnv("POSTGRES_USER", "demo")
	pwd := utils.GetEnv("POSTGRES_PASS", "demo")
	port := utils.GetEnv("POSTGRES_PORT", "5432")
	db := utils.GetEnv("POSTGRES_DB", "demo")
	host := utils.GetEnv("POSTGRES_HOST", "localhost")

	connString := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable", user, pwd, host, db, port)

	conn, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, err
}
