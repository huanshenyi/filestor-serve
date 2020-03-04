package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET"{
		// TODOアップロードのhtmlページを返す
		file, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil{
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(file))
	}else if r.Method == "POST"{
		// ファイルのストリームを受け取って保存
		file, head, err := r.FormFile("file")
		if err != nil{
			fmt.Printf("Failed to get data, err%s", err.Error())
			return
		}
		defer file.Close()
		dir, _ := os.Getwd()
		newFile, err := os.Create(dir+"\\tmp\\"+head.Filename)
		if err != nil{
			fmt.Printf("Failed to create file, err%s", err.Error())
			return
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save into file, err:%s\n", err.Error())
			return
		}
		http.Redirect(w,r, "/file/upload/suc", http.StatusFound)
	}
}

// アップロード完了の知らせ用
func UploadSucHandler(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "Upload finished!")
}