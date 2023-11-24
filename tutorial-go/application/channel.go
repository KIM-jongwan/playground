package main
 
func main() {
  // 정수형 채널을 생성한다 
  ch := make(chan int)
 
  go func() {
    ch <- 123   //채널에 123을 보낸다
				ch <- 456
				ch <- 789
  }()
 

		elem := <- ch;
		println(elem);
		elem = <- ch;
		println(elem);
		elem = <- ch;
		println(elem);
}