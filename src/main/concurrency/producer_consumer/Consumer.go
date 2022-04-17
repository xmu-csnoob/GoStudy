package producer_consumer

import "time"

type Consumer struct {
	id      int
	channel chan int
}

func (consumer *Consumer) NewConsumer(id int, channel chan int) {
	consumer.id = id
	consumer.channel = channel
}
func (consumer *Consumer) Consume(serialId int) {
	for {
		value := <-consumer.channel
		println("consumer ", consumer.id, " consumes", ":", value, " serial_id:", serialId)
	}
}
func (consumer *Consumer) ConsumeAtIntervals(serialId int) {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		value := <-consumer.channel
		println("consumer ", consumer.id, " consumes", ":", value, " serial_id:", serialId)
	}
}
