package database

import (
	"github.com/jinzhu/gorm"
	// Gorm dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// PostgresConnector implements Connector interface
type PostgresConnector struct{}

// NewPostgresConnector construct a connector for Postgres
func NewPostgresConnector() Connector {
	return &PostgresConnector{}
}

// Variables return variables from env
func (connector *PostgresConnector) Variables() map[string]string {
	return make(map[string]string)
}

// Connect to Postgres
func (connector *PostgresConnector) Connect() (*gorm.DB, error) {
	return gorm.Open(connector.Service(), connector.Args())
}

// Service is Postgres!
func (connector *PostgresConnector) Service() string {
	return "postgres"
}

// Args to create this connection
func (connector *PostgresConnector) Args() string {
	return "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
}
