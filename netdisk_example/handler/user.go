package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	dblayer "netdisk_example/db"
	"netdisk_example/util"
	"time"
)

const (
	pwdSalt   = "hh123"
	tokenSalt = "_tokensalt"
)

func SignUpHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			fmt.Printf("Read File Error: %s \n", err.Error())
			return
		}
		_, _ = w.Write(data)
	}
	_ = r.ParseForm()

	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	if len(username) < 3 || len(passwd) < 5 {
		_, _ = w.Write([]byte("Invalid parameter"))
		return
	}

	encPasswd := util.Sha1([]byte(passwd + pwdSalt))
	suc := dblayer.UserSignUp(username, encPasswd)
	if suc {
		_, _ = w.Write([]byte("SUCCESS"))
	} else {
		_, _ = w.Write([]byte("FAILED"))
	}
}

func SignInHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/signin.html")
		if err != nil {
			fmt.Printf("Read File Error: %s \n", err.Error())
			return
		}
		_, _ = w.Write(data)
		return
	}
	_ = r.ParseForm()

	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	encPasswd := util.Sha1([]byte(passwd + pwdSalt))
	suc := dblayer.UserSignIn(username, encPasswd)
	if !suc {
		_, _ = w.Write([]byte("FAILED"))
		return
	}

	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + tokenSalt))
	token := tokenPrefix + ts[:8]
	suc = dblayer.UpdateToken(username, token)
	if !suc {
		_, _ = w.Write([]byte("FAILED"))
		return
	}

	res := util.RespMsg{
		Code: 0,
		Msg:  "ok",
		Data: struct {
			Username string
			Token    string
			Location string
		}{
			Username: username,
			Token:    token,
			Location: "http://" + r.Host + "/static/view/home.html",
		},
	}
	_, _ = w.Write(res.JSONBytes())
}

func UserInfoHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	username := r.Form.Get("username")
	user, err := dblayer.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	res := util.RespMsg{
		Code: 0,
		Msg:  "ok",
		Data: user,
	}
	_, _ = w.Write(res.JSONBytes())

}
