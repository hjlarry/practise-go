package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)


func sayhelloName(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	_, _ = fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		_, _ = io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.html")
		_ = t.Execute(w, token)
		//log.Println(t.Execute(w, nil))
		//v := url.Values{}
		//v.Set("name", "Ava")
		//v.Add("friend", "111")
		//v.Add("friend", "Sarah")
		//v.Add("friend", "Zoe")
		//// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
		//fmt.Println(v.Get("name"))
		//fmt.Println(v.Get("friend"))
		//fmt.Println(v["friend"])
	} else {
		_ = r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// to valid token
		} else {
			//	error
		}
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", r.Form["password"])
		t, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		_ = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		_, _ = io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("upload.html")
		_ = t.Execute(w, token)
	} else {
		_ = r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		_, _ = fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		_, _ = io.Copy(f, file)
	}
}

func testHTTP() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("listen And Serve:", err)
	}
}

