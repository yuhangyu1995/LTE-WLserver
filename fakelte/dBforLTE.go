package fakelte

import (
	"dkay/fake-wlserver/common"
	"dkay/fake-wlserver/fake"
	"dkay/fake-wlserver/lte"
	"errors"
	"fmt"
	"strings"
)

var (
	defaultGetRes      = "\r\nCMD       :    get %s......[OK]\r\nRESULT    :    ACK 000\r\nHINT      :    get %s success!\r\nCONTENT   :    %s\r\n"
	defaultSetRes      = "\r\nCMD       :    set %s......[OK]\r\nRESULT    :    ACK 000\r\nHINT      :    set %s success!\r\nCONTENT   :\r\n\r\n->"
	defaultAddRes      = "\r\nCMD       :    add %s......[OK]\r\nRESULT    :    ACK 000\r\nHINT      :    set %s success!\r\nCONTENT   :\r\n\r\n->"
	defaultAddFailRes  = "\r\nCMD       :    add %s......[ERR]\r\nRESULT    :    NACK 004\r\nHINT      :    Database exception: SQLITE_CONSTRAINT[19]: PRIMARY KEY must be unique\r\nCONTENT   :\r\n\r\n->"
	defaultAddEmptyRes = "\r\nCMD       :    add %s......[ERR]\r\nRESULT    :    NACK 004\r\nHINT      :    Database exception: SQLITE_CONSTRAINT[20]: datatype mismatch\r\nCONTENT   :\r\n\r\n->"
	defaultDeleteRes   = "\r\nCMD       :    delete %s......[OK]\r\nRESULT    :    ACK 000\r\nHINT      :    delete %s success!\r\nCONTENT   :\r\n\r\n->"
	listpackage        = "\r\nCMD       :    list %s......[OK]\r\nRESULT    :    ACK 000\r\nHINT      :    list %s success!\r\nCONTENT   :    \r\n"
)

//查询只有一行的表
func queryFirst(s interface{}, hostname string) error {
	d := mysql.DBer.Where("hostname=?", hostname).First(s)
	if nil != d.Error {
		common.Log.Info(d.Error)
		return d.Error
	}
	return nil
}

func queryAll(s interface{}, hostname string) error {
	d := mysql.DBer.Where("hostname=?", hostname).Find(s)
	if nil != d.Error {
		common.Log.Info(d.Error)
		return d.Error
	}
	return nil
}

//普通响应的 get的头
func getResHead(t string) string {
	var res = fmt.Sprintf(defaultGetRes, t, t, t)
	return res
}

//set的头
func setResHead(t string) string {
	var res = fmt.Sprintf(defaultSetRes, t, t)
	return res
}

func addResHead(t string) string {
	var res = fmt.Sprintf(defaultAddRes, t, t)
	return res
}

func deleteResHead(t string) string {
	var res = fmt.Sprintf(defaultDeleteRes, t, t)
	return res
}

//ArgsToMap DlATT="11" 等号后面没有空格的情况
func ArgsToMap(a string) (map[string]string, string) {
	var m map[string]string
	var re string
	m = make(map[string]string)
	sa := strings.Split(a, `, `)
	for i := 0; i < len(sa); i++ {
		//saa:=sa[i]
		//saa=saa[i:len(saa)]
		re += sa[i] + "\n"
		f := strings.Split(sa[i], "=")
		v := f[1]
		v = v[1 : len(v)-1]
		m[f[0]] = v
	}
	return m, re
}

//ArgsToMap2 BackIfID= "3" 等号后面多了一个空格的情况
func ArgsToMap2(a string) (map[string]string, string) {
	var m map[string]string
	var re string
	m = make(map[string]string)
	sa := strings.Split(a, `, `)
	for i := 0; i < len(sa); i++ {
		//saa:=sa[i]
		//saa=saa[i:len(saa)]
		re += sa[i] + "\n"
		f := strings.Split(sa[i], `= `)
		v := f[1]
		v = v[1 : len(v)-1]
		m[f[0]] = v
	}
	return m, re
}

