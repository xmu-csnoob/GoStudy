package array

import (
	"errors"
	"fmt"
)

type Array struct {
	buffer []int
	len    int
}

func (array *Array) CreateArray(len int) {
	array.len = len
	array.buffer = make([]int, len)
}
func (array *Array) Assign(index int, value int) error {
	if index >= array.len || index < 0 {
		return errors.New("数组越界")
	}
	array.buffer[index] = value
	return nil
}
func (array *Array) CopyFromArray(arr []int) error {
	if array.len > len(arr) {
		return errors.New("数组长度大于所给参数数组")
	}
	if array.len < len(arr) {
		err := array.Resize(len(arr))
		if err != nil {
			return err
		}
	}
	for key, value := range arr {
		array.buffer[key] = value
	}
	return nil
}
func (array *Array) Resize(size int) error {
	if size == array.len {
		return errors.New("数组容量未改变")
	} else if size <= array.len {
		return errors.New("不允许缩小数组容量")
	}
	newArr := make([]int, size)
	for key, value := range array.buffer {
		newArr[key] = value
	}
	array.buffer = newArr
	return nil
}
func (array Array) PrintAll() {
	fmt.Println(array.buffer)
}
