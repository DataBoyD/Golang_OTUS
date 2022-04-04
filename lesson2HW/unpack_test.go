package lesson2

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	example := ""
	if Unpack(example) != "" {
		t.Errorf("OK")

	}
}

func TestUnpackError(t *testing.T) {
	example := "45"
	if Unpack(example) != "" {
		t.Errorf("OK")

	}
}
