package person

import (
	"awesomeProject/src/main/concurrency/fruit_table_unlimited/table"
)

type Son struct {
}

func (son *Son) Eat(table *table.Table) {
	for {
		f := <-table.AppleChan
		println("Son Eats An Apple Weights ", f.Weight)
	}
}