//SetFirst ...
func SetFirst(s string, temp interface{}, hostname string) string {
	m, re := ArgsToMap(s)
	d := mysql.DBer.Model(temp).Where("hostname=?", hostname).Updates(m)
	if nil != d.Error {
		common.Log.Info(d.Error)
	}
	return re
}

func SetPLMNListInfo(s string, temp interface{}, hostname string, sn string) string {
	m, re := ArgsToMap(s)
	id := m["PLMNListInfoId"]
	if mysql.DBer.Model(temp).Where("plmn_list_info_id = ?", id).Where("hostname=?", hostname).First(temp).RecordNotFound() {
		common.Log.Info("plmn_list_info_id not found")
	} else {
		if id == "1" {
			var temp2 *lte.ServingCellInfoQueryACK
			temp22, _ := fake.CommonFunc(sn, 0xF027)
			temp2 = temp22.(*lte.ServingCellInfoQueryACK)

			pl := m["PLMNID"]
			if "" != pl {
				plbyte := []byte(pl)
				if 5 == len(plbyte) {
					for i, v := range plbyte {
						temp2.PLMN[i] = v
					}
					temp2.PLMN[5] = 0
				} else {
					return "error plmnid"
				}
			}
			var ack *lte.CommonACK
			aa, _ := fake.CommonFuncSet(sn, 0xF003, temp2)
			if aa == nil {
				errors.New("ack nil")
			} else {
				ack = aa.(*lte.CommonACK)
			}
			if ack.Result == 0 {
				common.Log.Info("修改PLMN成功")
			} else {
				common.Log.Info("修改PLMN失败")
			}
		} else {
			mysql.DBer.Model(temp).Where("plmn_list_info_id = ?", id).Where("hostname=?", hostname).Updates(m)
		}
	}
	return re
}

