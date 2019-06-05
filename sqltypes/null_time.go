package sqltypes

import (
	"database/sql/driver"
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

// NullTime aliases sql.NullTime
type NullTime mysql.NullTime

// MarshalJSON for NullTime
func (n NullTime) MarshalJSON() ([]byte, error) {
	var a *time.Time
	if n.Valid {
		a = &n.Time
	}
	return json.Marshal(a)
}

// Value for NullTime
func (n NullTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}

// UnmarshalJSON for NullTime
func (n *NullTime) UnmarshalJSON(b []byte) error {
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
func (n *NullTime) Scan(src interface{}) error {
	// Set initial state for subsequent scans.
	n.Valid = false

	var a mysql.NullTime
	if err := a.Scan(src); err != nil {
		return err
	}
	n.Time = a.Time
	if reflect.TypeOf(src) != nil {
		n.Valid = true
	}
	return nil
}
