package polyline_test

import (
	"testing"

	"github.com/paganotoni/polyline"
)

func TestEncode(t *testing.T) {
	t.Run("Standard case", func(t *testing.T) {
		data := []polyline.Point{
			{38.5, -120.2},
			{40.7, -120.95},
			{43.252, -126.453},
		}

		expected := "_p~iF~ps|U_ulLnnqC}lqNvxq`@"
		result := polyline.Encode(data)

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Empty slice", func(t *testing.T) {
		data := []polyline.Point{}
		expected := ""
		result := polyline.Encode(data)

		if result != expected {
			t.Errorf("Expected empty string, got %s", result)
		}
	})

	t.Run("Single point", func(t *testing.T) {
		data := []polyline.Point{
			{41.0, -76.0},
		}

		expected := "_yfyF~vznM"
		result := polyline.Encode(data)

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Same consecutive points", func(t *testing.T) {
		data := []polyline.Point{
			{38.5, -120.2},
			{38.5, -120.2},
			{40.7, -120.95},
		}

		expected := "_p~iF~ps|U??_ulLnnqC"
		result := polyline.Encode(data)

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Negative coordinates", func(t *testing.T) {
		data := []polyline.Point{
			{-38.5, 120.2},
			{-40.7, 120.95},
		}

		expected := "~o~iF_qs|U~tlLonqC"
		result := polyline.Encode(data)

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Zero coordinates", func(t *testing.T) {
		data := []polyline.Point{
			{0, 0},
			{0.0001, 0.0001},
		}

		expected := "??SS"
		result := polyline.Encode(data)

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Large coordinates", func(t *testing.T) {
		data := []polyline.Point{
			{89.99, 179.99},
			{-89.99, -179.99},
		}

		expected := "odgdPohqia@~ioia@~qctcA"
		result := polyline.Encode(data)

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Fractional values", func(t *testing.T) {
		data := []polyline.Point{
			{35.6, -82.55},
			{35.59985, -82.55015},
			{35.6, -82.55},
		}

		expected := "_chxEn`zvN\\]]"
		result := polyline.Encode(data)

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
