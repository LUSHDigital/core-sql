package sqltypes

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

// NullInt64 aliases sql.NullInt64
type NullInt64 sql.NullInt64

// MarshalJSON for NullInt64
func (n NullInt64) MarshalJSON() ([]byte, error) {
	var a *int64
	if n.Valid {
		a = &n.Int64
	}
	return json.Marshal(a)
}

// Value for NullInt64
func (n NullInt64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int64, nil
}

// UnmarshalJSON for NullInt64
func (n *NullInt64) UnmarshalJSON(b []byte) error {
	if bytes.EqualFold(b, nullLiteral) {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.Int64)
	n.Valid = err == nil
	return err
}

// Scan for NullInt64
func (n *NullInt64) Scan(src interface{}) error {
	// Set initial state for subsequent scans.
	n.Valid = false

	var a sql.NullInt64
	if err := a.Scan(src); err != nil {
		return err
	}
	n.Int64 = a.Int64
	if reflect.TypeOf(src) != nil {
		n.Valid = true
	}
	return nil
}