func addPLMNListInfo(s string, hostname string) int {
	m, _ := ArgsToMap2(s)
	id := m["PLMNListInfoId"]
	var deft = PLMNListInfo{
		"1",
		"",
		"46000",
		"1",
		"fdd-u",
	}
	deft.Hostname = hostname
	var flag = 0 //0:插入失败   1：插入成功   2：传入id为空
	if id == `` {
		flag = 2
		common.Log.Info("传入id为空")
		return flag
	}
	if mysql.DBer.Model(&deft).Where("plmn_list_info_id=?", id).Where("hostname=?", hostname).First(&deft).RecordNotFound() {
		//插入逻辑
		flag = 1
		deft.PLMNListInfoId = id
		mysql.DBer.Create(&deft)
		mysql.DBer.Model(deft).Where("plmn_list_info_id=?", id).Where("hostname=?", hostname).Updates(m)
	} else {
		//ID重复 无法插入
		common.Log.Info("ID exist")
	}
	return flag
}
func addGSMNeighborCell(s string, hostname string) int {
	m, _ := ArgsToMap2(s)
	id := m["GSMNeighborCellId"]
	var deft = GSMNeighborCell{
		"1",
		"",
		"46000",
		"34322",
		"51",
		"12976",
		"0",
		"68",
		"7",
		"1",
		"22",
		"8",
		"0",
		"30",
		"1",
		"1",
		"",
		"1",
		"1",
		"1",
		"",
		"-3",
		"fdd-u",
	}
	deft.Hostname = hostname
	var flag = 0 //0:插入失败   1：插入成功   2：传入id为空
	if id == `` {
		flag = 2
		common.Log.Info("传入id为空")
		return flag
	}
	if mysql.DBer.Model(&deft).Where("gsm_neighbor_cell_id=?", id).Where("hostname=?", hostname).First(&deft).RecordNotFound() {
		//插入逻辑
		flag = 1
		deft.GSMNeighborCellId = id
		mysql.DBer.Create(&deft)
		mysql.DBer.Model(deft).Where("gsm_neighbor_cell_id=?", id).Where("hostname=?", hostname).Updates(m)
	} else {
		//ID重复 无法插入
		common.Log.Info("ID exist")
	}
	return flag
}
func addCDMA2000NeighborCell(s string, hostname string) int {
	m, _ := ArgsToMap2(s)
	id := m["CDMA2000NeighborCellId"]
	var deft = CDMA2000NeighborCell{
		"1",
		"",
		"0",
		"2",
		"26",
		"14",
		"2",
		"15",
		"0",
		"231",
		"37",
		"fdd-u",
	}
	deft.Hostname = hostname
	var flag = 0 //0:插入失败   1：插入成功   2：传入id为空
	if id == `` {
		flag = 2
		common.Log.Info("传入id为空")
		return flag
	}
	if mysql.DBer.Model(&deft).Where("cdma2000_neighbor_cell_id=?", id).Where("hostname=?", hostname).First(&deft).RecordNotFound() {
		//插入逻辑
		flag = 1
		deft.CDMA2000NeighborCellId = id
		mysql.DBer.Create(&deft)
		mysql.DBer.Model(deft).Where("cdma2000_neighbor_cell_id=?", id).Where("hostname=?", hostname).Updates(m)
	} else {
		//ID重复 无法插入
		common.Log.Info("ID exist")
	}
	return flag
}
func addLTENeighborCell(s string, hostname string) int {
	m, _ := ArgsToMap2(s)
	id := m["LTENeighborCellId"]
	var deft = LTENeighborCell{
		"1",
		"",
		"46000",
		"1",
		"38850",
		"1",
		"-3",
		"0",
		"0",
		"0",
		"4528",
		"0x0FFFFFFF",
		"0x0FFFFFFF",
		"0",
		"1",
		"",
		"1",
		"0",
		"0",
		"4",
		"0",
		"1",
		"4",
		"0",
		"0",
		"1",
		"0",
		"fdd-u",
	}
	deft.Hostname = hostname
	var flag = 0 //0:插入失败   1：插入成功   2：传入id为空
	if id == `` {
		flag = 2
		common.Log.Info("传入id为空")
		return flag
	}
	if mysql.DBer.Model(&deft).Where("lte_neighbor_cell_id=?", id).Where("hostname=?", hostname).First(&deft).RecordNotFound() {
		//插入逻辑
		flag = 1
		deft.LTENeighborCellId = id
		mysql.DBer.Create(&deft)
		mysql.DBer.Model(deft).Where("lte_neighbor_cell_id=?", id).Where("hostname=?", hostname).Updates(m)
	} else {
		//ID重复 无法插入
		common.Log.Info("ID exist")
	}
	return flag
}

func deletePLMNListInfo(s string, hostname string) string {
	m, re := ArgsToMap2(s)
	id := m["PLMNListInfoId"]
	var temp PLMNListInfo
	mysql.DBer.Model(&temp).Where("plmn_list_info_id =?", id).Where("hostname=?", hostname).Delete(&temp)
	return re
}
func deleteGSMNeighborCell(s string, hostname string) string {
	m, re := ArgsToMap2(s)
	id := m["GSMNeighborCellId"]
	var temp GSMNeighborCell
	mysql.DBer.Model(&temp).Where("gsm_neighbor_cell_id =?", id).Where("hostname=?", hostname).Delete(&temp)
	return re
}
func deleteCDMA2000NeighborCell(s string, hostname string) string {
	m, re := ArgsToMap2(s)
	id := m["CDMA2000NeighborCellId"]
	var temp CDMA2000NeighborCell
	mysql.DBer.Model(&temp).Where("cdma2000_neighbor_cell_id =?", id).Where("hostname=?", hostname).Delete(&temp)
	return re
}
func deleteLTENeighborCell(s string, hostname string) string {
	m, re := ArgsToMap2(s)
	id := m["LTENeighborCellId"]
	var temp LTENeighborCell
	mysql.DBer.Model(&temp).Where("lte_neighbor_cell_id =?", id).Where("hostname=?", hostname).Delete(&temp)
	return re
}

