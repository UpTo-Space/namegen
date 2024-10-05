// Package namegen implements a Name Generator that can be used to generate
// random names for different things like servers, containers, connections etc.
package namegen

import (
	"fmt"
	"math/rand"
)

type Generator interface {
	Generate() string
	GenerateLength(int) string
}

type NameGen struct {
	random *rand.Rand
}

// Create a new Generator object
func NewNameGen(seed int64) Generator {
	gen := &NameGen{
		random: rand.New(rand.NewSource(seed)),
	}

	return gen
}

// Generate a random name with default length.
// It is built out of two adjectives and one name.
func (gen *NameGen) Generate() string {
	return gen.GenerateLength(3)
}

// Generate a random name that ends with a name.
// If length is greated than one, multiple adjectives are added before the name.
func (gen *NameGen) GenerateLength(length int) string {
	if length < 0 {
		panic(fmt.Errorf("length must be greater than 0"))
	}
	ret := names[gen.random.Intn(len(names))]
	for i := 1; i < length; i++ {
		ret = fmt.Sprintf("%v_%v", adjectives[gen.random.Intn(len(adjectives))], ret)
	}
	return ret
}
