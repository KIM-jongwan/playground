package main

import "fmt"

/*
	go slice
	사용하여 queue 구현
*/

type Queue []interface{
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(data interface{}){
	*q = append(*q, data);
	fmt.Println("append queue")
}



func main()  {

	var a int = 3;

	q := Queue{}
	pq := &q

	q.Enqueue(1)

	fmt.Println(&a)
	fmt.Println(q)
	fmt.Println(&q)
	fmt.Println(pq)
	fmt.Println(*pq)
}