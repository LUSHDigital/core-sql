package sqltypes

import (
	"database/sql"
	"log"
	"time"
)

// ToNullString returns a new NullString
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func ToNullString(s *string) NullString {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	if s == nil {
		return NullString(sql.NullString{Valid: false})
	}
	return NullString(sql.NullString{String: *s, Valid: true})
}

// ToNullInt64 returns a new NullInt64
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func ToNullInt64(i *int64) NullInt64 {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	if i == nil {
		return NullInt64(sql.NullInt64{Valid: false})
	}
	return NullInt64(sql.NullInt64{Int64: *i, Valid: true})
}

// ToNullFloat64 returns a new NullFloat64
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func ToNullFloat64(i *float64) NullFloat64 {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	if i == nil {
		return NullFloat64(sql.NullFloat64{Valid: false})
	}
	return NullFloat64(sql.NullFloat64{Float64: *i, Valid: true})
}

// ToNullBool creates a new NullBool
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func ToNullBool(b *bool) NullBool {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	if b == nil {
		return NullBool(sql.NullBool{Valid: false})
	}
	return NullBool(sql.NullBool{Bool: *b, Valid: true})
}

// ToNullTime creates a new NullTime
// Deprecated: consinder using github.com/LUSHDigital/core-lush/nullable
func ToNullTime(t time.Time) NullTime {
	log.Println("package sqltypes is deprecated: consinder using github.com/LUSHDigital/core-lush/nullable")
	if t == emptyTime {
		return NullTime(sql.NullTime{Valid: false})
	}
	return NullTime(sql.NullTime{Time: t, Valid: true})
}
