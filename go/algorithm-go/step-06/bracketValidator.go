package main

import (
	"fmt"
)

/*
	'[(){}]' 와 같은 문자열이 주어졌을 때, 괄호가 유효한지 판단

	1. 단순화
	-> input으로 입력되는 문자열은 소괄호만을 포함하고있음
		-> bracketValidator Stack(array) 초기화
		-> 문자열 탐색 log(n)
			-> 열린 괄화일 경우: stack app 괄호 닫기
			-> 닫힌 괄호일 경우: stack pop()
		-> 문자열 탐색 종료 후 stack 이 비어있을 경우 true / 비어있지 않을 경우 false
*/

func bracketValidator(input string) bool {

	if len(input) == 0 {
		fmt.Println("empty input")
		return false
	}

	var charArray []byte = []byte(input)
	for i := 0; i < len(input); i++ {
		fmt.Println(charArray[i])
	}

	return true
}

func main() {

	param1 := "param1"

	bracketValidator(param1)

}
