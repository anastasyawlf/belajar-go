//challenge 1

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// // Tulis kode struct disini
// type Penduduk struct {
// 	nama   string
// 	umur   int
// 	alamat string
// }

// func main() {
// 	var penduduk1 Penduduk
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	penduduk1.nama = scanner.Text()
// 	scanner.Scan()
// 	penduduk1.umur, _ = strconv.Atoi(scanner.Text())
// 	scanner.Scan()
// 	penduduk1.alamat = scanner.Text()

// 	//Tulis kode format disini
// 	fmt.Printf("Hello, my name is %s. Im %d years old. I live in %s", penduduk1.nama, penduduk1.umur, penduduk1.alamat)
// }

// challenge2
package main

import "fmt"

type Counter struct {
	// Isi struct ini
	Value int
}

func (c *Counter) Increment() {
	// Tulis kode disini
	c.Value++
}

func main() {
	// Tulis kode disini
	counter := &Counter{3}
	counter.Increment()
	fmt.Println(counter.Value)
}
