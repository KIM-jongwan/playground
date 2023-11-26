package main

/*
Linked List 구현
*/

type List interface {
	//get(idx int)
	get(idx int) (data Node)
	//insert_back(data object)
	insert_back(data Node)
	//insert_front(data object)
	insert_front(data Node)
	//insert_at(data object, idx int)
	insert_at(data Node, idx int)
	//remove_back() (data object)
	remove_back() (data Node)
	//remove_front() (data object)
	remove_front() (data Node)
	//remove_at(idx int) (data object)
	remove_at(idx int) (data Node)
}

type Node struct {
 data string
	next *Node
}
type LinkedList struct {

}

func main(){

}