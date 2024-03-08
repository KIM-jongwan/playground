package main

import "fmt"

/*
	go slice
	사용하여 queue 구현
*/

type Queue []interface {
}

func (q *Queue) Enqueue(data interface{}) {
	*q = append(*q, data)
}

func (q *Queue) Dequeue() interface{} {

	if len(*q) == 0 {
		fmt.Println("empty Queue")
		return nil
	}

	data := (*q)[0]
	*q = (*q)[1:]

	return data

}

func (q *Queue) printElements() {
	if len(*q) == 0 {
		fmt.Println("empty queue")
		return
	}

	for i := 0; i < len(*q); i++ {
		fmt.Println((*q)[i])
	}

}

func main() {

	queue := Queue{}

	fmt.Println(queue.Dequeue()) //empty

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.printElements()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.printElements()
}
