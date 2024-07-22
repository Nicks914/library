package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Book represents a book in the library
type Book struct {
	Name            string `json:"name"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publication_year"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	userType := GetUserType(r)

	var books []Book
	var booksUser []Book

	if userType == "admin" {
		books = ReadBooks("./adminUser.csv")
		booksUser = ReadBooks("./regularUser.csv")
		for _, singleBook := range booksUser {
			books = append(books, singleBook)
		}
	} else {
		books = ReadBooks("./regularUser.csv")
	}

	json.NewEncoder(w).Encode(books)
}

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	userType := GetUserType(r)

	if userType != "admin" {
		http.Error(w, "Only admin users can access this endpoint", http.StatusUnauthorized)
		return
	}

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if book.Name == "" || book.Author == "" || book.PublicationYear == 0 {
		http.Error(w, "Invalid book details", http.StatusBadRequest)
		return
	}

	// Validate publication year
	if book.PublicationYear < 0 || book.PublicationYear > time.Now().Year() {
		http.Error(w, "Invalid publication year", http.StatusBadRequest)
		return
	}

	// Add book to regularUser.csv
	err = AddBookToFile("regularUser.csv", book)
	if err != nil {
		http.Error(w, "Error adding book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	userType := GetUserType(r)

	if userType != "admin" {
		http.Error(w, "Only admin users can access this endpoint", http.StatusUnauthorized)
		return
	}

	var bookName string
	err := json.NewDecoder(r.Body).Decode(&bookName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if bookName == "" {
		http.Error(w, "Invalid book name", http.StatusBadRequest)
		return
	}

	// Delete book from regularUser.csv
	err = DeleteBookFromFile("regularUser.csv", bookName)
	if err != nil {
		http.Error(w, "Error deleting book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ReadBooks(filename string) []Book {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	var books []Book
	for _, record := range records {
		year, _ := strconv.Atoi(record[2])
		book := Book{Name: record[0], Author: record[1], PublicationYear: year}
		books = append(books, book)
	}

	return books
}

func AddBookToFile(filename string, book Book) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{book.Name, book.Author, strconv.Itoa(book.PublicationYear)})
	if err != nil {
		return err
	}

	// Append a newline character to signify the end of this book's entry
	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}

	return nil
}

func DeleteBookFromFile(filename string, bookName string) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var updatedRecords [][]string
	for _, record := range records {
		if strings.EqualFold(record[0], bookName) {
			continue
		}
		updatedRecords = append(updatedRecords, record)
	}

	file.Truncate(0)
	file.Seek(0, 0)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range updatedRecords {
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}
