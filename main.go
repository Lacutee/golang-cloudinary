package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-cloudinary/cdn"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gorilla/mux"
)

func ServiceRouter(router *mux.Router) {
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode("Welcome To Server")
	}).Methods("GET")

	router.HandleFunc("/upload", func(rw http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")

		// fileName := handler.Filename

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(err)
			fmt.Println(err)
			return
		}

		defer file.Close()

		cld, err := cdn.CdnSetting()
		var ctx = context.Background()

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(err)
			fmt.Println("error open file", err)
			return
		}

		uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(err)
			fmt.Println("error read file", err)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		fmt.Println(uploadResult)
		json.NewEncoder(rw).Encode(uploadResult.SecureURL)
	}).Methods("POST")
}

func StartRouter() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println(`Running on port 8090`)
	ServiceRouter(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func main() {
	StartRouter()
}
