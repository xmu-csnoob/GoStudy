package producer_consumer

import "math/rand"

type Producer struct {
	id      int
	channel chan Goods
}

var names = []string{"Apple", "Orange", "Banana", "Toy", "Car", "Ring"}

func (producer *Producer) NewProducer(id int, channel chan Goods) {
	producer.id = id
	producer.channel = channel
}

func (producer *Producer) Produce(serialId int) {
	for {
		key := rand.Int() % len(names)
		id := rand.Int() % 100
		var goods Goods
		goods.NewGoods(id, names[key])
		producer.channel <- goods
		println("Producer ", producer.id, " produces goods named ", goods.name, goods.id, ",serial_id:", serialId)
	}
}
