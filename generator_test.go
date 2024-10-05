package namegen

import (
	"strings"
	"testing"
	"time"
)

func TestNameGen(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	_ = NewNameGen(seed)
}

func TestGeneate(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	name := NewNameGen(seed).Generate()
	if len(name) == 0 {
		t.Errorf("Name is empty")
	}
	parts := strings.Split(name, "_")

	if len(parts) != 3 {
		t.Errorf("Name doesn't match expected length")
	}
}
