package main

import (
	"fmt"
	"net/http"

	"github.com/C-Keenan/FileServer/controllers/prodcontroller"
)

func main() {
	http.HandleFunc("/upload", prodcontroller.Upload)
	http.HandleFunc("/upload/singleupload", prodcontroller.SingleUpload)
	http.HandleFunc("/upload/multiupload", prodcontroller.MultiUpload)
	http.HandleFunc("/", prodcontroller.ViewFiles)
	http.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("./uploads"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	port := ":3000"
	fmt.Printf("Server is listening on %s.\n", port)
	http.ListenAndServe(port, nil)
}
