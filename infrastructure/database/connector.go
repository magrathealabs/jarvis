package database

import (
	"github.com/jinzhu/gorm"
	// Gorm dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connector connects to database configured in environment variables
type Connector interface {
	Variables() map[string]string
	Connect() (*gorm.DB, error)
	Service() string
	Args() string
}
