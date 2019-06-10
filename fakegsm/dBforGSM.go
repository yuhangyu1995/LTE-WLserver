package fakegsm

import (
	"dkay/fake-wlserver/common"
	"dkay/fake-wlserver/fake"
	"dkay/fake-wlserver/gsm"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var (
	dev   = &common.AppContext.RealDev
	app   = &common.AppContext
	json  = jsoniter.ConfigCompatibleWithStandardLibrary
	mysql = &common.AppContext.DB.Mysql
)

func QuerryOne(key string, Type string, hostname string) (string, error) {
	var re = ""
	rows, err := mysql.DBer.DB().Query("SELECT val FROM gsm_basic_info where ki='" + key + "' and ty='" + Type + "' and hostname='" + hostname + "'")
	if nil != err {
		return "", err
	}
	for rows.Next() {
		rows.Scan(&re)
	}
	return re, nil
}

func UpdateOne(key string, Type string, val string, hostname string) error {
	stmt, err := mysql.DBer.DB().Prepare("UPDATE gsm_basic_info SET val =? where ki=? and ty=? and hostname=?")
	if nil != err {
		return err
	}
	_, err = stmt.Exec(val, key, Type, hostname)
	if nil != err {
		return err
	}
	return nil
}

func GetGsmWorkState(hostname string, sn string) GsmWorkState {
	var re GsmWorkState
	//re.RunningTime,_=QuerryOne("RunningTime","SystemStateParam",hostname)
	var startTime int64
	rows, err := mysql.DBer.DB().Query("SELECT StartTime FROM host_info where SN='"+sn+"'")
	if nil != err {
		common.Log.Info("主控IP查询失败")
	}
	for rows.Next() {
		rows.Scan(&startTime)
	}
	nowTime:=time.Now().Unix()
	day := (nowTime-startTime)/(60*60*24)
	hour := (nowTime-startTime-day*60*60*24)/(60*60)
	min := (nowTime-startTime-day*60*60*24-hour*60*60)/60
	sec := nowTime-startTime-day*60*60*24-hour*60*60-min*60
	re.RunningTime = fmt.Sprintf("%d %02d:%02d:%02d", day, hour, min, sec)

	sns := strings.Split(sn, "-")
	sn1 := sns[0]
	sn2 := sns[1]
	re.CellState, _ = QuerryOne("CellState", "SystemStateParam", hostname)

	//re.Arfcn,_=QuerryOne("Arfcn","RfMgmt",hostname)
	var cp *gsm.CellPara
	c, _ := fake.CommonFunc(sn1, "GetCellPara")
	cp = c.(*gsm.CellPara)
	for _, v := range cp.Cell.Item {
		mc := fmt.Sprintf("%v", v.Mnc)
		if sn2 == mc {
			re.Arfcn = fmt.Sprintf("%v", v.ArfcnList.Arfcn)
			re.Mcc = fmt.Sprintf("%v", v.Mcc)
			re.Mnc = mc
		}
	}

	re.CurrentClockSyncMode, _ = QuerryOne("CurrentClockSyncMode", "SystemStateParam", hostname)
	re.CurrentClockSyncState, _ = QuerryOne("CurrentClockSyncState", "SystemStateParam", hostname)

	re.SignalSwitch, _ = QuerryOne("SignalSwitch", "SignalCtrlParam", hostname)
	re.CellId, _ = QuerryOne("CellId", "CellIdentityInfo", hostname)

	re.TimeSyncState, _ = QuerryOne("TimeSyncState", "SystemStateParam", hostname)

	//re.Mcc,_=QuerryOne("Mcc","CellIdentityInfo",hostname)
	//re.Mnc,_=QuerryOne("Mnc","CellIdentityInfo",hostname)

	re.Ncc, _ = QuerryOne("Ncc", "CellIdentityInfo", hostname)
	re.Bcc, _ = QuerryOne("Bcc", "CellIdentityInfo", hostname)
	re.Lac, _ = QuerryOne("Lac", "CellIdentityInfo", hostname)
	re.Ba1List, _ = QuerryOne("ba1List", "NeighbourCellInfo", hostname)
	re.Ba2List, _ = QuerryOne("ba2List", "NeighbourCellInfo", hostname)
	re.ForbitLacList, _ = QuerryOne("ForbitLacList", "SonMgmt", hostname)

	return re
}

func GetGsmBasicSetting(hostname string, sn string) GsmBasicSetting {
	var re GsmBasicSetting
	var mainIP string
	var IP string
	rows, err := mysql.DBer.DB().Query("SELECT IP FROM host_info where Name='main'")
	if nil != err {
		common.Log.Info("主控IP查询失败")
	}
	for rows.Next() {
		rows.Scan(&mainIP)
	}
	rows2, err := mysql.DBer.DB().Query("SELECT IP FROM host_info where SN='"+sn+"'")
	if nil != err {
		common.Log.Info("主控IP查询失败")
	}
	for rows2.Next() {
		rows2.Scan(&IP)
	}

	//re.LanIpAddress, _ = QuerryOne("LanIpAddress", "LANSettingInfo", hostname)
	re.LanIpAddress=IP
	//re.OutputDataIp, _ = QuerryOne("OutputDataIp", "OutputDataInfo", hostname)
	re.OutputDataIp=mainIP
	re.OutputDataPort, _ = QuerryOne("OutputDataPort", "OutputDataInfo", hostname)
	re.Switch, _ = QuerryOne("Switch", "ExternalInterfaceInfo", hostname)
	//re.IpAddress, _ = QuerryOne("IpAddress", "ExternalInterfaceInfo", hostname)
	re.IpAddress=mainIP
	re.Port, _ = QuerryOne("Port", "ExternalInterfaceInfo", hostname)

	re.Type, _ = QuerryOne("Type", "ExternalInterfaceInfo", hostname)
	//re.Mcc,_=QuerryOne("Mcc","CellIdentityInfo",hostname)
	//re.Mnc,_=QuerryOne("Mnc","CellIdentityInfo",hostname)
	re.Lac, _ = QuerryOne("Lac", "CellIdentityInfo", hostname)
	re.CellId, _ = QuerryOne("CellId", "CellIdentityInfo", hostname)
	re.Ncc, _ = QuerryOne("Ncc", "CellIdentityInfo", hostname)
	re.Bcc, _ = QuerryOne("Bcc", "CellIdentityInfo", hostname)
	re.LacHoppingPeriod, _ = QuerryOne("LacHoppingPeriod", "LacHoppingInfo", hostname)
	re.LacRange, _ = QuerryOne("LacRange", "LacHoppingInfo", hostname)
	re.Switch2, _ = QuerryOne("Switch", "SonMgmt", hostname)

	re.Band, _ = QuerryOne("Band", "RfMgmt", hostname)
	re.SonBand, _ = QuerryOne("SonBand", "SonMgmt", hostname)
	re.InterferenceTh, _ = QuerryOne("InterferenceTh", "SonMgmt", hostname)
	re.AvailableArfcnList, _ = QuerryOne("AvailableArfcnList", "SonMgmt", hostname)
	re.ForbitSonArfcnList, _ = QuerryOne("ForbitSonArfcnList", "SonMgmt", hostname)

	re.SignalSwitch, _ = QuerryOne("SignalSwitch", "SignalCtrlParam", hostname)
	re.BcchTxPowerDbm, _ = QuerryOne("BcchTxPowerDbm", "CellIdentityInfo", hostname)
	//re.Arfcn,_=QuerryOne("Arfcn","RfMgmt",hostname)
	re.MaxTxPower, _ = QuerryOne("MaxTxPower", "CellIdentityInfo", hostname)
	re.MinTxPower, _ = QuerryOne("MinTxPower", "CellIdentityInfo", hostname)
	sns := strings.Split(sn, "-")
	sn1 := sns[0]
	sn2 := sns[1]
	var cp *gsm.CellPara
	c, _ := fake.CommonFunc(sn1, "GetCellPara")
	cp = c.(*gsm.CellPara)
	for _, v := range cp.Cell.Item {
		mc := fmt.Sprintf("%v", v.Mnc)
		if sn2 == mc {
			re.Arfcn = fmt.Sprintf("%v", v.ArfcnList.Arfcn)
			re.Mcc = fmt.Sprintf("%v", v.Mcc)
			re.Mnc = mc
		}
	}

	re.Period, _ = QuerryOne("Period", "IncomeDataCheck", hostname)
	re.Switch3, _ = QuerryOne("Switch", "RemoteCtrlInfo", hostname)
	//re.IpAddress2, _ = QuerryOne("IpAddress", "RemoteCtrlInfo", hostname)
	re.IpAddress2=mainIP

	return re
}

func GetRebootCtrlInfo(hostname string) RebootCtrlInfo {
	var re RebootCtrlInfo
	re.RebootSwitch, _ = QuerryOne("RebootSwitch", "RebootCtrlInfo", hostname)
	re.RebootTimeRangesEnd, _ = QuerryOne("RebootTimeRangesEnd", "RebootCtrlInfo", hostname)
	re.RebootTimeRangesStart, _ = QuerryOne("RebootTimeRangesStart", "RebootCtrlInfo", hostname)
	return re
}

func StructToStr(obj interface{}) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var str string
	for i := 0; i < t.NumField()-1; i++ {
		str += fmt.Sprintf(`{"%s":"%v"}`+",", t.Field(i).Name, v.Field(i).Interface())
	}
	str += fmt.Sprintf(`{"%s":"%v"}`, t.Field(t.NumField()-1).Name, v.Field(t.NumField()-1).Interface())
	return str
}

