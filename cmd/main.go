package main

import (
	"log"
	"net/http"
	"tracker/functions"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", functions.Home)
	mux.HandleFunc("/tracker/", functions.Posthandler)
	mux.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	log.Println("Запуск сервера на http://127.0.0.1:3000")
	err := http.ListenAndServe(":3000", mux)
	log.Println(err)
}
