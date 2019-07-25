package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// Gorm dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/magrathealabs/jarvis/infrastructure/env"
)

// PostgresConnector implements Connector interface
type PostgresConnector struct {
	Host   string
	Port   string
	User   string
	Dbname string
	Pass   string
}

// NewPostgresConnector construct a connector for Postgres
func NewPostgresConnector() Connector {
	return &PostgresConnector{
		Host:   env.Postgres().Host(),
		Port:   env.Postgres().Port(),
		User:   env.Postgres().User(),
		Dbname: env.Postgres().Dbname(),
		Pass:   env.Postgres().Pass(),
	}
}

// Variables return variables from env
func (connector *PostgresConnector) Variables() map[string]string {
	return map[string]string{
		"Host":   connector.Host,
		"Port":   connector.Port,
		"User":   connector.User,
		"Dbname": connector.Dbname,
		"Pass":   connector.Pass,
	}
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
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		connector.Host,
		connector.Port,
		connector.User,
		connector.Dbname,
		connector.Pass,
	)
}
