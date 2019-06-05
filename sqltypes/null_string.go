package sqltypes

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

// NullString aliases sql.NullString
type NullString sql.NullString

// MarshalJSON for NullString
func (n NullString) MarshalJSON() ([]byte, error) {
	var a *string
	if n.Valid {
		a = &n.String
	}
	return json.Marshal(a)
}

// UnmarshalJSON for NullString
func (n *NullString) UnmarshalJSON(b []byte) error {
	if bytes.EqualFold(b, nullLiteral) {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.String)
	n.Valid = err == nil
	return err
}

// Value for NullString
func (n NullString) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.String, nil
}

// Scan for NullString
func (n *NullString) Scan(src interface{}) error {
	var a sql.NullString
	if err := a.Scan(src); err != nil {
		return err
	}
	n.String = a.String
	if reflect.TypeOf(src) != nil {
		n.Valid = true
	}
	return nil
}
