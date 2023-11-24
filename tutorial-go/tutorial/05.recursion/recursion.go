package main

import("fmt")


func Recursion(a int){

	if a == 0 {
		fmt.Println("탈출 조건");
		return
	}

	fmt.Println("a != 0")
	Recursion(a - 1)
	fmt.Printf("After %d \n",a);

}

func main(){

	Recursion(3);

}