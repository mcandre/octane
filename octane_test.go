package octane_test

import (
	"github.com/mcandre/octane"

	"testing"
)

func TestTransposeKeySymmetric(t *testing.T) {
	for k := uint8(0); k < 128; k++ {
		kTransposed := octane.TransposeKey(k, 128)

		if kTransposed != k {
			t.Errorf("expected symmetric transpose key for %v", k)
		}
	}
}