func ParamToUpdate(pa string, hostname string, sn string) {
	sns := strings.Split(sn, "-")
	sn1 := sns[0]
	sn2 := sns[1]
	var cp *gsm.CellPara
	var index int
	c, _ := fake.CommonFunc(sn1, "GetCellPara")
	cp = c.(*gsm.CellPara)
	for i, v := range cp.Cell.Item {
		mc := fmt.Sprintf("%v", v.Mnc)
		if sn2 == mc {
			index = i
		}
	}
	t := strings.Split(pa, "](")
	var ty = t[0]
	var key string
	var val string
	ty = ty[1:]
	tt := t[1]
	tt = tt[0 : len(tt)-2]
	t2 := strings.Split(tt, `",`)
	//`[RfMgmt](Band="1",name="user",pass="123")`
	for i := 0; i < len(t2); i++ {
		t3 := strings.Split(t2[i], `="`)
		key = t3[0]
		val = t3[1]
		UpdateOne(key, ty, val, hostname)
		if key == "Mcc" {
			t, _ := strconv.Atoi(t3[1])
			cp.Cell.Item[index].Mcc = t
		}
		if key == "Mnc" {
			t, _ := strconv.Atoi(t3[1])
			cp.Cell.Item[index].Mnc = t
		}
		if key == "Arfcn" {
			t, _ := strconv.Atoi(t3[1])
			cp.Cell.Item[index].ArfcnList.Arfcn = t
		}
		if key == "RebootSwitch" {
			stmt, err := mysql.DBer.DB().Prepare("UPDATE host_info SET RebootSW =? where SN=?")
			if nil != err {
				common.Log.Info("更新rebootEnable失败")
			}
			_, err = stmt.Exec(val, sn)
			if nil != err {
				common.Log.Info("更新rebootEnable失败")
			}
		}
		if key == "RebootTimeRangesStart" {
			stmt, err := mysql.DBer.DB().Prepare("UPDATE host_info SET RebootStartTime =? where SN=?")
			if nil != err {
				common.Log.Info("更新rebootEnable失败")
			}
			_, err = stmt.Exec(val, sn)
			if nil != err {
				common.Log.Info("更新rebootEnable失败")
			}
			mysql.DBer.DB().Exec("UPDATE host_info SET CanReboot ='1' where SN='" + sn + "'")
		}
		if key == "RebootTimeRangesEnd" {
			stmt, err := mysql.DBer.DB().Prepare("UPDATE host_info SET RebootEndTime =? where SN=?")
			if nil != err {
				common.Log.Info("更新rebootEnable失败")
			}
			_, err = stmt.Exec(val, sn)
			if nil != err {
				common.Log.Info("更新rebootEnable失败")
			}
			mysql.DBer.DB().Exec("UPDATE host_info SET CanReboot ='1' where SN='" + sn + "'")
		}
	}
	fake.CommonFuncSet(sn1, "SetCellPara", cp)
	//`[RfMgmt](Band="1",name="user",pass="123")`
}

