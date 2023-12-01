package main

import "fmt"

/*
 인터넷 방문 기록과 동일한 동작을 하는 BrowserHistory class를 구현
구현할 브라우저는 homepage에서 시작하고, url을 입력받아 이동 할 수 있다.
또 이동 히스토리를 기록하여 "뒤로 가기", "앞으로 가기" 기능을 사용할 수 있다.
*/

type Node struct {
	Head *Node
	Tail *Node
	Data string
}

/*
BrowserHistory(homepage string) 호출하면 브라우저는 homepage에서 시작
void visit(url string)을 호출하면 현재 page 앞에 있는 페이지 기록은 삭제되고 url로 방문
back(int steps) (string) 을 호출하면 steps수 만큼 뒤로가기한다. ....
forward(int steps) (string) 을 호출하면 steps수 만큼 앞으로 가낟. ...
*/
type BrowserHistoryInterface interface {
	BrowseHistory(url string) Node
	new(url string) Node
	visit(url string) Node
	forward(steps int) *Node
	back(steps int) *Node
}

func BrowseHistory(url string) Node {
	var node Node
	node.Head = &node
	node.Tail = &node
	node.Data = url

	return node
}

func (node *Node) new(url string) Node {

}

func (node *Node) visit(url string) Node {

}

func (node *Node) forward(steps int) Node {

}

func (node *Node) back(steps int) Node {

}

func main() {

	node := BrowseHistory("https://www.google.com")

	fmt.Println(node.Head)
	fmt.Println(node.Tail)
	fmt.Println(node.Data)
}
