package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	giver, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	receiver, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	marble, _ := strconv.Atoi(scanner.Text())

	giveMarble(&giver, &receiver, marble)
	fmt.Printf("%d \n", giver)
	fmt.Printf("%d", receiver)
}

//Tulis kode disini

func giveMarble(giver *int, receiver *int, marble int) (int, int) {
	if marble < *giver {
		*giver -= marble
		*receiver += marble
	} else {
		fmt.Println("Kelereng tidak cukup untuk dibagikan")
	}
	return *giver, *receiver
}
