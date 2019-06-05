package sqltypes

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

// ToNullString returns a new NullString
func ToNullString(s *string) NullString {
	if s == nil {
		return NullString(sql.NullString{Valid: false})
	}
	return NullString(sql.NullString{String: *s, Valid: true})
}

// ToNullInt64 returns a new NullInt64
func ToNullInt64(i *int64) NullInt64 {
	if i == nil {
		return NullInt64(sql.NullInt64{Valid: false})
	}
	return NullInt64(sql.NullInt64{Int64: *i, Valid: true})
}

// ToNullFloat64 returns a new NullFloat64
func ToNullFloat64(i *float64) NullFloat64 {
	if i == nil {
		return NullFloat64(sql.NullFloat64{Valid: false})
	}
	return NullFloat64(sql.NullFloat64{Float64: *i, Valid: true})
}

// ToNullBool creates a new NullBool
func ToNullBool(b *bool) NullBool {
	if b == nil {
		return NullBool(sql.NullBool{Valid: false})
	}
	return NullBool(sql.NullBool{Bool: *b, Valid: true})
}

// ToNullTime creates a new NullTime
func ToNullTime(t time.Time) NullTime {
	if t == emptyTime {
		return NullTime(mysql.NullTime{Valid: false})
	}
	return NullTime(mysql.NullTime{Time: t, Valid: true})
}
