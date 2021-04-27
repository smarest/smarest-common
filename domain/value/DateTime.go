package value

import (
	"database/sql/driver"
	"fmt"
	"time"
)

func NewDateTime() DateTime {
	return DateTime{time.Now().Local()}
}

type DateTime struct {
	time.Time
}

func (t DateTime) ToString() string {
	return t.Time.Format(time.RFC3339)
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.ToString())
	return []byte(stamp), nil
}

func (t DateTime) MarshalText() ([]byte, error) {
	return t.MarshalJSON()
}

// Sql driver interface

func (t *DateTime) Scan(value interface{}) error {
	t.Time = value.(time.Time).Local()
	return nil
}

func (t DateTime) Value() (driver.Value, error) {
	return t.ToString(), nil
}
