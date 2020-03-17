package meta

import (
	dblayer "netdisk_example/db"
	"sort"
)

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UpdateAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

func GetMeta(fsha1 string) FileMeta {
	tmeta, _ := dblayer.GetFileMeta(fsha1)
	fmeta := FileMeta{
		FileSha1: tmeta.TfileHash,
		FileName: tmeta.TfileName.String,
		FileSize: tmeta.TfileSize.Int64,
		Location: tmeta.TfileAddr.String,
	}
	return fmeta
}

func GetLastFileMetasDB(limit int) []FileMeta {
	fMetaArray := make([]FileMeta, 0, len(fileMetas))
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}
	sort.Sort(ByUploadTime(fMetaArray))

	return fMetaArray[:limit]
}

func RemoveFileMeta(fsha1 string) {
	delete(fileMetas, fsha1)
}
