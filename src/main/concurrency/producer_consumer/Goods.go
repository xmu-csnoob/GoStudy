package producer_consumer

type Goods struct {
	id   int
	name string
}

func (goods *Goods) NewGoods(id int, name string) {
	goods.id = id
	goods.name = name
}
