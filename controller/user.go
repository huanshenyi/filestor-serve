package controller

import (
	"filestor-serve/db"
	"filestor-serve/util"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	pwd_salt = "*#890"
)

// 新規ユーザー用のコントローラー
func SignupHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	fmt.Println("username:"+username, "password:"+passwd)
	if len(username)<3 || len(passwd)<5 {
		w.Write([]byte("Invalid parameter"))
		return
	}
	enc_passwd := util.Sha1([]byte(passwd+pwd_salt))
	suc := db.UserSignup(username, enc_passwd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}
