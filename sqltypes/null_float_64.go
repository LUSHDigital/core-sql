package sqltypes

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

// NullFloat64 aliases sql.NullFloat64
type NullFloat64 sql.NullFloat64

// MarshalJSON for NullFloat64
func (n NullFloat64) MarshalJSON() ([]byte, error) {
	var a *float64
	if n.Valid {
		a = &n.Float64
	}
	return json.Marshal(a)
}

// Value for NullFloat64
func (n NullFloat64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Float64, nil
}

// UnmarshalJSON for NullFloat64
func (n *NullFloat64) UnmarshalJSON(b []byte) error {
	if bytes.EqualFold(b, nullLiteral) {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.Float64)
	n.Valid = err == nil
	return err
}

// Scan for NullFloat64
func (n *NullFloat64) Scan(src interface{}) error {
	var a sql.NullFloat64
	if err := a.Scan(src); err != nil {
		return err
	}
	n.Float64 = a.Float64
	if reflect.TypeOf(src) != nil {
		n.Valid = true
	}
	return nil
}
