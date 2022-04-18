package writer_reader

import (
	"math/rand"
	"time"
)

type Writer struct {
	id      int
	channel chan int
}

func (writer *Writer) NewWriter(id int, channel chan int) {
	writer.id = id
	writer.channel = channel
}
func (writer Writer) Write(serialId int) {
	for {
		value := rand.Int() % 10
		writer.channel <- value
		println("writer ", writer.id, " produces", ":", value, " serial_id:", serialId)
	}
}
func (writer Writer) WriteInIntervals(serialId int) {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		value := rand.Int() % 10
		writer.channel <- value
		println("writer ", writer.id, " produces", ":", value, " serial_id:", serialId)
	}
}
