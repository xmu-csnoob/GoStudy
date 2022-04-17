package producer_consumer

import "math/rand"

type Producer struct {
	id      int
	channel chan int
}

func (producer *Producer) NewProducer(id int, channel chan int) {
	producer.id = id
	producer.channel = channel
}
func (producer Producer) Produce(serial_id int) {
	for {
		value := rand.Int() % 10
		producer.channel <- value
		println("producer ", producer.id, " produces", ":", value, " serial_id:", serial_id)
	}
}
