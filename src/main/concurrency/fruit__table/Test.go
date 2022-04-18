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
	appleChannel := make(chan *fruit.Apple)
	orangeChannel := make(chan *fruit.Orange)
	t.NewTable(appleChannel, orangeChannel)
	go t.Produce()
	var son person.Son
	var daughter person.Daughter
	go son.Eat(t)
	go daughter.Eat(t)
	wg.Wait()
}
