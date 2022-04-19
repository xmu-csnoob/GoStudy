package table

import (
	"awesomeProject/src/main/concurrency/fruit__table/fruit"
	"math/rand"
)

type Table struct {
	AppleChan  chan *fruit.Apple
	OrangeChan chan *fruit.Orange
	EmptyChan  chan bool
}

func (table *Table) NewTable(appleChan chan *fruit.Apple, orangeChan chan *fruit.Orange, emptyChannel chan bool) {
	table.AppleChan = appleChan
	table.OrangeChan = orangeChan
	table.EmptyChan = emptyChannel
}
func (table *Table) Produce(serialId int) {
	for {
		flag := rand.Int() % 2
		_ = <-table.EmptyChan
		if flag == 0 {
			var apple fruit.Apple
			println("Produce An Apple")
			apple.Produce(table.AppleChan)
		} else {
			var orange fruit.Orange
			println("Produce An Orange")
			orange.Produce(table.OrangeChan)
		}
	}
}
