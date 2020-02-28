package sqltypes

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"log"
)

// RawJSON aliases json.RawMessage
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
type RawJSON json.RawMessage

// MarshalJSON for NullString
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n RawJSON) MarshalJSON() ([]byte, error) {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	if len(n) == 0 {
		return []byte("null"), nil
	}
	a := json.RawMessage(n)
	return a.MarshalJSON()
}

// Value for NullString
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n RawJSON) Value() (driver.Value, error) {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	return string(n), nil
}

// UnmarshalJSON for NullString
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n *RawJSON) UnmarshalJSON(b []byte) error {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	var a json.RawMessage
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	c := RawJSON(a)
	*n = c
	return nil
}

// Scan for NullString
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n *RawJSON) Scan(src interface{}) error {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	var a sql.NullString
	if err := a.Scan(src); err != nil {
		return err
	}
	jsn := RawJSON([]byte(a.String))
	*n = jsn
	return nil
}
