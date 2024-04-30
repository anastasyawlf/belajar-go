package main

import "fmt"

func main() {

	// Nilai maksimum bernilai n
	fmt.Print("Masukkan nilai maksimum (n): ")
	var n int
	fmt.Scanf("%d", &n)

	// Proses nilai n berdasarkan fungsi fizzBuzz
	for i := 1; i <= n; i++ {
		fizzBuzz(i)
	}

}

func fizzBuzz(i int) {

	if i%3 == 0 && i%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if i%3 == 0 {
		fmt.Println("Fizz")
	} else if i%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(i)
	}

}
