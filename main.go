package main
import (
    "fmt"
    "net/http"
"os"
"home"
)
func uploadFile(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Uploading File")
}
func setupRoutes() {

}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	
	mux.HandleFunc("/upload", home.Index)
	
	http.ListenAndServe(":"+port, mux)
}