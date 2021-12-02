package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vcycyv/bookshop/entity"
	"github.com/vcycyv/bookshop/representation"
)

func TestBook_Add(t *testing.T) {
	var bookName = "sqlite"
	book := addBook(bookName)
	assert.Equal(t, bookName, book.Name)

	_ = db.Migrator().DropTable(&entity.Book{})
	_ = db.Migrator().CreateTable(&entity.Book{})
}

func TestBook_Get(t *testing.T) {
	w := httptest.NewRecorder()

	var bookName = "sqlite"
	book := addBook(bookName)
	bookName = book.Name

	uri := fmt.Sprintf("/books/%s", book.ID)
	r := gin.Default()
	req := httptest.NewRequest(
		http.MethodGet,
		uri,
		nil,
	)
	r.GET("/books/:id", bookHdlr.Get)
	r.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	got := &representation.Book{}
	_ = json.Unmarshal(body, &got)
	assert.True(t, len(got.ID) > 0)
	assert.Equal(t, bookName, got.Name)

	_ = db.Migrator().DropTable(&entity.Book{})
	_ = db.Migrator().CreateTable(&entity.Book{})
}

func TestBook_GetAll(t *testing.T) {
	var bookName = "sqlite"
	book := addBook(bookName)

	books := getBooks()

	assert.Equal(t, len(books), 1)
	assert.Equal(t, book.Name, books[0].Name)

	_ = db.Migrator().DropTable(&entity.Book{})
	_ = db.Migrator().CreateTable(&entity.Book{})
}

func TestBook_Update(t *testing.T) {
	w := httptest.NewRecorder()

	var bookName = "sqlite"
	var renamed = "sqlite_renamed"
	book := addBook(bookName)

	book.Name = renamed
	reqBody, _ := json.Marshal(book)

	uri := fmt.Sprintf("/books/%s", book.ID)
	r := gin.Default()
	req := httptest.NewRequest(
		http.MethodPut,
		uri,
		bytes.NewReader(reqBody),
	)
	req.Header.Set("Content-Type", "application/json")

	r.PUT("/books/:id", bookHdlr.Update)
	r.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	got := &representation.Book{}
	_ = json.Unmarshal(body, &got)
	assert.Equal(t, renamed, got.Name)

	_ = db.Migrator().DropTable(&entity.Book{})
	_ = db.Migrator().CreateTable(&entity.Book{})
}

func TestBook_Delete(t *testing.T) {
	w := httptest.NewRecorder()

	var bookName = "sqlite"
	book := addBook(bookName)

	uri := fmt.Sprintf("/books/%s", book.ID)
	r := gin.Default()
	req := httptest.NewRequest(
		http.MethodDelete,
		uri,
		nil,
	)

	r.DELETE("/books/:id", bookHdlr.Delete)
	r.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	got := &representation.Book{}
	_ = json.Unmarshal(body, &got)

	_ = db.Migrator().DropTable(&entity.Book{})
	_ = db.Migrator().CreateTable(&entity.Book{})
}

func addBook(name string) *representation.Book {
	w := httptest.NewRecorder()

	uri := "/books"
	r := gin.Default()

	book := representation.Book{
		Name: name,
	}
	reqBody, _ := json.Marshal(book)
	req := httptest.NewRequest(
		http.MethodPost,
		uri,
		bytes.NewReader(reqBody),
	)
	req.Header.Set("Content-Type", "application/json")

	r.POST(uri, bookHdlr.Add)
	//r.GET(uri, bookController.GetAll)
	r.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	_ = json.Unmarshal(body, &book)
	return &book
}

func getBooks() []*representation.Book {
	w := httptest.NewRecorder()

	uri := "/books"
	r := gin.Default()
	req := httptest.NewRequest(
		http.MethodGet,
		uri,
		nil,
	)
	r.GET(uri, bookHdlr.GetAll)
	r.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)
	var books []*representation.Book
	_ = json.Unmarshal(body, &books)
	return books
}
