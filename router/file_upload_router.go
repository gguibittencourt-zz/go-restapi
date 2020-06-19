package router

import (
	"gopkg.in/mgo.v2"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func FileUpload(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	fileName := r.FormValue("file_name")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.WriteString(w, "File "+fileName+" Uploaded successfully")
	_, _ = io.Copy(f, file)
}

func ServeFromDB(w http.ResponseWriter, r *http.Request) {
	var gridfs *mgo.GridFS // Obtain GridFS via Database.GridFS(prefix)

	fileName := r.FormValue("file_name")
	f, err := gridfs.Open(fileName)
	if err != nil {
		log.Printf("Failed to open %s: %v", fileName, err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	http.ServeContent(w, r, fileName, time.Now(), f) // Use proper last mod time
}
