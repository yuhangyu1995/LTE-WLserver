package common

import (
	"bytes"
	"fmt"

	"github.com/jlaffaye/ftp"
)

// FTPConfigure ...
type FTPConfigure struct {
	Flag        bool   `json:"flag"`
	UploadCycle int    `json:"uploadCycle"`
	ServerIP    string `json:"serverIP"`
	Port        string `json:"port"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	FilePath    string `json:"filePath"`
}

func ftpUploadFile(ftpserver, ftpuser, pw, remoteSavePath, saveName string, value *bytes.Reader) {
	ftp, err := ftp.Connect(ftpserver)

	if err != nil {
		fmt.Println(err)
	}
	err = ftp.Login(ftpuser, pw)
	if err != nil {
		fmt.Println(err)
	}
	//注意是 pub/log，不能带“/”开头
	ftp.ChangeDir("pub/log")
	dir, err := ftp.CurrentDir()
	fmt.Println(dir)
	ftp.MakeDir(remoteSavePath)
	ftp.ChangeDir(remoteSavePath)
	dir, _ = ftp.CurrentDir()
	fmt.Println(dir)
	err = ftp.Stor(saveName, value)
	if err != nil {
		fmt.Println(err)
	}
	ftp.Logout()
	ftp.Quit()
	fmt.Println("success upload file:")
}

func ftpUpload() {

}
