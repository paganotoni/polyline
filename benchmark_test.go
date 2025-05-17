package polyline_test

import (
	"testing"

	"github.com/paganotoni/polyline"
)

func BenchmarkEncode(b *testing.B) {
	points := []polyline.Point{
		{Lat: 38.5, Lng: -120.2},
		{Lat: 40.7, Lng: -120.95},
		{Lat: 43.252, Lng: -126.453},
	}
	for b.Loop() {
		_ = polyline.Encode(points)
	}
}

func BenchmarkDecode(b *testing.B) {
	encoded := "_p~iF~ps|U_ulLnnqC}lqNvxq`@"
	for b.Loop() {
		_, err := polyline.Decode(encoded)
		if err != nil {
			b.Fatalf("Decode failed: %v", err)
		}
	}
}
