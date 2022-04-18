package writer_reader

import "time"

type Reader struct {
	id      int
	channel chan int
}

func (reader *Reader) NewReader(id int, channel chan int) {
	reader.id = id
	reader.channel = channel
}
func (reader *Reader) Read(serialId int) {
	for {
		value := <-reader.channel
		println("reader ", reader.id, " consumes", ":", value, " serial_id:", serialId)
	}
}
func (reader *Reader) ReadAtIntervals(serialId int) {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		value := <-reader.channel
		println("reader ", reader.id, " consumes", ":", value, " serial_id:", serialId)
	}
}
