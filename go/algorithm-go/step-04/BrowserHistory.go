package main

/*
 인터넷 방문 기록과 동일한 동작을 하는 BrowserHistory class를 구현
구현할 브라우저는 homepage에서 시작하고, url을 입력받아 이동 할 수 있다.
또 이동 히스토리를 기록하여 "뒤로 가기", "앞으로 가기" 기능을 사용할 수 있다.
*/

type Node struct {
	Next *Node
	Data string
}

type BrowserHistoryInfo struct {
	Head        *Node
	CurrentNode *Node
}

type BrowserHistory struct {
	Head *Node
	Tail *Node
}

/*
BrowserHistory(homepage string) 호출하면 브라우저는 homepage에서 시작
void visit(url string)을 호출하면 현재 page 앞에 있는 페이지 기록은 삭제되고 url로 방문
back(int steps) (string) 을 호출하면 steps수 만큼 뒤로가기한다. ....
forward(int steps) (string) 을 호출하면 steps수 만큼 앞으로 간다. ...
*/
type BrowserHistoryInterface interface {
	BrowseHistory(url string) BrowserHistoryInfo
	visit(url string) Node
	forward(steps int) *Node
	back(steps int) *Node
}

func BrowserHistory(url string) BrowserHistoryInfo {
	newNode := Node{}
	newNode.Next = &newNode
	newNode.Data = url

	browserHistory := BrowserHistoryInfo{}
	browserHistory.Head = &newNode
	browserHistory.CurrentNode = &newNode

	return browserHistory
}

func (browserHistory *BrowserHistoryInfo) visit(url string) *Node {
	newNode := Node{}
	newNode.Next = &newNode
	newNode.Data = url

	browserHistory.CurrentNode.Next = &newNode

	return &newNode
}

func (browserHistory *BrowserHistoryInfo) forward(steps int) Node {

	currnet := browserHistory.CurrentNode

	for steps > 0 {
		currnet = currnet.Next
		steps--
	}

	return *currnet
}

/*
func (node *Node) back(steps int) Node {

}
*/
func main() {

}
