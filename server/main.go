package main

import (
	"log"
	"multipart-upload/server/controllers"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Upload File First!"))
}

func main() {
	addr := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	mux.HandleFunc("/upload", controllers.FileUpload)
	mux.HandleFunc("/upload_multipart", controllers.FileUploadMultipart)

	log.Printf("server is listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
