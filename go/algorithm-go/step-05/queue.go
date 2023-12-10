package main

import "fmt"

/*
	Linked List로
	Queue 구현
 */

type Node struct {
	Front *Node
	Rear *Node
	Data int
}

type Queue struct {
	Head *Node
	Tail *Node
}

type QueueInterface interface{
	initQueue() (Queue)
	append(data int)
	pop() int
}

func initQueue() (Queue){
	queue := Queue{}
	return queue
}

func (queue *Queue) append(data int){
	newNode := Node{}
	newNode.Data = data

	if(queue.Head == nil){
		queue.Head = &newNode
		queue.Tail = &newNode
	}else{
		newNode.Rear = queue.Head
		queue.Head.Front = &newNode
		queue.Head = &newNode
	}	
}

func (queue *Queue) pop() (int){

	head := queue.Head

	queue.Head = head.Rear
	head.Front = head

	return head.Data;
}


func main() {

	queue := initQueue()
	(&queue).append(1)
	(&queue).append(2)
	(&queue).append(3)

	fmt.Println((&queue).pop())
	fmt.Println((&queue).pop())
	fmt.Println((&queue).pop())

	queue.append(5)
	queue.append(1)
	queue.append(2)

	fmt.Println(queue.pop())
	fmt.Println(queue.pop())
	fmt.Println(queue.pop())

	


}
