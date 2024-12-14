package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("Deferred function")
	}()

	fmt.Println("Before Panic")
	panic("panic occured")

	// var input string
	// fmt.Print("Masukkan nama: ")
	// fmt.Scanln(&input)

	// fmt.Println("Start")
	// if !isEmpty(input) {
	// 	fmt.Println(input)
	// }
	// fmt.Println("Done")
}

// func isEmpty(input string) (empty bool) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Terjadi Panic:", r)
// 		}
// 	}()
// 	if input == "" {
// 		panic("Nama tidak boleh kosong")
// 	}
// 	empty = true
// 	return
// }
