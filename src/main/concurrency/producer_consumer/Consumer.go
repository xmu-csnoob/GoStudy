package producer_consumer

type Consumer struct {
	id      int
	channel chan Goods
}

func (consumer *Consumer) NewConsumer(id int, channel chan Goods) {
	consumer.id = id
	consumer.channel = channel
}

func (consumer *Consumer) Consume(serialId int) {
	for {
		goods := <-consumer.channel
		println("Consumer ", consumer.id, " consumes goods named ", goods.name, goods.id, ",serial_id:", serialId)
	}
}
