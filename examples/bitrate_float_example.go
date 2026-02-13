package main

import (
	"fmt"

	"github.com/dvcrn/go-rekordbox/rekordbox"
	nulltype "github.com/mattn/go-nulltype"
)

func main() {
	fmt.Println("BitRate Float Example")
	fmt.Println("=====================")
	fmt.Println()

	// Example 1: Creating content with integer bitrate (backward compatible)
	fmt.Println("Example 1: Integer bitrate (e.g., 320 kbps)")
	content1 := &rekordbox.DjmdContent{
		ID:      nulltype.NullStringOf("track001"),
		Title:   nulltype.NullStringOf("High Quality Track"),
		BitRate: rekordbox.BitRateFromInt(320), // Old style still works!
	}
	fmt.Printf("  Track: %s\n", content1.Title.String)
	fmt.Printf("  BitRate (Int64): %d kbps\n", content1.BitRate.Int64())
	fmt.Printf("  BitRate (Float64): %.1f kbps\n", content1.BitRate.Float64())
	fmt.Println()

	// Example 2: Creating content with decimal bitrate (new capability!)
	fmt.Println("Example 2: Decimal bitrate (e.g., 256.5 kbps)")
	content2 := &rekordbox.DjmdContent{
		ID:      nulltype.NullStringOf("track002"),
		Title:   nulltype.NullStringOf("Variable Bitrate Track"),
		BitRate: rekordbox.BitRateOf(256.75), // New style with decimals
	}
	fmt.Printf("  Track: %s\n", content2.Title.String)
	fmt.Printf("  BitRate (Float64): %.2f kbps\n", content2.BitRate.Float64())
	fmt.Printf("  BitRate (Int64): %d kbps (truncated)\n", content2.BitRate.Int64())
	fmt.Println()

	// Example 3: FLAC/Lossless with high bitrate
	fmt.Println("Example 3: High-quality lossless (e.g., 1411.2 kbps)")
	content3 := &rekordbox.DjmdContent{
		ID:      nulltype.NullStringOf("track003"),
		Title:   nulltype.NullStringOf("FLAC Track"),
		BitRate: rekordbox.BitRateOf(1411.2),
	}
	fmt.Printf("  Track: %s\n", content3.Title.String)
	fmt.Printf("  BitRate: %.1f kbps\n", content3.BitRate.Float64())
	fmt.Println()

	// Example 4: Null/missing bitrate
	fmt.Println("Example 4: Track without bitrate information")
	content4 := &rekordbox.DjmdContent{
		ID:      nulltype.NullStringOf("track004"),
		Title:   nulltype.NullStringOf("Unknown Bitrate Track"),
		BitRate: rekordbox.BitRate{}, // Not set
	}
	fmt.Printf("  Track: %s\n", content4.Title.String)
	if content4.BitRate.Valid() {
		fmt.Printf("  BitRate: %.1f kbps\n", content4.BitRate.Float64())
	} else {
		fmt.Println("  BitRate: Not available")
	}
	fmt.Println()

	// Example 5: Checking bitrate and formatting output
	fmt.Println("Example 5: Processing multiple tracks")
	tracks := []*rekordbox.DjmdContent{
		{
			Title:   nulltype.NullStringOf("MP3 Standard"),
			BitRate: rekordbox.BitRateFromInt(128),
		},
		{
			Title:   nulltype.NullStringOf("MP3 High Quality"),
			BitRate: rekordbox.BitRateFromInt(320),
		},
		{
			Title:   nulltype.NullStringOf("AAC Variable"),
			BitRate: rekordbox.BitRateOf(256.5),
		},
		{
			Title:   nulltype.NullStringOf("FLAC Lossless"),
			BitRate: rekordbox.BitRateOf(1411.2),
		},
	}

	for i, track := range tracks {
		if track.BitRate.Valid() {
			quality := getQualityLabel(track.BitRate.Float64())
			fmt.Printf("  %d. %-20s: %7.1f kbps (%s)\n",
				i+1,
				track.Title.String,
				track.BitRate.Float64(),
				quality)
		}
	}
	fmt.Println()

	// Example 6: Why float matters for variable bitrate
	fmt.Println("Example 6: Variable Bitrate (VBR) average")
	fmt.Println("  VBR tracks often have non-integer average bitrates")
	vbrTrack := &rekordbox.DjmdContent{
		Title:   nulltype.NullStringOf("VBR Encoded Track"),
		BitRate: rekordbox.BitRateOf(245.67), // Precise VBR average
	}
	fmt.Printf("  Track: %s\n", vbrTrack.Title.String)
	fmt.Printf("  Average BitRate: %.2f kbps (VBR)\n", vbrTrack.BitRate.Float64())
	fmt.Println()

	fmt.Println("Summary:")
	fmt.Println("--------")
	fmt.Println("✓ BitRate supports decimal values: BitRateOf(256.75)")
	fmt.Println("✓ Backward compatible: BitRateFromInt(320) or .Int64()")
	fmt.Println("✓ New code uses .Float64() for precision")
	fmt.Println("✓ Old code using .Int64() still works!")
}

// getQualityLabel returns a quality label based on bitrate
func getQualityLabel(bitrate float64) string {
	switch {
	case bitrate >= 1000:
		return "Lossless"
	case bitrate >= 256:
		return "High Quality"
	case bitrate >= 192:
		return "Good Quality"
	case bitrate >= 128:
		return "Standard"
	default:
		return "Low Quality"
	}
}
