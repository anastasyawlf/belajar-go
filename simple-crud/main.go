package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Id          int
	Title       string
	Author      string
	ReleaseYear string
	Pages       int
}

var books []Book

var fileName string = "data.csv"

func main() {
	// createFile(fileName)
	// addNewBook()
	loadDataFromCSV(fileName)
	viewAllBooks()
}

// func createFile(fileName string) {
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	} else {
// 		fmt.Println("File", fileName, "berhasil dibuat.")
// 	}
// 	defer file.Close()
// }

func addNewBook() error {
	var newBook Book

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Book Details:")
	fmt.Print("Book ID:")
	scanner.Scan()
	newBook.Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Book Title:")
	scanner.Scan()
	newBook.Title = scanner.Text()

	fmt.Print("Book Author:")
	scanner.Scan()
	newBook.Author = scanner.Text()

	fmt.Print("Book Release Year:")
	scanner.Scan()
	newBook.ReleaseYear = scanner.Text()

	fmt.Print("Book Pages:")
	scanner.Scan()
	newBook.Pages, _ = strconv.Atoi(scanner.Text())

	_, err := findBookById(newBook.Id)
	if err != nil {
		books = append(books, newBook)
	} else {
		fmt.Errorf("Book with id: %d already exists", newBook.Id)
	}

	err = saveDataToCSV(fileName)
	if err != nil {
		return err
	}
	fmt.Println("Book added successfully")

	return nil
}

func findBookById(id int) (Book, error) {
	for _, book := range books {
		if book.Id == id {
			return book, nil
		}
	}
	return Book{}, fmt.Errorf("id: %d not found", id)
}

func saveDataToCSV(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	for _, book := range books {
		row := strconv.Itoa(book.Id) + "," + book.Title + "," + book.Author + "," + book.ReleaseYear + "," + strconv.Itoa(book.Pages) + "\n"

		file.WriteString(row)
	}
	return nil
}

func loadDataFromCSV(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		record := strings.Split(scanner.Text(), ",")
		// fmt.Println(record)
		id, _ := strconv.Atoi(record[0])
		pages, _ := strconv.Atoi(record[4])

		book := Book{
			Id:          id,
			Title:       record[1],
			Author:      record[2],
			ReleaseYear: record[3],
			Pages:       pages,
		}
		books = append(books, book)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error opening csv file: %w", err)
	}

	return nil
}

func viewAllBooks() error {
	if len(books) == 0 {
		return fmt.Errorf("No books available.")
	}

	for i, book := range books {
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("Book - ", i+1)
		fmt.Println("Book Id : ", book.Id)
		fmt.Println("Book Title : ", book.Title)
		fmt.Println("Book Author : ", book.Author)
		fmt.Println("Book Release Year : ", book.ReleaseYear)
		fmt.Println("Book Pages : ", book.Pages)
		fmt.Println(strings.Repeat("=", 50))
	}
	return nil
}
