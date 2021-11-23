// **** file: main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type Book struct {
	isbn   string
	title  string
	author string
	price  float32
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:changeme@localhost:5432/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", bookShow)
	http.HandleFunc("/books/create", bookCreate)
	http.HandleFunc("/books/delete", bookDelete)
	http.HandleFunc("/books/union", bookUnion)

	http.ListenAndServe(":3000", nil)
	fmt.Sprintln("Server launch on http://localhost:3000")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, %.2f€\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}

func bookShow(w http.ResponseWriter, r *http.Request) {

	if "GET" != r.Method {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	isbn := r.FormValue("isbn")
	if "" == isbn {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	rows := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	book := new(Book)
	err := rows.Scan(&book.isbn, &book.title, &book.author, &book.price)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if nil != err {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, %.2f€\n", book.isbn, book.title, book.author, book.price)
}

func bookCreate(w http.ResponseWriter, r *http.Request) {

	if "POST" != r.Method {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	isbn := r.FormValue("isbn")
	title := r.FormValue("title")
	author := r.FormValue("author")
	if isbn == "" || title == "" || author == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	price, err := strconv.ParseFloat(r.FormValue("price"), 32)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	result, err := db.Exec("INSERT INTO books VALUES($1, $2, $3, $4)", isbn, title, author, price)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Book %s created successfully (%d row affected)\n", isbn, rowsAffected)
}

func bookDelete(w http.ResponseWriter, r *http.Request) {
	if "DELETE" != r.Method {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	isbn := r.FormValue("isbn")
	if "" == isbn {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	rows := db.QueryRow("DELETE FROM books WHERE isbn = $1 RETURNING *", isbn)

	book := new(Book)
	err := rows.Scan(&book.isbn, &book.title, &book.author, &book.price)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if nil != err {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "%s, %s, %s, %.2f€\n", book.isbn, book.title, book.author, book.price)
}

func bookUnion(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	rows, err := db.Query("SELECT * FROM books UNION ALL SELECT * FROM books2 ORDER BY price")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	bks := make([]*Book, 0)
	for rows.Next() {
		bk := new(Book)
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, %.2f€\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}
