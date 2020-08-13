package home



import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)
var tpl *template.Template
// Compile templates on start of the application
func init(){
tpl = template.Must(template.ParseGlob("src/views/*.html"))
}
// Display the named template
func display(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w,"index",nil)

}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, r)
	case "POST":
		uploadFile(w, r)
	}
}

