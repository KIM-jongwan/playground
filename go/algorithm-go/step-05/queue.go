package main

/*
	Linked List로
	Queue 구현
 */

type Node struct {
	Front *Node
	Rear *Node
	Data int
}

type Queue interface{
	initQueue()
	append(data int)
	pop() int
}

func initQueue(){

}

func (node *Node) append(data int){

}

func pop() (int){

	return 0;
}


func main() {

}
