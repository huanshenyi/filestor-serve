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
		//meta.UpdateFileMeta(fileMeta)
         _ = meta.UpdateFileMetaDB(fileMeta)

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
	//fMeta := meta.GetFileMeta(filehash)
	fMeta, err := meta.GetFileMetaDB(filehash)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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

// ファイアをダウンロード
func DownloadHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha1)
	f, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
    if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// httpのheatを設定して、ブラウザがダウンロードできるようにする
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=\""+fm.FileName+"\"")

	w.Write(data)
}
// ファイア内容を更新する(rename)
func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	// opは操作のタイプ 0=rename, 1=その他
	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	if opType != "0"{
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST"{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	curFileMeta := meta.GetFileMeta(fileSha1)
	curFileMeta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)

	data, err := json.Marshal(curFileMeta)
	if err != nil{
        w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// ファイルを削除
func FileDeleteHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fileSha1 := r.Form.Get("filehash")
	fMeta := meta.GetFileMeta(fileSha1)

	err := os.Remove(fMeta.Location)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	meta.RemoveFileMeta(fileSha1)
	w.WriteHeader(http.StatusOK)
}