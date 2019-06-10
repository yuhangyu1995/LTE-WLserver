package fake

import (
	"dkay/fake-wlserver/common"
	"github.com/json-iterator/go"
	"sync"
)

var (
	log   =common.Log
	app   = &common.AppContext
	dev   = &common.AppContext.RealDev
	json                   = jsoniter.ConfigCompatibleWithStandardLibrary
	mysql                  = &common.AppContext.DB.Mysql
	memlist                = &common.AppContext.DB.MemoryList
	checkError             = common.CheckError
	printError             = common.PrintError
	bytetoString           = common.BytetoString
	readBufUseLittleEndian = common.ReadBufUseLittleEndian
)

type receiver struct {
	Mu     sync.RWMutex
	revMap map[string]interface{}
}

var rec receiver //存放对应的key和结构对象

//查询只有一行的表
func queryFirst(s interface{}) error{
	d := mysql.DBer.First(s)
	if nil != d.Error {
		common.Log.Info(d.Error)
		return d.Error
	}
	return nil
}

//查询有多行的表
func queryAll(s interface{}) error{
	d := mysql.DBer.Find(s)
	if nil != d.Error {
		common.Log.Info(d.Error)
		return d.Error
	}
	return nil
}

//修改只有一行的表
func SetFirst(s string,temp interface{})string{
	m,re:=ArgsToMap(s)
	d:=mysql.DBer.Model(temp).Updates(m)
	if nil != d.Error {
		common.Log.Info(d.Error)
	}
	return re
}

//根据backifid修改具体某一行的值
func SetSpecified(s string,temp interface{})string{
	m,re:=ArgsToMap(s)
	id:=m["BackIfID"]
	if mysql.DBer.Model(temp).Where("back_if_id = ?", id).First(temp).RecordNotFound(){
		common.Log.Info("back_if_id not found")
	}else{
		mysql.DBer.Model(temp).Where("back_if_id = ?", id).Updates(m)
	}
	return re
}

//只针对文件上传设置页面，插入行。用flag来标志 能插入/不能插入/传入id为空
func addBackIfParam(s string)(string,int){
	m,re:=ArgsToMap2(s)
	id:=m["BackIfID"]
	var flag=0 //0:插入失败   1：插入成功   2：传入id为空
	var deft=BackIfParam{
		"",
		"0",
		"1",
		"120",
		"180",
		"",
		"",
		"",
		"",
		"0",
		"",
		"",
	}
	if id==``{
		flag=2
		common.Log.Info("back_if_id 传入id为空")
		return re,flag
	}
	if mysql.DBer.Model(&deft).Where("back_if_id = ?", id).First(&deft).RecordNotFound(){
		//插入逻辑
		flag=1
		deft.BackIfID=id
		mysql.DBer.Create(&deft)
		mysql.DBer.Model(deft).Where("back_if_id = ?", id).Updates(m)
	}else{
		//ID 重复无法插入
		common.Log.Info("ID exist")
	}
	return re,flag
}

//只针对文件上传设置页面，删除行
func deleteBackIfParam(s string)string{
	m,re:=ArgsToMap2(s)
	id:=m["BackIfID"]
	var temp BackIfParam
	mysql.DBer.Model(&temp).Where("back_if_id = ?", id).Delete(&temp)
	return re
}

//还原数据库变量初始值
func SetDefault (adress interface{},temp interface{}){
	d:=mysql.DBer.Model(adress).Updates(temp)
	if nil != d.Error {
		common.Log.Info(d.Error)
	}
}

func GetHostInfo(sn string) HostInfo{
	rows, err := mysql.DBer.DB().Query("SELECT ID,SN,IP,Port,Name,Alias,Type,PLMN,RebootSW,RebootStartTime,RebootEndTime,StartTime,CanReboot from host_info where SN='"+sn+"'")
	if nil != err {
		log.Panic(err)
	}
	var temp HostInfo
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.SN, &temp.IP, &temp.Port, &temp.Name, &temp.Alias, &temp.Type, &temp.PLMN, &temp.RebootSW, &temp.RebootStartTime, &temp.RebootEndTime, &temp.StartTime,&temp.CanReboot)
	}
	return temp
}

func GetAllHostInfo() {
	rows, err := mysql.DBer.DB().Query("SELECT ID,SN,IP,Port,Name,Alias,Type,PLMN,RebootSW,RebootStartTime,RebootEndTime,StartTime,CanReboot from host_info")
	if nil != err {
		log.Panic(err)
	}
	for rows.Next() {
		var temp HostInfo
		rows.Scan(&temp.ID, &temp.SN, &temp.IP, &temp.Port, &temp.Name, &temp.Alias, &temp.Type, &temp.PLMN, &temp.RebootSW, &temp.RebootStartTime, &temp.RebootEndTime, &temp.StartTime, &temp.CanReboot)
		HostInfoMap[temp.IP] = temp
	}
}
