package main
import(
	"fmt"
)


func Divide(a, b int) (result int, success bool){
	if b == 0 {
		result = 1;
		success = true;
		return;
	}
	return;
}

func main(){

	c, boolResult1 := Divide(1,2);
	fmt.Println(c, boolResult1);
	
}