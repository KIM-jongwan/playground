package main

import "fmt"

/*
Linked List 구현
*/

type List interface {
	get(idx int) (data *Node)
	insert(data Node)
	insert_at(data Node, idx int)
	remove() (data Node)
	remove_at(idx int) (data Node)
}

type Node struct {
	Data string //데이터
	Next *Node  //다음 노드 주소
}
type LinkedList struct {
	Head  *Node //List Head Node 주소
	Tail  *Node //List Tail Node 주소
	Count int   //List Count
}

func (list *LinkedList) get(idx int) *Node {

	if list.Count == 0 {
		fmt.Println("empty list")
		return nil
	}

	if list.Count < idx {
		fmt.Println("Index Out of Bound")
		return nil
	}

	current := list.Head
	for i := 0; i < idx; i++ {
		current = current.Next
	}
	fmt.Println(current)

	return current
}

func main() {

	var node1 Node
	var list LinkedList

	node1.Data = "this is node 1"
	node1.Next = nil

	list.Head = &node1

	list.get(1)

}
