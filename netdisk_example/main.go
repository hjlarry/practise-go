package main

import (
	"net/http"
	"netdisk_example/handler"
)

func main() {
	http.HandleFunc("/file/upload", handler.FileUploadHandle)
	http.HandleFunc("/file/upload/success", handler.FileUploadSuccessHandle)
	http.HandleFunc("/file/meta", handler.FileMetaHandle)
	http.HandleFunc("/file/query", handler.FileQueryListHandle)
	http.HandleFunc("/file/download", handler.FileDownloadHandle)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandle)
	http.HandleFunc("/file/delete", handler.FileDeleteHandle)

	http.HandleFunc("/user/signup", handler.SignUpHandle)
	http.HandleFunc("/user/signin", handler.SignInHandle)
	http.HandleFunc("/user/info", handler.HttpInterceptor(handler.UserInfoHandle))

	_ = http.ListenAndServe(":8080", nil)
}
