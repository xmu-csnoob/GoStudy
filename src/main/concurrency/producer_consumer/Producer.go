package producer_consumer

import (
	"math/rand"
	"time"
)

type Producer struct {
	id      int
	channel chan int
}

func (producer *Producer) NewProducer(id int, channel chan int) {
	producer.id = id
	producer.channel = channel
}
func (producer Producer) Produce(serialId int) {
	for {
		value := rand.Int() % 10
		producer.channel <- value
		println("producer ", producer.id, " produces", ":", value, " serial_id:", serialId)
	}
}
func (producer Producer) ProduceInIntervals(serialId int) {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		value := rand.Int() % 10
		producer.channel <- value
		println("producer ", producer.id, " produces", ":", value, " serial_id:", serialId)
	}
}
