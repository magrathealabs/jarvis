package database

import (
	"github.com/jinzhu/gorm"
	// Gorm dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Sqlite3Connector implements Connector interface
type Sqlite3Connector struct{}

// NewSqlite3Connector construct a connector for Sqlite3
func NewSqlite3Connector() Connector {
	return &Sqlite3Connector{}
}

// Variables return variables from env
func (connector *Sqlite3Connector) Variables() map[string]string {
	return make(map[string]string)
}

// Connect to Sqlite3
func (connector *Sqlite3Connector) Connect() (*gorm.DB, error) {
	return gorm.Open(connector.Service(), connector.Args())
}

// Service is Sqlite3!
func (connector *Sqlite3Connector) Service() string {
	return "sqlite3"
}

// Args to create this connection
func (connector *Sqlite3Connector) Args() string {
	return "/tmp/gorm.db"
}
