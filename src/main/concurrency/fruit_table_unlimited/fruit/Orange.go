package fruit

import (
	"math/rand"
)

type Orange struct {
	Name   string
	Weight int
}

func (orange *Orange) Produce(channel chan *Orange) {
	orange.Weight = rand.Int() % 100
	channel <- orange
}
