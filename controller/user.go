package controller

import (
	"filestor-serve/db"
	"filestor-serve/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

// ログインコントローラー
func SignlnHandler(w http.ResponseWriter, r *http.Request){

  r.ParseForm()
  username := r.Form.Get("username")
  password := r.Form.Get("password")

  encPasswd := util.Sha1([]byte(password + pwd_salt))

  // 1. ユーザーネームとパスワードの検証
  pwdChecked := db.UserSignin(username, encPasswd)
  if !pwdChecked {
  	  w.Write([]byte("FAILED"))
	  return
  }

  // 2. tokenを生成する
  token := GenToken(username)
  upRes := db.UpdateToken(username, token)
  if !upRes {
	  w.Write([]byte("FAILED"))
	  return
  }
  // 3. ホームページへリダイレクト
  w.Write([]byte("http://"+r.Host+"/static/view/home.html"))

}

// Tokenを生成
func GenToken(username string) string {
	//長さ40: md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username+ts+"_tokensalt"))
	return tokenPrefix + ts[:8]
}