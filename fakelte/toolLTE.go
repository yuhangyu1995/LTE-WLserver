package fakelte

import (
	"dkay/fake-wlserver/common"
	"dkay/fake-wlserver/fake"
	"dkay/fake-wlserver/lte"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func structToString(obj interface{}) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var str string
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name != "Id" && t.Field(i).Name != "Hostname" {
			str += fmt.Sprintf(`%s = "%v"`+"\r\n", t.Field(i).Name, v.Field(i).Interface())
		}

	}
	return str
}

func GetResStringLTE(cmd string, hostname string, sn string) string {
	m := strings.Split(cmd, ";")
	var res, args string
	for _, v := range m {
		if i := strings.Index(v, ":"); i > 0 {
			args = v[i+1:]
			v = v[:i]
			args = strings.TrimPrefix(args, " ")
			v = strings.TrimSuffix(v, " ")
		}
		//用于处理前段多个cmd的情况
		if strings.HasPrefix(v, "\017") {
			v = v[len("\017"):]
			res += "\017"
		}
		switch v {
		case "get WebConfig":
			res = `
CMD       :    get WebConfig......[ERR]
RESULT    :    NACK 004
HINT      :    get WebConfig not defined!
CONTENT   :    

-> `
		case "web getFile":
			if args == "path=/OAM/software/web/web_page/IMEI.imei" {
				quertSQL := fmt.Sprintf("select No,imsi,imei,time,mode from imsi_imei where sn='%v';", sn)
				rows, _ := mysql.DBer.DB().Query(quertSQL)
				var imsiStr string
				for rows.Next() {
					var No int64
					var imsi, imei, time, mode string
					rows.Scan(&No, &imsi, &imei, &time, &mode)
					imsiStr += fmt.Sprintf("NO:%v IMSI:%v IMEI:%v TIME:%v MODE:%v\n", No, imsi, imei, time, mode)
				}
				res += "\r\nCMD    :    web web command......[OK]\r\nRESULT    :    ACK 000\r\nHINT    :    \r\nCONTENT:\r\ndate = \"" + imsiStr + "\" \r\n\r\n\r\n->\r\n"
			}
		case "get AlarmDictionary":
			str, err := ioutil.ReadFile("./dist/conf/AlarmDictionaryLTE.txt")
			if err != nil {
				res = "get AlarmDictionaryLTE failed."
			}
			res = string(str)
		case "get MetaData":
			str, err := ioutil.ReadFile("./dist/conf/MetaDataLTE.txt")
			if err != nil {
				res = "get MetaData failed."
			}
			res = string(str)
		case "get DeviceInfo":
			var temp DeviceInfo
			queryFirst(&temp, hostname)
			res = res + getResHead("DeviceInfo") + structToString(temp) + "\n\n" + `-> `
		case "get RFControl":
			var temp RFControl
			queryFirst(&temp, hostname)
			res = res + getResHead("RFControl") + structToString(temp) + "\n\n" + `-> `
		case "get ClockSynTemp":
			var temp ClockSynTemp
			queryFirst(&temp, hostname)

			var temp2 *lte.SyncInfoQueryACK
			temp22, _ := fake.CommonFunc(sn, 0xF02D)
			if temp22==nil{

			}else{
				temp2 = temp22.(*lte.SyncInfoQueryACK)
				temp.CurrentSynMode = fmt.Sprintf("%v", temp2.SyncMode)
				temp.CurrentSynStatus = fmt.Sprintf("%v", temp2.State)

				res = res + getResHead("ClockSynTemp") + structToString(temp) + "\n\n" + `-> `
			}

		case "get DeviceStatus":
			var temp DeviceStatus
			queryFirst(&temp, hostname)
			//RunningTime = "0 00:29:14"
			var startTime int64
			rows, err := mysql.DBer.DB().Query("SELECT StartTime FROM host_info where SN='" + sn + "'")
			if nil != err {
				common.Log.Info("主控IP查询失败")
			}
			for rows.Next() {
				rows.Scan(&startTime)
			}
			nowTime := time.Now().Unix()
			day := (nowTime - startTime) / (60 * 60 * 24)
			hour := (nowTime - startTime - day*60*60*24) / (60 * 60)
			min := (nowTime - startTime - day*60*60*24 - hour*60*60) / 60
			sec := nowTime - startTime - day*60*60*24 - hour*60*60 - min*60
			temp.RunningTime = fmt.Sprintf("%d %02d:%02d:%02d", day, hour, min, sec)

			res = res + getResHead("DeviceStatus") + structToString(temp) + "\n\n" + `-> `
		case "get MainCtrlParam":
			var temp MainCtrlParam
			queryFirst(&temp, hostname)
			res = res + getResHead("MainCtrlParam") + structToString(temp) + "\n\n" + `-> `
		case "get ScanSelfOptParam":
			var temp ScanSelfOptParam
			queryFirst(&temp, hostname)
			res = res + getResHead("ScanSelfOptParam") + structToString(temp) + "\n\n" + `-> `
		case "get LAN":
			var temp LAN
			queryFirst(&temp, hostname)

			var mainIP string
			var IP string
			rows, err := mysql.DBer.DB().Query("SELECT IP FROM host_info where Name='main'")
			if nil != err {
				log.Info("主控IP查询失败")
			}
			for rows.Next() {
				rows.Scan(&mainIP)
			}
			rows2, err := mysql.DBer.DB().Query("SELECT IP FROM host_info where SN='" + sn + "'")
			if nil != err {
				log.Info("主控IP查询失败")
			}
			for rows2.Next() {
				rows2.Scan(&IP)
			}
			temp.IPAddress = IP
			temp.DefaultGateway = mainIP
			temp.AssistDNSServer = mainIP
			temp.MainDNSServer = mainIP

			res = res + getResHead("LAN") + structToString(temp) + "\n\n" + `-> `
		case "get ImsiRouteCtrlParam":
			var temp ImsiRouteCtrlParam
			queryFirst(&temp, hostname)

			var mainIP string
			rows, err := mysql.DBer.DB().Query("SELECT IP FROM host_info where Name='main'")
			if nil != err {
				log.Info("主控IP查询失败")
			}
			for rows.Next() {
				rows.Scan(&mainIP)
			}
			temp.IpAddress = mainIP

			res = res + getResHead("ImsiRouteCtrlParam") + structToString(temp) + "\n\n" + `-> `
		case "get PCCH":
			var temp PCCH
			queryFirst(&temp, hostname)
			res = res + getResHead("PCCH") + structToString(temp) + "\n\n" + `-> `
		case "get BCCH":
			var temp BCCH
			queryFirst(&temp, hostname)
			res = res + getResHead("BCCH") + structToString(temp) + "\n\n" + `-> `
		case "get ClockSyn":
			var temp ClockSyn
			queryFirst(&temp, hostname)
			res = res + getResHead("ClockSyn") + structToString(temp) + "\n\n" + `-> `
		case "get RFConfig":
			var temp RFConfig
			queryFirst(&temp, hostname)

			var temp2 *lte.ServingCellInfoQueryACK
			temp22, _ := fake.CommonFunc(sn, 0xF027)
			temp2 = temp22.(*lte.ServingCellInfoQueryACK)
			temp.PhyCellId = fmt.Sprintf("%v", temp2.PCI)
			temp.EARFCNDL = fmt.Sprintf("%v", temp2.DLEarfcn)
			temp.EARFCNUL = fmt.Sprintf("%v", temp2.DLEarfcn)
			temp.FrequencyBandIndicator = fmt.Sprintf("%v", temp2.Band)

			res = res + getResHead("RFConfig") + structToString(temp) + "\n\n" + `-> `
		case "get EUTRAN":
			var temp EUTRAN
			queryFirst(&temp, hostname)
			res = res + getResHead("EUTRAN") + structToString(temp) + "\n\n" + `-> `
		case "get EPC":
			var temp EPC
			queryFirst(&temp, hostname)
			res = res + getResHead("EPC") + structToString(temp) + "\n\n" + `-> `
		case "get FAPControl":
			var temp FAPControl
			queryFirst(&temp, hostname)
			res = res + getResHead("FAPControl") + structToString(temp) + "\n\n" + `-> `
		case "get TacIncrease":
			var temp TacIncrease
			queryFirst(&temp, hostname)
			res = res + getResHead("TacIncrease") + structToString(temp) + "\n\n" + `-> `
		case "get SonConfigParam":
			var temp SonConfigParam
			queryFirst(&temp, hostname)
			res = res + getResHead("SonConfigParam") + structToString(temp) + "\n\n" + `-> `
		case "get TraceFunParam":
			var temp TraceFunParam
			queryFirst(&temp, hostname)
			res = res + getResHead("TraceFunParam") + structToString(temp) + "\n\n" + `-> `
		case "get PhyUlMonitorOpen":
			var temp PhyUlMonitorOpen
			queryFirst(&temp, hostname)
			res = res + getResHead("PhyUlMonitorOpen") + structToString(temp) + "\n\n" + `-> `
		case "get RemoteMgmtParam":
			var temp RemoteMgmtParam
			queryFirst(&temp, hostname)

			var mainIP string
			rows, err := mysql.DBer.DB().Query("SELECT IP FROM host_info where Name='main'")
			if nil != err {
				log.Info("主控IP查询失败")
			}
			for rows.Next() {
				rows.Scan(&mainIP)
			}
			temp.IpAddress = mainIP

			res = res + getResHead("RemoteMgmtParam") + structToString(temp) + "\n\n" + `-> `
		case "get MasterSlaveLteIndParam":
			var temp MasterSlaveLteIndParam
			queryFirst(&temp, hostname)
			res = res + getResHead("MasterSlaveLteIndParam") + structToString(temp) + "\n\n" + `-> `

		case "get NetworkTransCtrlParam":
			var temp NetworkTransCtrlParam
			queryFirst(&temp, hostname)

			//var temp2 lte.TAURejectCauseCfg
			//temp22, _ := CommonFunc(sn, 0xF057)
			//temp2 = temp22.(lte.TAURejectCauseCfg)
			//temp.TAURejectCause = fmt.Sprintf("%v", temp2.Cause)

			res = res + getResHead("NetworkTransCtrlParam") + structToString(temp) + "\n\n" + `-> `

		case "get LTENeighborCellInUse":
			res = res + getResHead("LTENeighborCellInUse") + `-> `
		case "list package":
			var temp Packagee
			queryFirst(&temp, hostname)
			res = res + fmt.Sprintf(listpackage, "package", "package") + structToString(temp) + "\n\n" + `-> `
		case "get RebootControl":
			var temp RebootControl
			queryFirst(&temp, hostname)
			var temp2 fake.HostInfo
			temp2 = fake.GetHostInfo(sn)
			temp.RebootEnable = temp2.RebootSW
			temp.RebootTime = temp2.RebootStartTime + "-" + temp2.RebootEndTime
			res = res + getResHead("RebootControl") + structToString(temp) + "\n\n" + `-> `
			//-----------------------------------------------------------------------------------------------------------------
		case "set MainCtrlParam":
			var temp MainCtrlParam
			SetFirst(args, &temp, hostname)
			res = res + setResHead("MainCtrlParam")

		case "set ScanSelfOptParam":
			var temp ScanSelfOptParam
			SetFirst(args, &temp, hostname)
			res = res + setResHead("ScanSelfOptParam")
		case "set LAN":
			var temp LAN
			SetFirst(args, &temp, hostname)
			res = res + setResHead("LAN")
		case "set ImsiRouteCtrlParam":
			var temp ImsiRouteCtrlParam
			SetFirst(args, &temp, hostname)
			res = res + setResHead("ImsiRouteCtrlParam")
		case "set PCCH":
			var temp PCCH
			SetFirst(args, &temp, hostname)
			res = res + setResHead("PCCH")
		case "set BCCH":
			var temp BCCH
			SetFirst(args, &temp, hostname)
			res = res + setResHead("BCCH")
		case "set ClockSyn":
			var temp ClockSyn
			SetFirst(args, &temp, hostname)
			res = res + setResHead("ClockSyn")
		case "set RFConfig":
			var temp RFConfig

			var temp2 *lte.ServingCellInfoQueryACK
			temp22, _ := fake.CommonFunc(sn, 0xF027)
			temp2 = temp22.(*lte.ServingCellInfoQueryACK)
			m, _ := ArgsToMap(args)
			//s1,_ :=strconv.Atoi(m["FrequencyBandIndicator"])
			//temp2.Band=uint32(s1)
			if _, ok := m["PhyCellId"]; ok {
				s2, _ := strconv.Atoi(m["PhyCellId"])
				temp2.PCI = uint16(s2)
			}
			if _, ok := m["EARFCNDL"]; ok {
				earfcdl := m["EARFCNDL"]
				s3, _ := strconv.Atoi(earfcdl)
				if hostname == "tdd" {
					if s3 >= 37750 && s3 <= 41589 {
						temp2.DLEarfcn = uint32(s3)
					} else {
						common.Log.Info("频点范围错误")
					}
					if s3 >= 37750 && s3 <= 38249 {
						temp2.Band = uint32(38)
					}
					if s3 >= 38250 && s3 <= 38649 {
						temp2.Band = uint32(39)
					}
					if s3 >= 38650 && s3 <= 39649 {
						temp2.Band = uint32(40)
					}
					if s3 >= 39650 && s3 <= 41589 {
						temp2.Band = uint32(41)
					}
				}
				if hostname == "fdd-t" || hostname == "fdd-u" {
					if s3 >= 0 && s3 <= 599 {
						temp2.Band = uint32(1)
						temp2.ULEarfcn = uint32(s3 + 18000)
						temp2.DLEarfcn = uint32(s3)
					} else if s3 >= 1200 && s3 <= 1949 {
						temp2.Band = uint32(3)
						temp2.ULEarfcn = uint32(s3 - 1200 + 19200)
						temp2.DLEarfcn = uint32(s3)
					} else if s3 >= 2750 && s3 <= 3449 {
						temp2.Band = uint32(7)
						temp2.ULEarfcn = uint32(s3 - 2750 + 20750)
						temp2.DLEarfcn = uint32(s3)
					} else {
						common.Log.Info("频点范围错误")
					}
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
				SetFirst(args, &temp, hostname)
				res = res + setResHead("RFConfig")
			} else {
				res = "\r\nCMD       :    set RFConfig......[ERR]\r\nRESULT    :    NACK 004\r\nHINT      :    set RFConfig not defined!\r\nCONTENT   :\r\n\r\n->"
			}

		case "set EUTRAN":
			var temp EUTRAN
			SetFirst(args, &temp, hostname)
			res = res + setResHead("EUTRAN")
		case "set FAPControl":
			var temp FAPControl
			SetFirst(args, &temp, hostname)
			res = res + setResHead("FAPControl")
		case "set EPC":
			var temp EPC
			SetFirst(args, &temp, hostname)
			res = res + setResHead("EPC")

		case "set TacIncrease":
			var temp TacIncrease
			SetFirst(args, &temp, hostname)
			res = res + setResHead("TacIncrease")
		case "set SonConfigParam":
			var temp SonConfigParam
			SetFirst(args, &temp, hostname)
			res = res + setResHead("SonConfigParam")
		case "set TraceFunParam":
			var temp TraceFunParam
			SetFirst(args, &temp, hostname)
			res = res + setResHead("TraceFunParam")
		case "set PhyUlMonitorOpen":
			var temp PhyUlMonitorOpen
			SetFirst(args, &temp, hostname)
			res = res + setResHead("PhyUlMonitorOpen")
		case "set RemoteMgmtParam":
			var temp RemoteMgmtParam
			SetFirst(args, &temp, hostname)
			res = res + setResHead("RemoteMgmtParam")
		case "set MasterSlaveLteIndParam":
			var temp MasterSlaveLteIndParam
			SetFirst(args, &temp, hostname)
			res = res + setResHead("MasterSlaveLteIndParam")

		case "set NetworkTransCtrlParam":
			var temp NetworkTransCtrlParam
			SetFirst(args, &temp, hostname)
			res = res + setResHead("NetworkTransCtrlParam")

			//var temp2 lte.TAURejectCauseCfg
			//temp22, _ := CommonFunc(sn, 0xF06B)
			//temp2 = temp22.(lte.TAURejectCauseCfg)
			//m, _ := ArgsToMap(args)
			//s1, _ := strconv.Atoi(m["TAURejectCause"])
			//temp2.Cause = uint32(s1)
			//
			//var ack lte.CommonACK
			//aa, _ := CommonFuncSet(sn, 0xF057, temp2)
			//ack = aa.(lte.CommonACK)
			//if ack.Result == 0 {
			//	SetFirst(args, &temp, hostname)
			//	res = res + setResHead("NetworkTransCtrlParam")
			//} else {
			//	res = "update fail"
			//	//to do
			//}
			//---------------------------------------------------------------------------------------------------
		case "get GSMNeighborCell":
			var temp []GSMNeighborCell
			queryAll(&temp, hostname)
			res = res + getResHead("GSMNeighborCell") + structToString(temp[0]) + "\r\n"
			for i := 1; i < len(temp); i++ {
				res = res + "CONTENT   :    GSMNeighborCell\r\n" + structToString(temp[i]) + "\r\n"
			}
			res = res + `-> `
		case "set GSMNeighborCell":
			var temp GSMNeighborCell
			SetFirst(args, &temp, hostname)
			res = res + setResHead("GSMNeighborCell ")
		case "add GSMNeighborCell":
			flag := addGSMNeighborCell(args, hostname)
			if flag == 1 {
				//插入成功
				res = res + addResHead("GSMNeighborCell ")
			} else if flag == 0 {
				//插入失败 ID重复
				res = res + fmt.Sprintf(defaultAddFailRes, "GSMNeighborCell")
			} else {
				//传入ID为空
				log.Info("id empty fail")
			}

		case "delete GSMNeighborCell":
			deleteGSMNeighborCell(args, hostname)
			res = res + deleteResHead("GSMNeighborCell")
			//-----------------------------------------------------------------------------------------------------
		case "get PLMNListInfo":
			var temp []PLMNListInfo
			queryAll(&temp, hostname)
			//plmn := bytetoString(msg.PLMN[:])
			var temp2 *lte.ServingCellInfoQueryACK
			temp22, _ := fake.CommonFunc(sn, 0xF027)
			temp2 = temp22.(*lte.ServingCellInfoQueryACK)
			plmn := common.BytetoString(temp2.PLMN[:])
			temp[0].PLMNID = plmn

			res = res + getResHead("PLMNListInfo") + structToString(temp[0]) + "\r\n"
			for i := 1; i < len(temp); i++ {
				res = res + "CONTENT   :    PLMNListInfo\r\n" + structToString(temp[i]) + "\r\n"
			}
			res = res + `-> `
		case "set PLMNListInfo":
			var temp PLMNListInfo
			SetPLMNListInfo(args, &temp, hostname, sn)
			res = res + setResHead("PLMNListInfo ")
		case "add PLMNListInfo":
			flag := addPLMNListInfo(args, hostname)
			if flag == 1 {
				//插入成功
				res = res + addResHead("PLMNListInfo ")
			} else if flag == 0 {
				//插入失败 ID重复
				res = res + fmt.Sprintf(defaultAddFailRes, "PLMNListInfo")
			} else {
				//传入ID为空
				res = res + fmt.Sprintf(defaultAddEmptyRes, "PLMNListInfo")
			}

		case "delete PLMNListInfo":
			deletePLMNListInfo(args, hostname)
			res = res + deleteResHead("PLMNListInfo")
			//-----------------------------------------------------------------------------------------------------
		case "get CDMA2000NeighborCell":
			var temp []CDMA2000NeighborCell
			queryAll(&temp, hostname)
			res = res + getResHead("CDMA2000NeighborCell") + structToString(temp[0]) + "\r\n"
			for i := 1; i < len(temp); i++ {
				res = res + "CONTENT   :    CDMA2000NeighborCell\r\n" + structToString(temp[i]) + "\r\n"
			}
			res = res + `-> `
		case "set CDMA2000NeighborCell":
			var temp CDMA2000NeighborCell
			SetFirst(args, &temp, hostname)
			res = res + setResHead("CDMA2000NeighborCell ")
		case "add CDMA2000NeighborCell":
			flag := addCDMA2000NeighborCell(args, hostname)
			if flag == 1 {
				//插入成功
				res = res + addResHead("CDMA2000NeighborCell ")
			} else if flag == 0 {
				//插入失败 ID重复
				res = res + fmt.Sprintf(defaultAddFailRes, "CDMA2000NeighborCell")
			} else {
				//传入ID为空
				log.Info("id empty fail")
			}

		case "delete CDMA2000NeighborCell":
			deleteCDMA2000NeighborCell(args, hostname)
			res = res + deleteResHead("CDMA2000NeighborCell")
			//----------------------------------------------------------------------------------------------
		case "get LTENeighborCell":
			var temp []LTENeighborCell
			queryAll(&temp, hostname)
			res = res + getResHead("LTENeighborCell") + structToString(temp[0]) + "\r\n"
			for i := 1; i < len(temp); i++ {
				res = res + "CONTENT   :    LTENeighborCell\r\n" + structToString(temp[i]) + "\r\n"
			}
			res = res + `-> `
		case "set LTENeighborCell":
			var temp LTENeighborCell
			SetFirst(args, &temp, hostname)
			res = res + setResHead("LTENeighborCell ")
		case "add LTENeighborCell":
			flag := addLTENeighborCell(args, hostname)
			if flag == 1 {
				//插入成功
				res = res + addResHead("LTENeighborCell ")
			} else if flag == 0 {
				//插入失败 ID重复
				res = res + fmt.Sprintf(defaultAddFailRes, "LTENeighborCell")
			} else {
				//传入ID为空
				log.Info("id empty fail")
			}

		case "delete LTENeighborCell":
			deleteLTENeighborCell(args, hostname)
			res = res + deleteResHead("CDMA2000NeighborCell")
			//-----------------------------------------------------------------------------------------
		case " get TargetImsiParam":
			var temp TargetImsiParam
			queryFirst(&temp, hostname)
			res = res + getResHead("TargetImsiParam") + structToString(temp) + "\n\n" + `-> `
		case "set TargetImsiParam":
			var temp TargetImsiParam
			SetFirst(args, &temp, hostname)
			res = res + setResHead("TargetImsiParam")
			//----------------------------------------------------
		case "set RebootControl":
			m, _ := ArgsToMap(args)
			if _, ok := m["RebootEnable"]; ok {
				stmt, err := mysql.DBer.DB().Prepare("UPDATE host_info SET RebootSW =? where SN=?")
				if nil != err {
					log.Info("更新rebootEnable失败")
				}
				_, err = stmt.Exec(m["RebootEnable"], sn)
				if nil != err {
					log.Info("更新rebootEnable失败")
				}
			}
			if _, ok := m["RebootTime"]; ok {
				time := strings.Split(m["RebootTime"], "-")
				stmt, err := mysql.DBer.DB().Prepare("UPDATE host_info SET RebootStartTime =? , RebootEndTime=? where SN=?")
				if nil != err {
					log.Info("更新rebootEnable失败")
				}
				_, err = stmt.Exec(time[0], time[1], sn)
				if nil != err {
					log.Info("更新rebootEnable失败")
				}
			}
			mysql.DBer.DB().Exec("UPDATE host_info SET CanReboot ='1' where SN='" + sn + "'")
			res = res + setResHead("RebootControl")
		case "reboot":
			//RebootCfg  0xF00B
			res = "\r\nCMD       :    reboot......[OK]\r\nRESULT    :    ACK 000\r\nHINT      :    reboot success!\r\nCONTENT   :\r\n\r\n->"
			fake.RebootLTE(sn)
		default:
		}
	}
	return res

}
