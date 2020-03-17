package db

import (
	"database/sql"
	"fmt"
	mydb "netdisk_example/db/mysql"
)

func OnFileUploadFinished(fsha1, fname, faddr string, fsize int64) bool {
	stmt, err := mydb.DBConn().Prepare("insert ignore into tbl_file (`file_sha1`,`file_name`,`file_size`, `file_addr`,`status`) values (?,?,?,?,1)")
	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(fsha1, fname, fsize, faddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	rf, err := ret.RowsAffected()
	if err != nil {
		return false
	}
	if rf <= 0 {
		fmt.Printf("File with hash:%s has been uploaded before", fsha1)
	}
	return true

}

type TableMeta struct {
	TfileHash string
	TfileName sql.NullString
	TfileSize sql.NullInt64
	TfileAddr sql.NullString
}

func GetFileMeta(fsha1 string) (*TableMeta, error) {
	stmt, err := mydb.DBConn().Prepare("select file_name,file_size,file_addr  from tbl_file where file_sha1 = ? and status=1 limit 1")
	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return nil, err
	}
	defer stmt.Close()

	tmeta := TableMeta{TfileHash: fsha1}
	err = stmt.QueryRow(fsha1).Scan(&tmeta.TfileName, &tmeta.TfileSize, &tmeta.TfileAddr)
	if err != nil {
		if err == sql.ErrNoRows {
			// 查不到对应记录， 返回参数及错误均为nil
			return nil, nil
		} else {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	return &tmeta, nil
}
