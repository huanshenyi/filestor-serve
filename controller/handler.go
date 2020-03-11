package controller

import (
	"encoding/json"
	"filestor-serve/meta"
	"filestor-serve/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
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
		fileMeta := meta.FileMeta{
			FileName:head.Filename,
			Location:dir+"\\tmp\\"+head.Filename,
			UploadAt:time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location)
		if err != nil{
			fmt.Printf("Failed to create file, err%s", err.Error())
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save into file, err:%s\n", err.Error())
			return
		}

		newFile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(w,r, "/file/upload/suc", http.StatusFound)
	}
}

// アップロード完了の知らせ用
func UploadSucHandler(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "Upload finished!")
}

// アプロード完了したファイアのデータを取得
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data, err := json.Marshal(fMeta)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// 複数のファイアデータを取得
func FileQueryHandler(w http.ResponseWriter, r *http.Request)  {
   r.ParseForm()
   limitCnt, _ := strconv.Atoi(r.Form.Get("limit"))
   fileMetas := meta.GetLastFileMetas(limitCnt)
   data, err := json.Marshal(fileMetas)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}