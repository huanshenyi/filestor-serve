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
    err := http.ListenAndServe(":5000", nil)
    if err != nil{
    	fmt.Printf("Failed to start server, err:%s",err)
	}
}
