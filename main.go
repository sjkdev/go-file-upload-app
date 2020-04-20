package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading File\n")

	// 1. parse input, type multi-aprt/ form-data
	r.ParseMultipartForm(10 << 20)

	// 2. retrieve file from psted form-data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Upload File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// 3. Write temporary file on our server
	tempFile, err := ioutil.TempFile("temp-images",
		"upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	// 4. retrun whether pr not this has been succesful
	fmt.Fprintf(w, "Succesfuly Uploaded File\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8040", nil)
}

func main() {
	fmt.Println("Go File Upload Tutorial")
	setupRoutes()
}

// package main

// import "fmt"

// func main() {
// 	fmt.Printf("hello, world\n")
// }
