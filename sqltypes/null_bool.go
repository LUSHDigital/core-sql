package sqltypes

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

// NullBool aliases sql.NullBool
type NullBool sql.NullBool

// MarshalJSON for NullBool
func (n NullBool) MarshalJSON() ([]byte, error) {
	var a *bool
	if n.Valid {
		a = &n.Bool
	}
	return json.Marshal(a)
}

// Value for NullBool
func (n NullBool) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Bool, nil
}

// UnmarshalJSON for NullBool
func (n *NullBool) UnmarshalJSON(b []byte) error {
	var field *bool
	err := json.Unmarshal(b, &field)
	if field != nil {
		n.Valid = true
		n.Bool = *field
	}
	return err
}

// Scan for NullBool
func (n *NullBool) Scan(src interface{}) error {
	var a sql.NullBool
	if err := a.Scan(src); err != nil {
		return err
	}
	n.Bool = a.Bool
	if reflect.TypeOf(src) != nil {
		n.Valid = true
	}
	return nil
}
