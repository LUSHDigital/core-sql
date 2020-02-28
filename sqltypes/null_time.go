package sqltypes

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"log"
	"reflect"
	"strings"
	"time"
)

// NullTime aliases sql.NullTime
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
type NullTime sql.NullTime

// MarshalJSON for NullTime
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n NullTime) MarshalJSON() ([]byte, error) {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	var a *time.Time
	if n.Valid {
		a = &n.Time
	}
	return json.Marshal(a)
}

// Value for NullTime
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n NullTime) Value() (driver.Value, error) {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}

// UnmarshalJSON for NullTime
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n *NullTime) UnmarshalJSON(b []byte) error {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	s := string(b)
	s = strings.Trim(s, `"`)

	var (
		zeroTime time.Time
		tim      time.Time
		err      error
	)

	if strings.EqualFold(s, "null") {
		return nil
	}

	if tim, err = time.Parse(time.RFC3339, s); err != nil {
		n.Valid = false
		return err
	}

	if tim == zeroTime {
		return nil
	}

	n.Time = tim
	n.Valid = true
	return nil
}

// Scan for NullTime
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func (n *NullTime) Scan(src interface{}) error {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	// Set initial state for subsequent scans.
	n.Valid = false

	var a sql.NullTime
	if err := a.Scan(src); err != nil {
		return err
	}
	n.Time = a.Time
	if reflect.TypeOf(src) != nil {
		n.Valid = true
	}
	return nil
}
