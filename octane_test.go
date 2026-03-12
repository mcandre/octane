package octane_test

import (
	"testing"

	"github.com/mcandre/octane"
)

func TestTransposeKeySymmetric(t *testing.T) {
	for k := range uint8(128) {
		kTransposed := octane.TransposeKey(k, 128)

		if kTransposed != k {
			t.Errorf("expected symmetric transpose key for %v", k)
		}
	}
}
