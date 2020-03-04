package controller

import (
	"io"
	"io/ioutil"
	"net/http"
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
	}
}
