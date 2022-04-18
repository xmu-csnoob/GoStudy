package producer_consumer

import "sync"

func Test() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var consumer Consumer
	channel := make(chan Goods)
	consumer.NewConsumer(1, channel)
	var producer Producer
	producer.NewProducer(1, channel)
	go producer.Produce(1)
	go consumer.Consume(1)
	wg.Wait()
}
