package table

import (
	"awesomeProject/src/main/concurrency/fruit__table/fruit"
	"math/rand"
	"time"
)

type Table struct {
	AppleChan  chan *fruit.Apple
	OrangeChan chan *fruit.Orange
	IsEmpty    chan bool
}

func (table *Table) NewTable(appleChan chan *fruit.Apple, orangeChan chan *fruit.Orange, boolChannel chan bool) {
	table.AppleChan = appleChan
	table.OrangeChan = orangeChan
	table.IsEmpty = boolChannel
}
func (table *Table) Produce() {
	timeTicker := time.NewTicker(1 * time.Second)
	for _ = range timeTicker.C {
		isEmpty := <-table.IsEmpty
		if isEmpty {
			flag := rand.Int() % 2
			if flag == 0 {
				var apple fruit.Apple
				apple.Produce(table.AppleChan)
				table.IsEmpty <- false
			} else {
				var orange fruit.Orange
				orange.Produce(table.OrangeChan)
				table.IsEmpty <- false
			}
		}
	}
}
