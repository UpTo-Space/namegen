package namegen

import (
	"fmt"
	"math/rand"
)

type Generator interface {
	Generate() string
}

type NameGen struct {
	random *rand.Rand
}

func (gen *NameGen) Generate() string {
	firstAjective := adjectives[gen.random.Intn(len(adjectives))]
	secondAdjective := adjectives[gen.random.Intn(len(adjectives))]
	name := names[gen.random.Intn(len(names))]

	return fmt.Sprintf("%s_%s_%s", firstAjective, secondAdjective, name)
}

func NewNameGen(seed int64) Generator {
	gen := &NameGen{
		random: rand.New(rand.NewSource(seed)),
	}

	return gen
}
