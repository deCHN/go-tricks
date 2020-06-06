package main_test

import (
	"fmt"
	"testing"
)

func TestClosures(t *testing.T) {
	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}
}