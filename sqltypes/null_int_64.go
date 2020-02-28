package sqltypes

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"log"
	"reflect"
)

// NullInt64 aliases sql.NullInt64
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
type NullInt64 sql.NullInt64

// MarshalJSON for NullInt64
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n NullInt64) MarshalJSON() ([]byte, error) {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	var a *int64
	if n.Valid {
		a = &n.Int64
	}
	return json.Marshal(a)
}

// Value for NullInt64
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n NullInt64) Value() (driver.Value, error) {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	if !n.Valid {
		return nil, nil
	}
	return n.Int64, nil
}

// UnmarshalJSON for NullInt64
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n *NullInt64) UnmarshalJSON(b []byte) error {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	if bytes.EqualFold(b, nullLiteral) {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.Int64)
	n.Valid = err == nil
	return err
}

// Scan for NullInt64
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n *NullInt64) Scan(src interface{}) error {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
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
