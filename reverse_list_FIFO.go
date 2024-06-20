package main

import (
	"container/list"
	"fmt"
)

func main() {
	mylist := list.New()
	Push(1, mylist)
	Push(2, mylist)
	Push(3, mylist)
	//Pop(mylist)
	//fmt.Println(Pop(mylist))
	printQueue(mylist)
	printQueue(ReverseList(mylist))

}

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {

	if queue.Front() == nil {
		return nil
	} else {
		elem := queue.Front()
		queue.Remove(queue.Front())
		return elem
	}
}

func printQueue(queue *list.List) {
	for i := queue.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value)
	}
}
func ReverseList(l *list.List) *list.List {
	reversedList := list.New()
	for i := l.Front(); i != nil; i = i.Next() {
		reversedList.PushFront(i.Value)
	}
	return reversedList
}
