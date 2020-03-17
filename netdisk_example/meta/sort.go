package meta

import "time"

const baseFormat = "2016-01-02 15:04:05"

type ByUploadTime []FileMeta

func (a ByUploadTime) Len() int {
	return len(a)
}

func (a ByUploadTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByUploadTime) Less(i, j int) bool {
	iTime, _ := time.Parse(baseFormat, a[i].UpdateAt)
	jTime, _ := time.Parse(baseFormat, a[j].UpdateAt)
	return iTime.UnixNano() > jTime.UnixNano()
}
