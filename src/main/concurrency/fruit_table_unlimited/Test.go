package fruit_table_unlimited

import (
	"awesomeProject/src/main/concurrency/fruit_table_unlimited/fruit"
	"awesomeProject/src/main/concurrency/fruit_table_unlimited/person"
	table2 "awesomeProject/src/main/concurrency/fruit_table_unlimited/table"
	"sync"
)

func Test() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	var t table2.Table
	appleChannel := make(chan *fruit.Apple, 3)
	orangeChannel := make(chan *fruit.Orange, 3)
	t.NewTable(appleChannel, orangeChannel)
	var son person.Son
	var daughter person.Daughter
	go t.Produce(1)
	go son.Eat(&t)
	go daughter.Eat(&t)
	wg.Wait()
}
