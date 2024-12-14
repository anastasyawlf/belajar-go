package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer loop()
	fmt.Println("Done")
}

func loop() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done Loop")
}
