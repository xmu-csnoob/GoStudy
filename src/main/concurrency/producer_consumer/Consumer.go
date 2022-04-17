package producer_consumer

type Consumer struct {
	id      int
	channel chan int
}

func (consumer *Consumer) NewConsumer(id int, channel chan int) {
	consumer.id = id
	consumer.channel = channel
}
func (consumer *Consumer) Consume(serial_id int) {
	for {
		value := <-consumer.channel
		println("consumer ", consumer.id, " consumes", ":", value, " serial_id:", serial_id)
	}
}
