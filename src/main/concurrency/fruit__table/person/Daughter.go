package person

import (
	"awesomeProject/src/main/concurrency/fruit__table/table"
)

type Daughter struct {
}

func (daughter *Daughter) Eat(table *table.Table) {
	for {
		f := <-table.OrangeChan
		println("Daughter Eats An Orange Weights ", f.Weight)
		table.EmptyChan <- true
	}
}
