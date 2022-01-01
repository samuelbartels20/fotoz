// Declares what package this code is part of 
package main

/*
Tells the compiler which packages i intend to use in my code
*/
import (
	"fmt"
	"net/http"
)

/* 
This is the function i use to process incoming web requests
*/
func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>This is the begining of my journey in Go programming</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3001", nil)
}

