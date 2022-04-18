package writer_reader

import (
	"sync"
)

func Test() {
	wg := sync.WaitGroup{}
	var writer Writer
	var reader Reader
	channel := make(chan int, 3)
	reader.NewReader(1, channel)
	writer.NewWriter(1, channel)
	wg.Add(2)
	go reader.Read(1)
	go writer.Write(1)
	go reader.Read(2)
	go reader.Read(3)
	wg.Wait()
}
