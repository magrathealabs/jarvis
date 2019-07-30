package models

import (
	"encoding/json"
	"time"

	"github.com/krakenlab/ternary"
)

// Base for other models
type Base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// NewBase constructor
func NewBase() *Base {
	return &Base{}
}

// ToJSON marshal
func (base *Base) ToJSON() string {
	data, err := json.Marshal(base)
	ternary.Func(err == nil, func() {}, func() { panic(err) })()
	return string(data)
}
