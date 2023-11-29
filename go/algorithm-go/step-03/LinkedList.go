package main

import "fmt"

/*
Linked List 구현
*/

type List interface {
	get(idx int) (data *Node)
	insert(data Node)
	insert_at(node Node, idx int)
	remove() (node Node)
	remove_at(idx int) (data Node)
}

type Node struct {
	Data string //데이터
	Next *Node  //다음 노드 주소
}
type LinkedList struct {
	Node  *Node //List Node 주소
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

	current := list.Node
	for i := 0; i < idx; i++ {
		current = current.Next
	}
	fmt.Println(current)

	return current
}

/*
 func (list *LinkedList)insert(node *Node)
  - LinkednList의 Tail 뒷부분으로 새로운 Node 추가

 00. List(L-value, 직접 값 수정) Node field 초기화
	01. Node nil check
	  -> (true) Node := parameter Node
			   Count += Count
			-> (false) for current.Next == current (Tail node)
			   current.Next := paramter Node
						Count += Count
*/
func (list *LinkedList) insert(node *Node) {
	current := list.Node

	if current == nil {
		list.Node = node
		list.Count = 1
	} else {
		for current.Next == current {
			current = current.Next
		}
		current.Next = node
		list.Count += list.Count
	}
}

/*
 func (list *LinkedList)insert_at(node *Node, idx int)
  - LinkednList의 idx parameter에 해당하는 순번에 새로운 Node 추가

 00. List(L-value, 직접 값 수정) Node field 초기화
	01. List Count > idx 검증
	02. newNode.Next = idxNode.Next
	03. idxNode.Next = newNode
*/
func (list *LinkedList) insert_at(newNode *Node, idx int) {

	idxNode := list.get(idx)

	newNode.Next = idxNode.Next
	idxNode.Next = newNode

	list.Count += list.Count
}

func (list *LinkedList) remove() (node Node) {
	return
}

func (list *LinkedList) remove_at(idx int) (node Node) {
	return
}

func main() {

	var list LinkedList
	var node1 Node
	var node2 Node

	node1.Data = "this is node 1"
	node1.Next = &node1

	node2.Data = "this is node 2"
	node2.Next = &node2

	list.insert(&node1)
	list.insert(&node2)

	fmt.Println(list.get(0))
	fmt.Println(list.get(1))

}
