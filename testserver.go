package main

import (
	"fmt"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
}
func main() {
	/*msgHandler := msg("Hello world")
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8080", msgHandler)*/

	/*fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Server is listening on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)*/

	// Обработчик для статических файлов
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Обработчик для API-запроса
	http.Handle("/api/message", msg("Hello world"))

	fmt.Println("Server is listening on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
