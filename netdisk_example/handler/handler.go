package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	dblayer "netdisk_example/db"
	"netdisk_example/meta"
	"netdisk_example/util"
	"os"
	"strconv"
	"time"
)

func FileUploadHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			fmt.Printf("Read File Error: %s \n", err.Error())
			return
		}
		_, _ = io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data: %s \n", err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "/tmp/" + head.Filename,
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		newfile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create file: %s \n", err.Error())
			return
		}
		defer newfile.Close()

		fileMeta.FileSize, err = io.Copy(newfile, file)
		if err != nil {
			fmt.Printf("Failed to copy: %s \n", err.Error())
			return
		}

		_, _ = newfile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newfile)

		dblayer.OnFileUploadFinished(fileMeta.FileSha1, fileMeta.FileName, fileMeta.Location, fileMeta.FileSize)

		http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	}
}

func FileUploadSuccessHandle(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "success upload!")
}

func FileMetaHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	filehash := r.Form["filehash"][0]
	fMeta := meta.GetMeta(filehash)
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
}

func FileQueryListHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	limitCnt, _ := strconv.Atoi(r.Form.Get("limit"))
	fileMetas := meta.GetLastFileMetasDB(limitCnt)
	data, err := json.Marshal(fileMetas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
}

func FileDownloadHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fm := meta.GetMeta(fsha1)

	f, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-Descrption", "attachment;filename=\""+fm.FileName+"\"")
	_, _ = w.Write(data)
}

func FileMetaUpdateHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	opType := r.Form.Get("op")
	fsha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	if opType != "0" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	curFileMeta := meta.GetMeta(fsha1)
	curFileMeta.FileName = newFileName
	//meta.UpdateMeta(curFileMeta)

	data, _ := json.Marshal(curFileMeta)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func FileDeleteHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fMeta := meta.GetMeta(fsha1)

	_ = os.Remove(fMeta.Location)
	meta.RemoveFileMeta(fsha1)
	w.WriteHeader(http.StatusOK)
}
