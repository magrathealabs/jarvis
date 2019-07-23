package database

import (
	"github.com/jinzhu/gorm"
	// Gorm dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MySQLConnector implements Connector interface
type MySQLConnector struct{}

// Variables return variables from env
func (connector *MySQLConnector) Variables() map[string]string {
	return make(map[string]string)
}

// Connect to MySQL
func (connector *MySQLConnector) Connect() (*gorm.DB, error) {
	return gorm.Open(connector.Service(), connector.Args())
}

// Service is mysql!
func (connector *MySQLConnector) Service() string {
	return "mysql"
}

// Args to create this connection
func (connector *MySQLConnector) Args() string {
	return "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
}
