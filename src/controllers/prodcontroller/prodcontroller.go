package testcontroller

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	tmplt, _ := template.ParseFiles("views/testcontroller/index1.html")
	tmplt.Execute(w, nil)
}

func SingleUpload(w http.ResponseWriter, r *http.Request) {
	// Parse File data
	r.ParseMultipartForm(5 * 1024 * 1024)
	file, handler, _ := r.FormFile("file")
	defer file.Close()
	fmt.Println("File Name: ", handler.Filename)
	fmt.Println("File Size: ", handler.Size)
	// Save File to disk
	dst, _ := os.Create("./uploads/" + handler.Filename)
	defer dst.Close()
	io.Copy(dst, file)
	tmplt, _ := template.ParseFiles("views/prodcontroller/index.html")
	tmplt.Execute(w, nil)
}

func MultiUpload(w http.ResponseWriter, r *http.Request) {
	// Parse File data
	r.ParseMultipartForm(5 * 1024 * 1024)
	files := r.MultipartForm.File["files"]
	fmt.Println("Files: ", len(files))
	for i, fileHandler := range files {
		fmt.Println("File ", i)
		fmt.Println("File Name: ", fileHandler.Filename)
		fmt.Println("File Size: ", fileHandler.Size)
		file, _ := files[i].Open()
		defer file.Close()
		dst, _ := os.Create("./uploads/" + fileHandler.Filename)
		defer dst.Close()
		io.Copy(dst, file)
	}
	tmplt, _ := template.ParseFiles("views/fileview/index.html")
	tmplt.Execute(w, nil)
}

func ViewFiles(w http.ResponseWriter, r *http.Request) {
	dir := "./uploads"
	tmplt, _ := template.ParseGlob("views/fileview/index.html")
	data := map[string]interface{}{
		"Dir":   dir,
		"Files": []os.FileInfo{},
	}
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			data["Files"] = append(data["Files"].([]os.FileInfo), info)
		}
		return nil
	})
	tmplt.Execute(w, data)
}
