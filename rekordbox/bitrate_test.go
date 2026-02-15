package rekordbox

import (
	"testing"

	nulltype "github.com/mattn/go-nulltype"
)

func TestBitRateBackwardCompatibility(t *testing.T) {
	// Test backward compatibility - old code using Int64
	content := &DjmdContent{
		ID:      nulltype.NullStringOf("test123"),
		Title:   nulltype.NullStringOf("Test Track"),
		BitRate: BitRateFromInt(320), // Old style: integer
	}

	if !content.BitRate.Valid() {
		t.Error("Expected BitRate to be valid")
	}

	// Old code still works
	if content.BitRate.Int64() != 320 {
		t.Errorf("Expected BitRate.Int64() to be 320, got %d", content.BitRate.Int64())
	}

	// New code also works
	if content.BitRate.Float64() != 320.0 {
		t.Errorf("Expected BitRate.Float64() to be 320.0, got %f", content.BitRate.Float64())
	}
}

func TestBitRateAsFloat(t *testing.T) {
	// Test that BitRate can handle float values
	content := &DjmdContent{
		ID:      nulltype.NullStringOf("test123"),
		Title:   nulltype.NullStringOf("Test Track"),
		BitRate: BitRateOf(320.5), // New style: float with decimals
	}

	if !content.BitRate.Valid() {
		t.Error("Expected BitRate to be valid")
	}

	if content.BitRate.Float64() != 320.5 {
		t.Errorf("Expected BitRate to be 320.5, got %f", content.BitRate.Float64())
	}

	// Int64() truncates decimals
	if content.BitRate.Int64() != 320 {
		t.Errorf("Expected BitRate.Int64() to be 320 (truncated), got %d", content.BitRate.Int64())
	}
}

func TestBitRateVariousValues(t *testing.T) {
	testCases := []struct {
		name     string
		bitRate  float64
		expected float64
	}{
		{"Integer value", 128.0, 128.0},
		{"Decimal value", 256.75, 256.75},
		{"High bitrate", 1411.2, 1411.2},
		{"Low bitrate", 64.5, 64.5},
		{"Very precise", 192.123456, 192.123456},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			content := &DjmdContent{
				BitRate: BitRateOf(tc.bitRate),
			}

			if content.BitRate.Float64() != tc.expected {
				t.Errorf("Expected BitRate to be %f, got %f", tc.expected, content.BitRate.Float64())
			}
		})
	}
}

func TestBitRateNull(t *testing.T) {
	// Test that BitRate can be null
	content := &DjmdContent{
		ID:      nulltype.NullStringOf("test456"),
		Title:   nulltype.NullStringOf("Test Track Without BitRate"),
		BitRate: BitRate{}, // Null/invalid value
	}

	if content.BitRate.Valid() {
		t.Error("Expected BitRate to be invalid/null")
	}

	// Both accessors return zero for invalid/null
	if content.BitRate.Int64() != 0 {
		t.Errorf("Expected Int64() to return 0 for null, got %d", content.BitRate.Int64())
	}
	if content.BitRate.Float64() != 0.0 {
		t.Errorf("Expected Float64() to return 0.0 for null, got %f", content.BitRate.Float64())
	}
}

func TestBitRateTypeChange(t *testing.T) {
	// This test documents backward compatibility
	// Old code: BitRate.Int64()
	// New code: BitRate.Float64() (supports decimals)
	// Both work!

	content := &DjmdContent{
		BitRate: BitRateOf(320.5),
	}

	// New: decimal values supported
	if content.BitRate.Float64() != 320.5 {
		t.Errorf("BitRate should support decimal values, got %f", content.BitRate.Float64())
	}

	// Old: Int64() still works (truncates)
	if content.BitRate.Int64() != 320 {
		t.Errorf("BitRate.Int64() should work for backward compatibility, got %d", content.BitRate.Int64())
	}
}
