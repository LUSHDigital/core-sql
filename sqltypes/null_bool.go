package sqltypes

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"log"
	"reflect"
)

// NullBool aliases sql.NullBool
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
type NullBool sql.NullBool

// MarshalJSON for NullBool
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n NullBool) MarshalJSON() ([]byte, error) {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	var a *bool
	if n.Valid {
		a = &n.Bool
	}
	return json.Marshal(a)
}

// Value for NullBool
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n NullBool) Value() (driver.Value, error) {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	if !n.Valid {
		return nil, nil
	}
	return n.Bool, nil
}

// UnmarshalJSON for NullBool
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n *NullBool) UnmarshalJSON(b []byte) error {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	var field *bool
	err := json.Unmarshal(b, &field)
	if field != nil {
		n.Valid = true
		n.Bool = *field
	}
	return err
}

// Scan for NullBool
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n *NullBool) Scan(src interface{}) error {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
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
