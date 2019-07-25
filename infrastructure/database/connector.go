package database

import (
	"github.com/jinzhu/gorm"
	// Gorm dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/magrathealabs/jarvis/infrastructure/env"
)

// Connector connects to database configured in environment variables
type Connector interface {
	Variables() map[string]string
	Connect() (*gorm.DB, error)
	Service() string
	Args() string
}

// NewConnector constructor
func NewConnector() Connector {
	switch env.Database().Engine() {
	case "mssql":
		return NewMsSQLConnector()
	case "mysql":
		return NewMySQLConnector()
	case "postgres":
		return NewPostgresConnector()
	case "sqlite3":
		return NewSqlite3Connector()
	}

	return NewSqlite3Connector()
}
