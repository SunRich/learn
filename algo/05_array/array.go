package _5_array

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []int
	length uint
}

//初始化
func NewArray(cap uint) *Array {
	if cap <= 0 {
		return nil
	}
	return &Array{
		data:   make([]int, cap, cap),
		length: 0,
	}
}

func (this *Array) Length() uint {
	return this.length
}
func (this *Array) isIndexOutofRange(index uint) bool {
	if index >= uint(cap(this.data)) {
		return true
	}
	return false
}

//访问
func (this *Array) FindIndex(index uint) (value int, err error) {
	if this.isIndexOutofRange(index) {
		err = errors.New("out of range")
		return
	}
	value = this.data[index]
	return
}

//插入
//判断是否已满
//判断越界
func (this *Array) Insert(index uint, value int) (err error) {
	if index<0{
		err = errors.New("index error")
		return
	}
	if this.isIndexOutofRange(index) {
		err = errors.New("out of range")
		return
	}
	if this.length == uint(cap(this.data)) {
		err = errors.New("array is full")
		return
	}
	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = value
	this.length++
	return
}

//删除
func (this *Array)Delete(index uint)(err error)  {
	if index<0{
		err = errors.New("index error")
		return
	}
	if this.isIndexOutofRange(index) {
		err = errors.New("out of range")
		return
	}
	for i := index; i < this.length-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	return
}
//打印数列
func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.Length(); i++ {
		format += fmt.Sprintf("|%+v", this.data[i])
	}
	fmt.Println(format)
}
