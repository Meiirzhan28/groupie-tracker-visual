package functions

import (
	"fmt"
	"net/http"
	"text/template"
)

type Errors struct {
	Message string
	Status  int
}

func Errorhandler(w http.ResponseWriter, status int) {
	html, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Internal Server Errror 500")
		return
	}
	var Result Errors
	Result.Message = http.StatusText(status) // func StatusText()
	Result.Status = status
	w.WriteHeader(status)
	err = html.Execute(w, Result)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Internal Server Errror 500")
		return
	}
}