func GetTargetList(hostname string) (string, error) {
	var res string
	var temp string
	res += `{"AccessList":"`
	rows, err := mysql.DBer.DB().Query("SELECT target_imsi FROM gsm_targetlist where hostname='" + hostname + "'")
	if nil != err {
		return "", err
	}
	for rows.Next() {
		rows.Scan(&temp)
		res = res + temp + ","
	}
	res = res[0 : len(res)-1]
	res += `"}`
	return res, nil
}

func ConfigTargetList(para string, id string, hostname string) error {
	//params: [TargetList](AccessList="112312121222321,121254345489111")
	stmt, err := mysql.DBer.DB().Prepare("delete from gsm_targetlist where id=? and hostname=?")
	if nil != err {
		return err
	}
	_, err = stmt.Exec(id, hostname)
	if nil != err {
		return err
	}
	s := strings.Split(para, `="`)
	ss := s[1]
	ss = ss[0 : len(ss)-2]
	sss := strings.Split(ss, ",")
	for i := 0; i < len(sss); i++ {
		stmt, err := mysql.DBer.DB().Prepare("insert into gsm_targetlist values (?,?,?)")
		if nil != err {
			return err
		}
		_, err = stmt.Exec(id, sss[i], hostname)
		if nil != err {
			return err
		}
	}
	return nil
}

