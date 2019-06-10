package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"reflect"
)

//CheckError exit
func CheckError(err error) {
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
}

//PrintError println
func PrintError(err error) {
	if nil != err {
		fmt.Println(err.Error())
	}
}

//ReadBufUseLittleEndian ...
func ReadBufUseLittleEndian(buf []byte, data interface{}) error {
	buff := bytes.NewBuffer(buf)
	err := binary.Read(buff, binary.LittleEndian, data)
	return err
}

//WriteBufUseLittleEndian ...
func WriteBufUseLittleEndian(data interface{}) []byte {
	buff := new(bytes.Buffer)
	err := binary.Read(buff, binary.LittleEndian, data)
	PrintError(err)
	return buff.Bytes()
}

//BytetoString ...
func BytetoString(buf []byte) string {
	index := bytes.IndexByte(buf, 0)
	if index < 0 {
		index = 0
	}
	return string(buf[0:index])
}

//StructToString 用于将struct 转成特殊的字符串
//只能用于结构体
func StructToString(obj interface{}) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var str string
	for i := 0; i < t.NumField(); i++ {
		str += fmt.Sprintf(`%s= "%v"`+"\r\n", t.Field(i).Name, v.Field(i).Interface())
	}
	return str
}
