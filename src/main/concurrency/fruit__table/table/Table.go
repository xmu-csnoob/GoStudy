package table

import (
	"awesomeProject/src/main/concurrency/fruit__table/fruit"
	"math/rand"
	"time"
)

type Table struct {
	AppleChan  chan *fruit.Apple
	OrangeChan chan *fruit.Orange
}

func (table *Table) NewTable(appleChan chan *fruit.Apple, orangeChan chan *fruit.Orange) {
	table.AppleChan = appleChan
	table.OrangeChan = orangeChan
}
func (table *Table) Produce() {
	timeTicker := time.NewTicker(5 * time.Second)
	for _ = range timeTicker.C {
		flag := rand.Int() % 2
		if flag == 0 {
			var apple fruit.Apple
			apple.Produce(table.AppleChan)
		} else {
			var orange fruit.Orange
			orange.Produce(table.OrangeChan)
		}
	}
}
