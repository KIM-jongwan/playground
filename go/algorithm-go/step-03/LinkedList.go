package main

/*
Linked List 구현
*/

type List interface {
 get(idx int) (data Node)
 insert_back(data Node)
 insert_front(data Node)
 insert_at(data Node, idx int)
 remove_back() (data Node)
 remove_front() (data Node)
 remove_at(idx int) (data Node)
}

type Node struct {
 data interface{}
 next *Node
}
type LinkedList struct {
 head Node
	tail Node
}

func main(){

}