package sqltypes

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"log"
	"reflect"
)

// NullString aliases sql.NullString
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
type NullString sql.NullString

// MarshalJSON for NullString
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n NullString) MarshalJSON() ([]byte, error) {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	var a *string
	if n.Valid {
		a = &n.String
	}
	return json.Marshal(a)
}

// UnmarshalJSON for NullString
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n *NullString) UnmarshalJSON(b []byte) error {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	if bytes.EqualFold(b, nullLiteral) {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.String)
	n.Valid = err == nil
	return err
}

// Value for NullString
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n NullString) Value() (driver.Value, error) {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	if !n.Valid {
		return nil, nil
	}
	return n.String, nil
}

// Scan for NullString
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n *NullString) Scan(src interface{}) error {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
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
