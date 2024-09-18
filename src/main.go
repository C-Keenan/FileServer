package main

import (
	"fmt"
	"net/http"

	"github.com/C-Keenan/FileServer/src/controllers/testcontroller"
)

func main() {
	http.HandleFunc("/upload", testcontroller.Index1)
	http.HandleFunc("/upload/singleupload", testcontroller.SingleUpload)
	http.HandleFunc("/upload/multiupload", testcontroller.MultiUpload)
	http.HandleFunc("/view", testcontroller.ViewFiles)
	port := ":3000"
	fmt.Printf("Server is listening on localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
