package polyline

import (
	"fmt"
	"strconv"
	"strings"
)

// Point is a representation of latitude and longitude.
type Point struct {
	Lat float32
	Lng float32
}

// Encode encodes a list of points into a polyline string and
// returns the encoded string.
func Encode(coordinates []Point) string {
	prevLat := float32(0)
	prevLng := float32(0)
	encoded := ""
	for _, point := range coordinates {
		lat := point.Lat
		lng := point.Lng

		encoded += encodeCoordinate(lat, prevLat)
		encoded += encodeCoordinate(lng, prevLng)

		prevLat = lat
		prevLng = lng
	}

	return encoded
}

// encodeCoordinate does the heavy lifting of encoding a single coordinate
// the polyline algorithm goes one by one and considers the difference between
// the current and previous values. Then applies the algorithm to encode the value.
func encodeCoordinate(original, previous float32) string {
	diff := original - previous

	// If current value is equal to previous value
	// Return a question mark
	if diff == 0 {
		return "?"
	}

	// multiply the number by 1e5
	lat := uint32(int(diff * 1e5))
	// left shift one bit
	lat = lat << 1
	// if original is negative invert its encoding
	if original < 0 {
		lat = ^lat
	}

	// Get the binary representation of the number
	bin := fmt.Sprintf("%b", lat)

	// add padding zeros to the string
	if pad := len(bin) % 5; pad > 0 {
		bin = strings.Repeat("0", 5-pad) + bin
	}

	// Split in chunks of 5 bits starting from the right
	chunks := make([]string, 0)
	for i := len(bin); i > 0; i -= 5 {
		chunks = append(chunks, bin[i-5:i])
	}

	encoded := ""
	for i, chunk := range chunks {
		// Parsevalue of the chunk
		n, _ := strconv.ParseUint(chunk, 2, 5)

		// If its not the last chunk OR the value with 0x20
		if i < len(chunks)-1 {
			n = n | 0x20
		}

		// Add 63 to the parsed value
		encoded += string(rune(n + 63))
	}

	return encoded
}
