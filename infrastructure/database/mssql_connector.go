package database

import (
	"github.com/jinzhu/gorm"
	// Gorm dialect
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// MsSQLConnector implements Connector interface
type MsSQLConnector struct{}

// Variables return variables from env
func (connector *MsSQLConnector) Variables() map[string]string {
	return make(map[string]string)
}

// Connect to MsSQL
func (connector *MsSQLConnector) Connect() (*gorm.DB, error) {
	return gorm.Open(connector.Service(), connector.Args())
}

// Service is MsSQL!
func (connector *MsSQLConnector) Service() string {
	return "mssql"
}

// Args to create this connection
func (connector *MsSQLConnector) Args() string {
	return "sqlserver://username:password@localhost:1433?database=dbname"
}
