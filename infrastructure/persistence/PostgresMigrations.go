package persistence

import (
	app "github.com/Wilo/url-shortener-echo/application"
	"github.com/Wilo/url-shortener-echo/domain/models"
	"gorm.io/gorm"
)

// PostgresMigrations migrator with postgres
type PostgresMigrations struct {
	connection *gorm.DB
}

// NewPostgresMigrations constructor for PostgresMigrations
func NewPostgresMigrations(connection *gorm.DB) app.PersistenceMigration {
	return &PostgresMigrations{
		connection: connection,
	}
}

// Migrate implementation
func (m *PostgresMigrations) Migrate() error {
	m.connection.AutoMigrate(
		&models.Url{},
	)
	return nil
}
