package main

import (
	"filestor-serve/controller"
	"fmt"
	"net/http"
)

func main()  {
    http.HandleFunc("/file/upload", controller.UploadHandler)
    http.HandleFunc("/file/upload/suc", controller.UploadSucHandler)
    http.HandleFunc("/file/meta", controller.GetFileMetaHandler)
    http.HandleFunc("/file/query", controller.FileQueryHandler)
    http.HandleFunc("/file/download", controller.DownloadHandler)
    err := http.ListenAndServe(":5000", nil)
    if err != nil{
    	fmt.Printf("Failed to start server, err:%s",err)
	}
}
