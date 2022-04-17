package producer_consumer

import "sync"

func Test() {
	wg := sync.WaitGroup{}
	var consumer Consumer
	var producer Producer
	channel := make(chan int, 3)
	consumer.NewConsumer(1, channel)
	producer.NewProducer(1, channel)
	wg.Add(2)
	go consumer.Consume(1)
	go producer.Produce(1)
	go consumer.Consume(2)
	go consumer.Consume(3)
	wg.Wait()
}
