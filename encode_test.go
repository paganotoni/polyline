package polyline_test

import (
	"testing"

	"github.com/paganotoni/polyline"
)

func Test(t *testing.T) {
	data := []polyline.Point{
		{38.5, -120.2},
		{40.7, -120.95},
		{43.252, -126.453},
	}

	if polyline.Encode(data) != "_p~iF~ps|U_ulLnnqC}lqNvxq`@" {
		t.Errorf("Expected _p~iF~ps|U_ulLnnqC}lqNvxq`@, got %s", polyline.Encode(data))
	}
}
