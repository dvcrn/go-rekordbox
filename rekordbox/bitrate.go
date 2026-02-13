package rekordbox

import (
	"database/sql/driver"
	"encoding/json"
	"math"

	nulltype "github.com/mattn/go-nulltype"
)

// BitRate represents a nullable audio bitrate that supports both integer and float values.
// It provides backward compatibility by exposing both Int64() and Float64() methods.
//
// Migration from old code:
//   Old: content.BitRate.Valid       -> New: content.BitRate.Valid()
//   Old: content.BitRate.Int64       -> New: content.BitRate.Int64()
//   New: content.BitRate.Float64()   -> Supports decimals (e.g., 256.75)
type BitRate struct {
	value nulltype.NullFloat64
}

// Int64 returns the bitrate as an integer (for backward compatibility).
// Decimal values are truncated.
func (b BitRate) Int64() int64 {
	if !b.value.Valid() {
		return 0
	}
	return int64(b.value.Float64Value())
}

// Float64 returns the bitrate as a float64 (supports decimals).
func (b BitRate) Float64() float64 {
	return b.value.Float64Value()
}

// Valid returns true if the BitRate is not null.
func (b BitRate) Valid() bool {
	return b.value.Valid()
}

// BitRateOf creates a new BitRate from a float64 value.
func BitRateOf(value float64) BitRate {
	return BitRate{value: nulltype.NullFloat64Of(value)}
}

// BitRateFromInt creates a new BitRate from an int64 value (backward compatibility).
func BitRateFromInt(value int64) BitRate {
	return BitRate{value: nulltype.NullFloat64Of(float64(value))}
}

// Scan implements the sql.Scanner interface.
func (b *BitRate) Scan(value interface{}) error {
	return b.value.Scan(value)
}

// Value implements the driver.Valuer interface.
func (b BitRate) Value() (driver.Value, error) {
	return b.value.Value()
}

// MarshalJSON implements json.Marshaler interface.
func (b BitRate) MarshalJSON() ([]byte, error) {
	if !b.value.Valid() {
		return []byte("null"), nil
	}
	floatVal := b.value.Float64Value()
	// Marshal as integer if no decimal part, otherwise as float
	if floatVal == math.Floor(floatVal) {
		return json.Marshal(int64(floatVal))
	}
	return json.Marshal(floatVal)
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (b *BitRate) UnmarshalJSON(data []byte) error {
	return b.value.UnmarshalJSON(data)
}
