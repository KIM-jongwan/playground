package main

import (
	"fmt"
)

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false;
	}
	return a/b, true;
}

func main(){

	c, success1 := Divide(1,2);
	d, success2 := Divide(3,0);

	fmt.Println(c, success1);
	fmt.Println(d, success2);

}
