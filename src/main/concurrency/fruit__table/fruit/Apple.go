package fruit

import (
	"math/rand"
)

type Apple struct {
	Name   string
	Weight int
}

func (apple *Apple) Produce(channel chan *Apple) {
	apple.Weight = rand.Int() % 100
	channel <- apple
}
