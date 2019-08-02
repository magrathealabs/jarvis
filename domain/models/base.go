package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/krakenlab/ternary"
)

// Base for other models
type Base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Errors []error `json:"errors"`
}

// NewBase constructor
func NewBase() *Base {
	return &Base{}
}

// AppendError in model
func (base *Base) AppendError(errs ...error) {
	base.Errors = append(base.Errors, errs...)
}

// Valid verify errors list
func (base *Base) Valid() bool {
	for _, err := range base.Errors {
		if err != nil {
			return false
		}
	}

	return true
}

// ToJSON marshal
func (base *Base) ToJSON() string {
	data, err := json.Marshal(base)
	ternary.Func(err == nil, func() {}, func() { panic(err) })()
	return string(data)
}

func (base *Base) clearMetricTag(tag string) string {
	tag = strings.Replace(tag, "/", "_", -1)
	tag = strings.TrimPrefix(tag, "_")
	tag = strings.TrimSuffix(tag, "_")
	tag = ternary.String(tag == "", "root", tag)

	return tag
}
