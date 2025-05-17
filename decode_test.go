package polyline_test

import (
	"testing"

	"github.com/paganotoni/polyline"
)

func TestDecode(t *testing.T) {
	t.Run("Standard case", func(t *testing.T) {
		encoded := "_p~iF~ps|U_ulLnnqC}lqNvxq`@"
		expected := []polyline.Point{
			{38.5, -120.2},
			{40.7, -120.95},
			{43.252, -126.453},
		}
		points, err := polyline.Decode(encoded)
		if err != nil {
			t.Fatalf("Decode failed: %v", err)
		}
		if len(points) != len(expected) {
			t.Fatalf("Expected %d points, got %d", len(expected), len(points))
		}
		for i := range points {
			if abs(points[i].Lat-expected[i].Lat) > 1e-5 || abs(points[i].Lng-expected[i].Lng) > 1e-5 {
				t.Errorf("Point %d: expected %+v, got %+v", i, expected[i], points[i])
			}
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		encoded := ""
		points, err := polyline.Decode(encoded)
		if err != nil {
			t.Fatalf("Decode failed: %v", err)
		}
		if len(points) != 0 {
			t.Errorf("Expected 0 points, got %d", len(points))
		}
	})

	t.Run("Single point", func(t *testing.T) {
		encoded := "_yfyF~vznM"
		expected := []polyline.Point{
			{41.0, -76.0},
		}
		points, err := polyline.Decode(encoded)
		if err != nil {
			t.Fatalf("Decode failed: %v", err)
		}
		if len(points) != len(expected) {
			t.Fatalf("Expected %d points, got %d", len(expected), len(points))
		}
		for i := range points {
			if abs(points[i].Lat-expected[i].Lat) > 1e-5 || abs(points[i].Lng-expected[i].Lng) > 1e-5 {
				t.Errorf("Point %d: expected %+v, got %+v", i, expected[i], points[i])
			}
		}
	})

	t.Run("Negative coordinates", func(t *testing.T) {
		encoded := "~o~iF_qs|U~tlLonqC"
		expected := []polyline.Point{
			{-38.5, 120.2},
			{-40.7, 120.95},
		}
		points, err := polyline.Decode(encoded)
		if err != nil {
			t.Fatalf("Decode failed: %v", err)
		}
		if len(points) != len(expected) {
			t.Fatalf("Expected %d points, got %d", len(expected), len(points))
		}
		for i := range points {
			if abs(points[i].Lat-expected[i].Lat) > 1e-5 || abs(points[i].Lng-expected[i].Lng) > 1e-5 {
				t.Errorf("Point %d: expected %+v, got %+v", i, expected[i], points[i])
			}
		}
	})

	t.Run("Invalid polyline", func(t *testing.T) {
		encoded := "_p~iF~ps|U_ulLnnqC}lqNvxq" // truncated
		_, err := polyline.Decode(encoded)
		if err == nil {
			t.Errorf("Expected error for truncated polyline, got nil")
		}
	})
}

// abs returns the absolute value of a float32.
func abs(f float32) float32 {
	if f < 0 {
		return -f
	}
	return f
}
