package person

import (
	"awesomeProject/src/main/concurrency/fruit__table/table"
)

type Son struct {
}

func (son *Son) Eat(table *table.Table) {
	for {
		f := <-table.AppleChan
		println("Son Eats An Apple Weights ", f.Weight)
		table.EmptyChan <- true
	}
}
