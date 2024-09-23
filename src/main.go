package main

import (
	"fmt"
	"net/http"

	"github.com/C-Keenan/FileServer/src/controllers/testcontroller"
)

func main() {
	http.HandleFunc("/upload", testcontroller.Upload)
	http.HandleFunc("/upload/singleupload", testcontroller.SingleUpload)
	http.HandleFunc("/upload/multiupload", testcontroller.MultiUpload)
	http.HandleFunc("/", testcontroller.ViewFiles)
	http.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("./uploads"))))
	port := ":3000"
	fmt.Printf("Server is listening on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
