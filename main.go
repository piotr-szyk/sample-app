package main
import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/blue", blueHandler)
	http.HandleFunc("/red", redHandler)
	http.ListenAndServe(":8080", nil)
}
func blueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body style='background-color:blue;'></body></html>")
}
func redHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body style='background-color:red;'></body></html>")
}
// version 2.0.1
