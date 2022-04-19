package fruit__table

import (
	"awesomeProject/src/main/concurrency/fruit__table/fruit"
	"awesomeProject/src/main/concurrency/fruit__table/person"
	"awesomeProject/src/main/concurrency/fruit__table/table"
	"sync"
)

func Test() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	var t table.Table
	appleChannel := make(chan *fruit.Apple, 3)
	orangeChannel := make(chan *fruit.Orange, 3)
	boolChannel := make(chan bool, 1)
	t.NewTable(appleChannel, orangeChannel, boolChannel)
	t.EmptyChan <- true
	var son person.Son
	var daughter person.Daughter
	go t.Produce(1)
	go son.Eat(&t)
	go daughter.Eat(&t)
	wg.Wait()
}
