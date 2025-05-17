package polyline

import (
	"errors"
)

// Decode decodes a polyline string into a slice of Point structs.
// Returns an error if the polyline is invalid.
func Decode(poly string) ([]Point, error) {
	var points []Point
	var lat, lng int32
	index := 0
	length := len(poly)

	for index < length {
		dlat, n, err := decodeCoordinate(poly, index)
		if err != nil {
			return nil, err
		}
		index = n
		lat += dlat

		dlng, n, err := decodeCoordinate(poly, index)
		if err != nil {
			return nil, err
		}
		index = n
		lng += dlng

		points = append(points, Point{
			Lat: float32(lat) / 1e5,
			Lng: float32(lng) / 1e5,
		})
	}

	return points, nil
}

// decodeCoordinate decodes a single coordinate from the polyline string.
func decodeCoordinate(poly string, start int) (int32, int, error) {
	var result, shift, b int32
	index := start

	for {
		if index >= len(poly) {
			return 0, index, errors.New("polyline decode: truncated string")
		}
		b = int32(poly[index]) - 63
		index++
		result |= (b & 0x1F) << shift
		shift += 5
		if b < 0x20 {
			break
		}
	}

	coord := result >> 1
	if result&1 != 0 {
		coord = ^coord
	}

	return coord, index, nil
}