func ConfigCarrierInfoSet(para string, id string, hostname string) error {
	//[CarrierInfoSet](CarrierArfcn="91",BcchInd="0",CarrierId="0")
	var s map[string]string
	s = make(map[string]string)
	var key string
	var val string
	t := strings.Split(para, "](")
	tt := t[1]
	tt = tt[0 : len(tt)-2]
	t2 := strings.Split(tt, `",`)
	for i := 0; i < len(t2); i++ {
		t3 := strings.Split(t2[i], `="`)
		key = t3[0]
		val = t3[1]
		s[key] = val
	}
	var temp GSMCarrierInfoSet
	mysql.DBer.Model(&temp).Where("carrier_id = ?", s["CarrierId"]).Where("id = ?", id).Where("hostname=?", hostname).Updates(s)
	return nil
}

func AddCarrierInfoSet(para string, id string, hostname string) error {
	var temp = GSMCarrierInfoSet{id,
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}
	temp.HostName = hostname
	var s map[string]string
	s = make(map[string]string)
	var key string
	var val string
	t := strings.Split(para, "](")
	tt := t[1]
	tt = tt[0 : len(tt)-2]
	t2 := strings.Split(tt, `",`)
	for i := 0; i < len(t2); i++ {
		t3 := strings.Split(t2[i], `="`)
		key = t3[0]
		val = t3[1]
		s[key] = val
	}
	temp.CarrierId = s["CarrierId"]
	mysql.DBer.Create(&temp)
	mysql.DBer.Model(&temp).Where("carrier_id = ?", s["CarrierId"]).Where("id = ?", id).Where("hostname=?", hostname).Updates(s)
	return nil
	//[CarrierInfoSet](CarrierId="3",CarrierArfcn="78",BcchInd="1",TrxType="1",EdgeCapable="0",CarrierState="1")
}
