package main

/*

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

>>>>>>> b9aa21474e4ea90b32a2f82bf1acaa14ee1aebfe
