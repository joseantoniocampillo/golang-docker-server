package main
import (
	"net/http"
)
func main() {
	 fileServer := http.FileServer(http.Dir("public"))
	 http.Handle("/", http.StripPrefix("/",fileServer))
	 http.ListenAndServe(":8001", nil)
 }
