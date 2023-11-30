package main

import (
	"container/list"
	"fmt"
)

type FrontMiddleBackQueue struct {
	left  *list.List
	right *list.List
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		left:  list.New(),
		right: list.New(),
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.left.PushFront(val)
	if this.left.Len() == this.right.Len()+2 {
		this.right.PushFront(this.left.Back().Value.(int))
		this.left.Remove(this.left.Back())
	}
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	if this.left.Len() > this.right.Len() {
		this.right.PushFront(this.left.Back().Value.(int))
		this.left.Remove(this.left.Back())
	}
	this.left.PushBack(val)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.right.PushBack(val)
	if this.left.Len()+2 == this.right.Len() {
		this.left.PushBack(this.right.Front().Value.(int))
		this.right.Remove(this.right.Front())
	}
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.left.Len() == 0 && this.right.Len() == 0 {
		return -1
	}
	var val int
	if this.left.Len() != 0 {
		val = this.left.Front().Value.(int)
		this.left.Remove(this.left.Front())
		if this.left.Len()+2 == this.right.Len() {
			this.left.PushBack(this.right.Front().Value.(int))
			this.right.Remove(this.right.Front())
		}
	} else {
		val = this.right.Front().Value.(int)
		this.right.Remove(this.right.Front())
	}
	return val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.left.Len() == 0 && this.right.Len() == 0 {
		return -1
	}
	var val int
	if this.left.Len() >= this.right.Len() {
		val = this.left.Back().Value.(int)
		this.left.Remove(this.left.Back())
	} else {
		val = this.right.Front().Value.(int)
		this.right.Remove(this.right.Front())
	}
	return val
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.left.Len() == 0 && this.right.Len() == 0 {
		return -1
	}
	var val int
	if this.right.Len() != 0 {
		val = this.right.Back().Value.(int)
		this.right.Remove(this.right.Back())
		if this.left.Len() == this.right.Len()+2 {
			this.right.PushFront(this.left.Back().Value.(int))
			this.left.Remove(this.left.Back())
		}
	} else {
		val = this.left.Back().Value.(int)
		this.left.Remove(this.left.Back())
	}
	return val
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();

["FrontMiddleBackQueue","pushMiddle","popMiddle","pushMiddle","popMiddle","pushFront","pushMiddle","pushMiddle","popMiddle","popMiddle","pushBack","popMiddle"]
[[],[493299],[],[75427],[],[753523],[677444],[431158],[],[],[47949],[]]

*/

func main() {
	obj := Constructor()
	obj.PushMiddle(493299)
	obj.PopMiddle()
	obj.PushMiddle(75427)
	obj.PopMiddle()
	obj.PushFront(753523)
	obj.PushMiddle(677444)
	obj.PushMiddle(431158)
	c := obj.PopMiddle()
	a := obj.PopMiddle()
	obj.PushBack(47949)
	b := obj.PopMiddle()
	fmt.Println(c, a, b)
}
