package prodcontroller

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	tmplt, _ := template.ParseFiles("views/fileview/index.html")
	tmplt.Execute(w, nil)
}

func SingleUpload(w http.ResponseWriter, r *http.Request) {
	// Parse File data
	r.ParseMultipartForm(5 * 1024 * 1024)
	file, handler, err := r.FormFile("file")
	if err != nil || handler == nil || handler.Filename == "" {
        // Show error on the upload page
        tmplt, _ := template.ParseFiles("views/fileview/index.html")
        tmplt.Execute(w, map[string]string{"Error": "Please supply a file before trying to submit for upload."})
        return
  }
	defer file.Close()
	fmt.Println("File Name: ", handler.Filename)
	fmt.Println("File Size: ", handler.Size)
	// Save File to disk
	dst, _ := os.Create("./uploads/" + handler.Filename)
	defer dst.Close()
	io.Copy(dst, file)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func MultiUpload(w http.ResponseWriter, r *http.Request) {
	// Parse File data
	r.ParseMultipartForm(5 * 1024 * 1024)
	files := r.MultipartForm.File["files"]
	if len(files) == 0 || (len(files) == 1 && files[0].Filename == "") {
        // Show error on the upload page
        tmplt, _ := template.ParseFiles("views/fileview/index.html")
        tmplt.Execute(w, map[string]string{"Error": "Please supply at least one file before trying to submit for upload."})
        return
  }
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
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ViewFiles(w http.ResponseWriter, r *http.Request) {
	dir := "./uploads"
	tmplt, err := template.ParseGlob("views/fileview/view.html")
	if err != nil {
    http.Error(w, "Template not found: "+err.Error(), http.StatusInternalServerError)
    return
	}
	data := map[string]interface{}{
		"Dir":   dir,
		"Files": []os.FileInfo{},
	}
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil {
        fmt.Println("Walk error:", err)
        return nil
    }
		if !info.IsDir() {
			data["Files"] = append(data["Files"].([]os.FileInfo), info)
		}
		return nil
	})
	tmplt.Execute(w, data)
}
