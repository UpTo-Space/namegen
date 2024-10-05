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

func TestGenerateLength(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	nameGen := NewNameGen(seed)

	onlyName := nameGen.GenerateLength(1)

	defer func() {
		for _, name := range names {
			if onlyName == name {
				return
			}
		}
		t.Errorf("Name %s is not in names", onlyName)
	}()

	assertPanicNameGen(t, nameGen.GenerateLength)

	longName := nameGen.GenerateLength(4)
	if len(strings.Split(longName, "_")) != 4 {
		t.Errorf("Length of generated name is wrong")
	}
}

func assertPanicNameGen(t *testing.T, f func(int) string) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f(-1)
}
