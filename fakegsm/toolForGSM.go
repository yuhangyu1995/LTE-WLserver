package fakegsm

import (
	"dkay/fake-wlserver/fake"
	"fmt"
	"strings"
)

var (
	workstate        = `{"resp":"1","content":{"mixData":[%s]},"hint":"0"}`
	rebootCtrlInfo   = `{"resp":"1","content":{"singleData":[%s]},"hint":"0"}`
	targetList       = `{"resp":"1","content":{"singleData":[%s,{"RejectList":""}]},"hint":"0"}`
	carrierInfoSet   = `{"resp":"1","content":{"multiData":[%s]},"hint":"0"}`
	currentAlarmInfo = `{"resp":"1","content":{"multiData":[%s]},"hint":"0"}`
)

func GetResStringForGSM(cmd string, param string, hostname string, sn string) string {
	var res string
	switch cmd {
	case "query_class_table":
		res = `{"resp":"1","content":"","hint":"0"}`
	case "query_version":
		res = `{"resp":"1","content":{"currentVersion":"4MS_HSG.21.09.1","backupVersion":"4MS_HSG.21.09.0"},"hint":"0"}`
	case "mix_query":
		if strings.Contains(param, "RunningTime") {
			t := fmt.Sprintf(workstate, StructToStr(GetGsmWorkState(hostname, sn)))
			s := strings.Replace(t, "Ba1List", "ba1List", 1)
			res = strings.Replace(s, "Ba2List", "ba2List", 1)
		}
		if strings.Contains(param, "LanIpAddress") {
			t := fmt.Sprintf(workstate, StructToStr(GetGsmBasicSetting(hostname, sn)))
			s := strings.Replace(t, "Switch2", "Switch", 1)
			s2 := strings.Replace(s, "Switch3", "Switch", 1)
			res = strings.Replace(s2, "IpAddress2", "IpAddress", 1)
		}
	case "config":
		if strings.Contains(param, "TargetList") { //cmd: config ;params: [TargetList](AccessList="112312121222321,121254345489111")
			ConfigTargetList(param, "1", hostname)
			res = `{"resp":"1","content":"","hint":"0"}`
		}
		if strings.Contains(param, "CarrierInfoSet") { //cmd: config ;params: [TargetList](AccessList="112312121222321,121254345489111")
			ConfigCarrierInfoSet(param, "1", hostname)
			res = `{"resp":"1","content":"","hint":"0"}`
		}
		//---------------------------------默认情况，基本设置的修改
		ParamToUpdate(param, hostname, sn) //cmd: config  params: [SonMgmt](InterferenceTh="-61")
		res = `{"resp":"1","content":"","hint":"0"}`
		//------------------------------------------------------------

	case "query":
		if strings.Contains(param, "TargetInfo") {
			res = `{"resp":"1","content":{"singleData":[{"MainCatchImsiInfo":""}]},"hint":"0"}`
		}
		if strings.Contains(param, "TargetList") {
			temp, _ := GetTargetList(hostname)
			res = fmt.Sprintf(targetList, temp)
			//{"resp":"1","content":{"singleData":[{"AccessList":"111112222233333,124598789874563,124598789874563"},{"RejectList":""}]},"hint":"0"}
		}
		if strings.Contains(param, "RebootCtrlInfo") {
			var re RebootCtrlInfo
			var temp fake.HostInfo
			temp=fake.GetHostInfo(sn)
			re.RebootSwitch=temp.RebootSW
			re.RebootTimeRangesStart=temp.RebootStartTime
			re.RebootTimeRangesEnd=temp.RebootEndTime

			res = fmt.Sprintf(rebootCtrlInfo, StructToStr(re))
			//{"resp":"1","content":{"singleData":[{"RebootSwitch":"1"},{"RebootTimeRangesStart":"02:00"},{"RebootTimeRangesEnd":"05:00"}]},"hint":"0"}
		}

	case "multi_query":
		if strings.Contains(param, "CarrierInfoSet") {
			//-----------------------------cmd=multi_query&params=%5BCarrierInfoSet%5D
			var t string
			var temp []GSMCarrierInfoSet
			mysql.DBer.Where("hostname=?", hostname).Find(&temp)
			for i := 0; i < len(temp)-1; i++ {
				tt, _ := json.Marshal(temp[i])
				t = t + string(tt) + ","
			}
			tt, _ := json.Marshal(temp[len(temp)-1])
			t += string(tt)
			res = fmt.Sprintf(carrierInfoSet, t)
		}

		if strings.Contains(param, "CurrentAlarmInfo") {
			//当前状态告警{"resp":"1","content":{"multiData":[{"InstanceId":"1","AlarmId":"10007","AlarmCnName":"侦听邻区失败","AlarmRaiseCause":"no ncell record","AlarmRaiseTime":"1970-01-01 00:00:39"}]},"hint":"0"}
			var t string
			var temp []CurrentAlarmInfo
			mysql.DBer.Where("hostname=?", hostname).Find(&temp)
			for i := 0; i < len(temp)-1; i++ {
				tt, _ := json.Marshal(temp[i])
				t = t + string(tt) + ","
			}
			tt, _ := json.Marshal(temp[len(temp)-1])
			t += string(tt)
			res = fmt.Sprintf(currentAlarmInfo, t)
		}

		if strings.Contains(param, "HistoryAlarmInfo") {
			//历史告警
			var t string
			var temp []HistoryAlarmInfo
			mysql.DBer.Where("hostname=?", hostname).Find(&temp)
			for i := 0; i < len(temp)-1; i++ {
				tt, _ := json.Marshal(temp[i])
				t = t + string(tt) + ","
			}
			tt, _ := json.Marshal(temp[len(temp)-1])
			t += string(tt)
			res = fmt.Sprintf(currentAlarmInfo, t)
		}

	case "add":
		if strings.Contains(param, "CarrierInfoSet") {
			AddCarrierInfoSet(param, "1", hostname)
			res = `{"resp":"1","content":"","hint":"0"}`
		}

	case "delete":
		//[CarrierInfoSet](CarrierId="1",CarrierId="0")
		if strings.Contains(param, "CarrierInfoSet") {
			w := strings.Split(param, `](`)
			e := w[1]
			e = e[0 : len(e)-2]
			ee := strings.Split(e, `",`)
			for i := 0; i < len(ee); i++ {
				t := strings.Split(ee[i], `="`)
				d := t[1]
				var temp GSMCarrierInfoSet
				mysql.DBer.Model(&temp).Where("carrier_id = ?", d).Where("id=?", "1").Where("hostname=?", hostname).Delete(&temp)
			}
			res = `{"resp":"1","content":"","hint":"0"}`
		}

	case "query_work_dir":
		res = `{"resp":"1","content":"/flash/appsys/Pack1","hint":"0"}`

	case "reboot":
		res = `{"resp":"1","content":"","hint":"0"}`
		fake.RebootGSM(sn)

	case "reboot_state":
		res = `{"resp":"1","content":"","hint":"0"}`
	}
	return res
}

