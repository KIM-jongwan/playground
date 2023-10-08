package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message string
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to %v\n", m.message, m.phoneNumber)
	fmt.Printf("=========================================================")
}

func main(){

	test(messageToSend{
		phoneNumber: 123123123123,
		message: "Thanks for order",
	})

}