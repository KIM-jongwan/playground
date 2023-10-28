package main

import (
	"log"
	"os"
)

func main(){
		f, err := os.Open("/Users/jongwan/Desktop/누가나게.jpeg")
		if err != nil {
			log.Fatal(err.Error())
		}
		println(f.Name())
}