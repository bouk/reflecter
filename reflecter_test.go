package reflecter

import (
	"testing"
)

func TestReflecter(t *testing.T) {
	for _, typ := range Types() {
		typ.String()
	}
}
