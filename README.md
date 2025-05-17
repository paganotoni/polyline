# Polyline

This repo is a Go implementation of the Google Maps Polyline encoding algorithm. It supports both encoding and decoding of polyline strings.

For more information, visit Google's explanation of the algorithm [here](https://developers.google.com/maps/documentation/utilities/polylinealgorithm).

## How to Use

```go
package main

import (
	"fmt"
	"github.com/paganotoni/polyline"
)

func main() {
	// Encoding example
	points := []polyline.Point{
		{Lat: 38.5, Lng: -120.2},
		{Lat: 40.7, Lng: -120.95},
		{Lat: 43.252, Lng: -126.453},
	}
	encoded := polyline.Encode(points)
	fmt.Println("Encoded polyline:", encoded)

	// Decoding example
	decodedPoints, err := polyline.Decode(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded points:", decodedPoints)
}
```
