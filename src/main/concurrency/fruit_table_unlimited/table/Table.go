package table

import (
	"awesomeProject/src/main/concurrency/fruit_table_unlimited/fruit"
	"math/rand"
)

type Table struct {
	AppleChan  chan *fruit.Apple
	OrangeChan chan *fruit.Orange
}

func (table *Table) NewTable(appleChan chan *fruit.Apple, orangeChan chan *fruit.Orange) {
	table.AppleChan = appleChan
	table.OrangeChan = orangeChan
}
func (table *Table) Produce(serialId int) {
	for {
		flag := rand.Int() % 2
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
