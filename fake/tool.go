package fake

import (
	"dkay/fake-wlserver/common"
	"dkay/fake-wlserver/gsm"
	"dkay/fake-wlserver/lte"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	gsm_on_in_min   = -10.6
	gsm_on_in_max   = -8.8
	gsm_on_out_min  = -9.6
	gsm_on_out_max  = -7.8
	gsm_on_rfx_min  = -8.6
	gsm_on_rfx_max  = -6.8
	gsm_off_in_min  = -10.8
	gsm_off_in_max  = -8.7
	gsm_off_out_min = -9.8
	gsm_off_out_max = -7.7
	gsm_off_rfx_min = -8.8
	gsm_off_rfx_max = -6.7

	ltdd_on_in_min   = -30.1
	ltdd_on_in_max   = -23.3
	ltdd_on_out_min  = 38.1
	ltdd_on_out_max  = 40.2
	ltdd_on_rfx_min  = 15.4
	ltdd_on_rfx_max  = 26.3
	ltdd_off_in_min  = -30.8
	ltdd_off_in_max  = -24.0
	ltdd_off_out_min = -6.2
	ltdd_off_out_max = -6.2
	ltdd_off_rfx_min = -2.2
	ltdd_off_rfx_max = -2.2

	lfdd_on_in_min   = -44.8
	lfdd_on_in_max   = -20.6
	lfdd_on_out_min  = 36.5
	lfdd_on_out_max  = 40.2
	lfdd_on_rfx_min  = 21.0
	lfdd_on_rfx_max  = 29.9
	lfdd_off_in_min  = -46.8
	lfdd_off_in_max  = -20.8
	lfdd_off_out_min = -2.1
	lfdd_off_out_max = -2.1
	lfdd_off_rfx_min = -5.0
	lfdd_off_rfx_max = -5.0
)

//概率返回随机数，60%返回（-high,high）,30%返回（-mid,mid）,10%返回（-low,low），保留小数点后两位
func change_value(high float64, mid float64, low float64) float64 {
	var num float64
	var r float64
	num = rand.Float64()
	if num < 0.6 {
		r = high
	}
	if num > 0.6 && num < 0.9 {
		r = mid
	}
	if num > 0.9 {
		r = low
	}
	return math.Trunc((rand.Float64()-0.5)*2*r*100+0.5) * 1e-2
}

var (
	defaulClires203  = "\r\nCMD: \r\nRESULT: 203\r\nHINT: \"%v\" is not defined.\r\nCONTENT: \r\n-> "
	defaulClires000  = "\r\nCMD: \r\nRESULT: 000\r\nHINT: %v success\r\nCONTENT: \r\n-> "
	defaulCliRes     = "\r\nCMD: %v\r\nRESULT: %v\r\nHINT: %v\r\nCONTENT: %v\r\n"
	defaulOneLineRes = "CONTENT: %v\r\n"

	pingSuccess = `TIME:%v RESULT:Success SUMMARY:NETWORK DIAG: Get WAN IP [192.168.1.76] Success
TIME:%v RESULT:Success SUMMARY:NETWORK DIAG: Get Route default Gateway[192.168.1.1] and Genmask[0.0.0.0] Success
TIME:%v RESULT:Success SUMMARY:NETWORK DIAG: Convert PING destination IP address[%s] Success
TIME:%v RESULT:Success SUMMARY:NETWORK DIAG: seq[0] recv PING echo reply from address[%v] Success
TIME:%v RESULT:Success SUMMARY:NETWORK DIAG: destination IP address[%v] Success
ccess
dress[`

	pingFail = `TIME:%v RESULT:Success SUMMARY:NETWORK DIAG: Get WAN IP [192.168.1.76] Success
TIME:%v RESULT:Success SUMMARY:NETWORK DIAG: Get Route default Gateway[192.168.1.1] and Genmask[0.0.0.0] Success
TIME:%v RESULT:Success SUMMARY:NETWORK DIAG: Convert PING destination IP address[%v] Success
TIME:%v RESULT:Failure SUMMARY:NETWORK DIAG: seq[0] wait PING echo reply from address[%v] Failure
TIME:%v RESULT:Failure SUMMARY:NETWORK DIAG: seq[1] wait PING echo reply from address[%v] Failure
TIME:%v RESULT:Failure SUMMARY:NETWORK DIAG: seq[2] wait PING echo reply from address[%v] Failure
TIME:%v RESULT:Failure SUMMARY:NETWORK DIAG: destination IP address[%v] Failure
="1"
Re`

	diaginfomode = `%v	%v	%v	%v	%v	%v	%v	%v	%v	%v	%v	%v			%v	
`
)

//普通响应的 get的头
func getResHead(t string) string {
	var res = fmt.Sprintf(defaulCliRes, `get `+t, `000`, `get `+t+` success`, t)
	return res
}

//多行数据返回时，第2行开始，只显示CONTENT信息那部分
func getOneLineHead(t string) string {
	var res = fmt.Sprintf(defaulOneLineRes, t)
	return res
}

//没有get/set/…… 的头
func NoGetResHead(t string) string {
	var res = fmt.Sprintf(defaulCliRes, ` `+t, `000`, ` `+t+` success`, t)
	return res
}

//networkDiagnosis DiagParam
func DiagParamHead(t string, arg string) string {
	var res = fmt.Sprintf(defaulCliRes, `networkDiagnosis `+t+` : `+arg+`;`, `000`, `networkDiagnosis `+t+` success`, t)
	return res
}

//set的头
func setResHead(t string, arg string) string {
	var res = fmt.Sprintf(defaulCliRes, `set `+t+` : `+arg+`;`, `000`, `set `+t+` success`, t)
	return res
}

//add成功的头
func addSuccessHead(t string, arg string) string {
	var res = fmt.Sprintf(defaulCliRes, `add `+t+` : `+arg+`;`, `000`, `add `+t+` success`, t)
	return res
}

//add空ID的头
func addEmptyIDHead(t string, arg string) string {
	var res = fmt.Sprintf(defaulCliRes, `add `+t+` : `+arg+`;`, `201`, `AddDBParam[`+t+`] fail`, t)
	return res
}

//删除数据返回的头
func deleteHead(t string, arg string) string {
	var res = fmt.Sprintf(defaulCliRes, `delete `+t+` : `+arg+`;`, `000`, `delete `+t+` success`, t)
	return res
}

func responseHead(req, res string) string {
	if "000" == res {
		t := strings.Split(req, " ")
		if len(t) < 2 {
			return "unknown args"
		}
		return fmt.Sprintf(defaulCliRes, req, res, req, t[1])
	} else if "203" == res {
		return fmt.Sprintf(defaulClires203, req)
	}
	return "unknown args"
}

func RefreshGsm(a PADisplayGSM) {
	var temp PowerAmplifierGSMweb
	queryFirst(&temp)
	queryFirst(&a)
	In, _ := strconv.ParseFloat(a.InputPowerValue, 64)
	Out, _ := strconv.ParseFloat(a.OutputPowerValue, 64)
	Rfx, _ := strconv.ParseFloat(a.ReflexPowerValue, 64)
	if temp.PASwitch == "1" {
		if (In+change_value(0.17, 0.5, 1.4)) > gsm_on_in_max || (In+change_value(0.17, 0.5, 1.4)) < gsm_on_in_min {
			In = gsm_on_in_min + (rand.Float64() * (gsm_on_in_max - gsm_on_in_min)) //超过正常区间，在正常区间随机生成
		} else {
			In = In + change_value(0.17, 0.5, 1.4) //在上次状态下随机变化
		}

		if (Out+change_value(0.15, 0.5, 1.0)) > gsm_on_out_max || (Out+change_value(0.15, 0.5, 1.0)) < gsm_on_out_min {
			Out = gsm_on_out_min + (rand.Float64() * (gsm_on_out_max - gsm_on_out_min))
		} else {
			Out = Out + change_value(0.15, 0.5, 1.0)
		}

		if (Rfx+change_value(0.12, 0.6, 1.4)) > gsm_on_rfx_max || (Rfx+change_value(0.12, 0.6, 1.4)) < gsm_on_rfx_min {
			Rfx = gsm_on_rfx_min + (rand.Float64() * (gsm_on_rfx_max - gsm_on_rfx_min))
		} else {
			Rfx = Rfx + change_value(0.12, 0.6, 1.4)
		}
	}
	if temp.PASwitch == "0" {
		if (In+change_value(0.12, 0.4, 1.1)) > gsm_off_in_max || (In+change_value(0.12, 0.4, 1.1)) < gsm_off_in_min {
			In = gsm_off_in_min + (rand.Float64() * (gsm_off_in_max - gsm_off_in_min))
		} else {
			In = In + change_value(0.12, 0.4, 1.1)
		}

		if (Out+change_value(0.12, 0.5, 1.0)) > gsm_off_out_max || (Out+change_value(0.12, 0.5, 1.0)) < gsm_off_out_min {
			Out = gsm_off_out_min + (rand.Float64() * (gsm_off_out_max - gsm_off_out_min))
		} else {
			Out = Out + change_value(0.12, 0.5, 1.0)
		}

		if (Rfx+change_value(0.1, 0.35, 1.0)) > gsm_off_rfx_max || (Rfx+change_value(0.1, 0.35, 1.0)) < gsm_off_rfx_min {
			Rfx = gsm_off_rfx_min + (rand.Float64() * (gsm_off_rfx_max - gsm_off_rfx_min))
		} else {
			Rfx = Rfx + change_value(0.1, 0.35, 1.0)
		}
	}
	//strconv.FormatFloat(float64, 'f', -1, 64)
	in := strconv.FormatFloat(In, 'f', 1, 64)
	out := strconv.FormatFloat(Out, 'f', 1, 64)
	rfx := strconv.FormatFloat(Rfx, 'f', 1, 64)
	mysql.DBer.Model(&a).Updates(map[string]interface{}{"input_power_value": in, "output_power_value": out, "reflex_power_value": rfx})
}

func RefreshLTEtdd(a PADisplayLTE) {
	var temp PowerAmplifierLTEweb
	queryFirst(&temp)
	queryFirst(&a)
	In, _ := strconv.ParseFloat(a.InputPowerValue, 64)
	Out, _ := strconv.ParseFloat(a.OutputPowerValue, 64)
	Rfx, _ := strconv.ParseFloat(a.ReflexPowerValue, 64)
	if temp.PASwitch == "1" {
		if (In+change_value(0.8, 2.2, 3.6)) > ltdd_on_in_max || (In+change_value(0.8, 2.2, 3.6)) < ltdd_on_in_min {
			In = ltdd_on_in_min + (rand.Float64() * (ltdd_on_in_max - ltdd_on_in_min))
		} else {
			In = In + change_value(0.8, 2.2, 3.6)
		}

		if (Out+change_value(0.1, 0.3, 0.7)) > ltdd_on_out_max || (Out+change_value(0.1, 0.3, 0.7)) < ltdd_on_out_min {
			Out = ltdd_on_out_min + (rand.Float64() * (ltdd_on_out_max - ltdd_on_out_min))
		} else {
			Out = Out + change_value(0.1, 0.3, 0.7)
		}

		if (Rfx+change_value(3.2, 2.1, 0.1)) > ltdd_on_rfx_max || (Rfx+change_value(3.2, 2.1, 0.1)) < ltdd_on_rfx_min {
			Rfx = ltdd_on_rfx_min + (rand.Float64() * (ltdd_on_rfx_max - ltdd_on_rfx_min))
		} else {
			Rfx = Rfx + change_value(3.2, 2.1, 0.1)
		}
	}
	if temp.PASwitch == "0" {
		if (In+change_value(1.2, 0.8, 0.2)) > ltdd_off_in_max || (In+change_value(1.2, 0.8, 0.2)) < ltdd_off_in_min {
			In = ltdd_off_in_min + (rand.Float64() * (ltdd_off_in_max - ltdd_off_in_min))
		} else {
			In = In + change_value(1.2, 0.8, 0.2)
		}

		if (Out+change_value(0, 0, 0)) > ltdd_off_out_max || (Out+change_value(0, 0, 0)) < ltdd_off_out_min {
			Out = ltdd_off_out_min + (rand.Float64() * (ltdd_off_out_max - ltdd_off_out_min))
		} else {
			Out = Out + change_value(0, 0, 0)
		}

		if (Rfx+change_value(0, 0, 0)) > ltdd_off_rfx_max || (Rfx+change_value(0, 0, 0)) < ltdd_off_rfx_min {
			Rfx = ltdd_off_rfx_min + (rand.Float64() * (ltdd_off_rfx_max - ltdd_off_rfx_min))
		} else {
			Rfx = Rfx + change_value(0, 0, 0)
		}
	}
	in := strconv.FormatFloat(In, 'f', 1, 64)
	out := strconv.FormatFloat(Out, 'f', 1, 64)
	rfx := strconv.FormatFloat(Rfx, 'f', 1, 64)
	mysql.DBer.Model(&a).Updates(map[string]interface{}{"input_power_value": in, "output_power_value": out, "reflex_power_value": rfx})
}

func RefreshLTEfdd(a PADisplayFddLTE) {
	var temp PowerAmplifierFddLTEweb
	queryFirst(&temp)
	queryFirst(&a)
	In, _ := strconv.ParseFloat(a.InputPowerValue, 64)
	Out, _ := strconv.ParseFloat(a.OutputPowerValue, 64)
	Rfx, _ := strconv.ParseFloat(a.ReflexPowerValue, 64)
	if temp.PASwitch == "1" {
		if (In+change_value(0.3, 4.2, 12.7)) > lfdd_on_in_max || (In+change_value(0.3, 4.2, 12.7)) < lfdd_on_in_min {
			In = lfdd_on_in_min + (rand.Float64() * (lfdd_on_in_max - lfdd_on_in_min)) //超过正常区间后，在正常区间随机生成
		} else {
			In = In + change_value(0.3, 4.2, 12.7) //在上次状态下随机变化
		}
		//--------------------------------------------------------------------------------------------
		if (Out+change_value(0.1, 0.4, 1.1)) > lfdd_on_out_max || (Out+change_value(0.1, 0.4, 1.1)) < lfdd_on_out_min {
			Out = lfdd_on_out_min + (rand.Float64() * (lfdd_on_out_max - lfdd_on_out_min))
		} else {
			Out = Out + change_value(0.1, 0.4, 1.1)
		}
		//-------------------------------------------------------------------------------------------------
		//rfx参数根据IN参数分情况变化
		if In < float64(-35.4) { //当in参数较小时，rfx在（-5，-1.6）变化
			if (Rfx+change_value(0.5, 0.1, 2.1)) > float64(-1.6) || (Rfx+change_value(0.5, 0.1, 2.1)) < float64(-5.0) {
				Rfx = -5 + (rand.Float64() * 4.4) //超过正常区间后，在正常区间随机生成
			} else {
				Rfx = Rfx + change_value(0.5, 0.1, 2.1) //在上次状态下随机变化
			}

		} else { //当in参数较大时，rfx在（21,29.9）（fdd_on_rfx_min，lfdd_on_rfx_max）变化
			if (Rfx+change_value(5.1, 0.3, 1.7)) > lfdd_on_rfx_max || (Rfx+change_value(5.1, 0.3, 1.7)) < lfdd_on_rfx_min {
				Rfx = lfdd_on_rfx_min + (rand.Float64() * (lfdd_on_rfx_max - lfdd_on_rfx_min))
			} else {
				Rfx = Rfx + change_value(5.1, 0.3, 1.7)
			}
		}

	}
	if temp.PASwitch == "0" {
		if (In+change_value(3.4, 8.7, 1.2)) > lfdd_off_in_max || (In+change_value(3.4, 8.7, 1.2)) < lfdd_off_in_min {
			In = lfdd_off_in_min + (rand.Float64() * (lfdd_off_in_max - lfdd_off_in_min))
		} else {
			In = In + change_value(3.4, 8.7, 1.2)
		}

		if (Out+change_value(0, 0, 0)) > lfdd_off_out_max || (Out+change_value(0, 0, 0)) < lfdd_off_out_min {
			Out = lfdd_off_out_min + (rand.Float64() * (lfdd_off_out_max - lfdd_off_out_min))
		} else {
			Out = Out + change_value(0, 0, 0)
		}

		if (Rfx+change_value(0, 0, 0)) > lfdd_off_rfx_max || (Rfx+change_value(0, 0, 0)) < lfdd_off_rfx_min {
			Rfx = lfdd_off_rfx_min + (rand.Float64() * (lfdd_off_rfx_max - lfdd_off_rfx_min))
		} else {
			Rfx = Rfx + change_value(0, 0, 0)
		}

	}
	in := strconv.FormatFloat(In, 'f', 1, 64)
	out := strconv.FormatFloat(Out, 'f', 1, 64)
	rfx := strconv.FormatFloat(Rfx, 'f', 1, 64)
	mysql.DBer.Model(&a).Updates(map[string]interface{}{"input_power_value": in, "output_power_value": out, "reflex_power_value": rfx})
}

func GSMSwitch() {
	var temp PowerAmplifierGSMweb
	queryFirst(&temp)
	if temp.PASwitch == "0" {
		mysql.DBer.Model(&temp).Updates(map[string]interface{}{"pa_switch": "1"})
	} else {
		mysql.DBer.Model(&temp).Updates(map[string]interface{}{"pa_switch": "0"})
	}
}

func FDDwitch() {
	var temp PowerAmplifierFddLTEweb
	queryFirst(&temp)
	if temp.PASwitch == "0" {
		mysql.DBer.Model(&temp).Updates(map[string]interface{}{"pa_switch": "1"})
	} else {
		mysql.DBer.Model(&temp).Updates(map[string]interface{}{"pa_switch": "0"})
	}
}

func TDDSwitch() {
	var temp PowerAmplifierLTEweb
	queryFirst(&temp)
	if temp.PASwitch == "0" {
		mysql.DBer.Model(&temp).Updates(map[string]interface{}{"pa_switch": "1"})
	} else {
		mysql.DBer.Model(&temp).Updates(map[string]interface{}{"pa_switch": "0"})
	}
}

func MasterSwitch(a int) {
	var temp1 PowerAmplifierLTEweb
	var temp2 PowerAmplifierGSMweb
	var temp3 PowerAmplifierFddLTEweb
	if a == 0 {
		mysql.DBer.Model(&temp1).Updates(map[string]interface{}{"pa_switch": "0"})
		mysql.DBer.Model(&temp2).Updates(map[string]interface{}{"pa_switch": "0"})
		mysql.DBer.Model(&temp3).Updates(map[string]interface{}{"pa_switch": "0"})
	} else {
		mysql.DBer.Model(&temp1).Updates(map[string]interface{}{"pa_switch": "1"})
		mysql.DBer.Model(&temp2).Updates(map[string]interface{}{"pa_switch": "1"})
		mysql.DBer.Model(&temp3).Updates(map[string]interface{}{"pa_switch": "1"})
	}
}

//DlATT="11" 等号后面没有空格的情况
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

//BackIfID= "3" 等号后面多了一个空格的情况
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

//判断是否pin通
func isping(ip string) bool {
	recvBuf1 := make([]byte, 2048)
	payload := []byte{0x08, 0x00, 0x4d, 0x4b, 0x00, 0x01, 0x00, 0x10, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67,
		0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69}
	Time, _ := time.ParseDuration("3s")
	conn, err := net.DialTimeout("ip4:icmp", ip, Time)
	if err != nil {
		return false
	}
	_, err = conn.Write(payload)
	if err != nil {
		return false
	}
	conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	num, err := conn.Read(recvBuf1[0:])
	if err != nil {
		//check 80 3389 443 22 port
		Timetcp, _ := time.ParseDuration("1s")
		conn1, err := net.DialTimeout("tcp", ip+":80", Timetcp)
		if err == nil {
			defer conn1.Close()
			return true
		}

		conn2, err := net.DialTimeout("tcp", ip+":443", Timetcp)
		if err == nil {
			defer conn2.Close()
			return true
		}

		conn3, err := net.DialTimeout("tcp", ip+":3389", Timetcp)
		if err == nil {
			defer conn3.Close()
			return true
		}

		conn4, err := net.DialTimeout("tcp", ip+":22", Timetcp)
		if err == nil {
			defer conn4.Close()
			return true
		}

		return false
	}
	conn.SetReadDeadline(time.Time{})
	if string(recvBuf1[0:num]) != "" {
		return true
	}
	return false

}

func DiagnosisInfoRefresh() {
	go func() {
		tt := time.NewTicker(time.Minute * 30)
		for {
			select {
			case <-tt.C:
				var t1 PADisplayGSM
				RefreshGsm(t1)
				var t2 PADisplayLTE
				RefreshLTEtdd(t2)
				var t3 PADisplayFddLTE
				RefreshLTEfdd(t3)

				var gsmTemper PowerAmplifierGSM
				var tddTemper PowerAmplifierLTE
				var fddTemper PowerAmplifierFddLTE

				var se SerialNumberParam
				var gsm PADisplayGSM
				var tdd PADisplayLTE
				var fdd PADisplayFddLTE
				var temp1 DiagInfo
				var temp2 DiagInfo
				var temp3 DiagInfo
				var temp4 DiagInfo
				var temp5 DiagInfo
				var t = time.Now().Format("20060102150405")
				queryFirst(&gsmTemper)
				queryFirst(&tddTemper)
				queryFirst(&fddTemper)
				queryFirst(&se)
				queryFirst(&gsm)
				queryFirst(&tdd)
				queryFirst(&fdd)
				{
					temp1.Id = se.McSN
					temp1.Time = string(t)
					temp1.Type = "G/32-0,0"
					temp1.Num = "0"
					temp1.CurAlarm = "0000000000"
					temp1.HisAlarm = "0000000000"
					temp1.Ipwr = gsm.InputPowerValue
					temp1.Opwr = gsm.OutputPowerValue
					temp1.Rpwr = gsm.ReflexPowerValue
					temp1.Tem = gsmTemper.Temperature
					temp1.Syn = "-"
					temp1.UnSynCnt = "-"
					temp1.Last = "46000"
				}
				{
					temp2.Id = se.McSN
					temp2.Time = string(t)
					temp2.Type = "G/96-0,0"
					temp2.Num = "0"
					temp2.CurAlarm = "0000000000"
					temp2.HisAlarm = "0000000000"
					temp2.Ipwr = gsm.InputPowerValue
					temp2.Opwr = gsm.OutputPowerValue
					temp2.Rpwr = gsm.ReflexPowerValue
					temp2.Tem = gsmTemper.Temperature
					temp2.Syn = "-"
					temp2.UnSynCnt = "-"
					temp2.Last = "46001"
				}
				{
					temp3.Id = se.McSN
					temp3.Time = string(t)
					temp3.Type = "L/38950-503"
					temp3.Num = "0"
					temp3.CurAlarm = "0000000000"
					temp3.HisAlarm = "0000000000"
					temp3.Ipwr = fdd.InputPowerValue
					temp3.Opwr = fdd.OutputPowerValue
					temp3.Rpwr = fdd.ReflexPowerValue
					temp3.Tem = fddTemper.Temperature
					temp3.Syn = "-1"
					temp3.UnSynCnt = "0"
					temp3.Last = "46000"
				}
				{
					temp4.Id = se.McSN
					temp4.Time = string(t)
					temp4.Type = "L/575-502"
					temp4.Num = "0"
					temp4.CurAlarm = "0000000000"
					temp4.HisAlarm = "0000000000"
					temp4.Ipwr = tdd.InputPowerValue
					temp4.Opwr = tdd.OutputPowerValue
					temp4.Rpwr = tdd.ReflexPowerValue
					temp4.Tem = tddTemper.Temperature
					temp4.Syn = "-1"
					temp4.UnSynCnt = "0"
					temp4.Last = "46001"
				}
				{
					temp5.Id = se.McSN
					temp5.Time = string(t)
					temp5.Type = "L/100-501"
					temp5.Num = "0"
					temp5.CurAlarm = "0000000000"
					temp5.HisAlarm = "0000000000"
					temp5.Ipwr = tdd.InputPowerValue
					temp5.Opwr = tdd.OutputPowerValue
					temp5.Rpwr = tdd.ReflexPowerValue
					temp5.Tem = tddTemper.Temperature
					temp5.Syn = "-1"
					temp5.UnSynCnt = "0"
					temp5.Last = "46011"
				}
				mysql.DBer.Create(&temp1)
				mysql.DBer.Create(&temp2)
				mysql.DBer.Create(&temp3)
				mysql.DBer.Create(&temp4)
				mysql.DBer.Create(&temp5)
			}
		}
	}()
}

func diagInfotoString() string {
	var temp []DiagInfo
	var re string
	queryAll(&temp)
	for i := 0; i < len(temp); i++ {
		//re=re+temp[i].Id+"\t"+temp[i].Time+"\t"+temp[i].Type+"\t"+temp[i].Num+"\t"+temp[i].CurAlarm+"\t"+
		//	temp[i].HisAlarm+"\t"+temp[i].Ipwr+"\t"+temp[i].Opwr+"\t"+ temp[i].Rpwr+"\t"+temp[i].Tem+"\t"+temp[i].Syn+"\t"+temp[i].UnSynCnt+"\t\t\t"+temp[i].Last+"\n"
		re = re + fmt.Sprintf(diaginfomode, temp[i].Id, temp[i].Time, temp[i].Type, temp[i].Num, temp[i].CurAlarm,
			temp[i].HisAlarm, temp[i].Ipwr, temp[i].Opwr, temp[i].Rpwr, temp[i].Tem, temp[i].Syn, temp[i].UnSynCnt, temp[i].Last)
	}
	return re
}

func CommonFunc(sn string, msgid interface{}) (interface{}, error) {
	//获取URL中传递ID参数
	//判断ID是否为后端已经注册的设备,如果存在返回设备
	if v, exits := dev.DeviceMap[sn]; exits {
		//通过msgID获取消息实体的空对象
		temp := v.GetMsgEntity(msgid)
		//具体设备处理具体的消息
		transID, _, err := v.SendMsg(msgid, temp)
		if nil != err {
			return nil, errors.New("sendMsg error")
			//TODO 处理SendMsg 错误的情况
		}
		//获取答复
		re := v.GetAck(transID)
		if nil != re {
			return re, nil
		} else {
			//TODO 超时
			return nil, errors.New("time runs out")
		}

	} else {
		return nil, errors.New("sn not exist")
		//TODO
	}
}

func CommonFuncSet(sn string, msgid interface{}, en interface{}) (interface{}, error) {
	//获取URL中传递ID参数
	//判断ID是否为后端已经注册的设备,如果存在返回设备
	if v, exits := dev.DeviceMap[sn]; exits {
		//通过msgID获取消息实体的空对象
		//具体设备处理具体的消息
		transID, _, err := v.SendMsg(msgid, en)
		if nil != err {
			return nil, errors.New("sendMsg error")
			//TODO 处理SendMsg 错误的情况
		}
		//获取答复
		re := v.GetAck(transID)
		if nil != re {
			return re, nil
		} else {
			//TODO 超时
			return nil, errors.New("time runs out")
		}

	} else {
		return nil, errors.New("sn not exist")
		//TODO
	}
}

func GetResStringDashboard(cmd string) string {
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
		case "get PowerAmplifierGSMweb":
			var temp PowerAmplifierGSMweb
			queryFirst(&temp)
			res = res + getResHead("PowerAmplifierGSMweb") + common.StructToString(temp) + `->`
		case "get PADisplayGSM":
			var temp PADisplayGSM
			RefreshGsm(temp)
			queryFirst(&temp)
			res = res + getResHead("PADisplayGSM") + common.StructToString(temp) + `->`
		case "get PowerAmplifierLTEweb":
			var temp PowerAmplifierLTEweb
			queryFirst(&temp)
			res = res + getResHead("PowerAmplifierLTEweb") + common.StructToString(temp) + `->`
		case "get PADisplayLTE":
			var temp PADisplayLTE
			RefreshLTEtdd(temp)
			queryFirst(&temp)
			res = res + getResHead("PADisplayLTE") + common.StructToString(temp) + `->`
		case "get PowerAmplifierFddLTEweb":
			var temp PowerAmplifierFddLTEweb
			queryFirst(&temp)
			res = res + getResHead("PowerAmplifierFddLTEweb") + common.StructToString(temp) + `->`
		case "get PADisplayFddLTE":
			var temp PADisplayFddLTE
			RefreshLTEfdd(temp)
			queryFirst(&temp)
			res = res + getResHead("PADisplayFddLTE") + common.StructToString(temp) + `->`
		case "get McDiagParam":
			var temp McDiagParam
			queryFirst(&temp)
			res = res + getResHead("McDiagParam") + common.StructToString(temp) + `->`
		case "get DeviceRunningState":
			res = res + getResHead("DeviceRunningState") + common.StructToString(common.AppContext.RunStat) + common.StructToString(common.AppContext.Config.Ver) + `->`
		case "get ScanSelfOptParam":
			var temp ScanSelfOptParam
			queryFirst(&temp)
			res = res + getResHead("ScanSelfOptParam") + common.StructToString(temp) + `->`
		case "get SysFlashImsiMgmt":
			var temp SysFlashImsiMgmt
			queryFirst(&temp)
			res = res + getResHead("SysFlashImsiMgmt") + common.StructToString(temp) + `->`
		case "get GPSPosition":
			var temp GPSPosition
			queryFirst(&temp)
			res = res + getResHead("GPSPosition") + common.StructToString(temp) + `->`
		case "get PowerAmplifierGSM":
			var temp PowerAmplifierGSM
			queryFirst(&temp)
			res = res + getResHead("PowerAmplifierGSM") + common.StructToString(temp) + `->`
		case "get PowerAmplifierLTE":
			var temp PowerAmplifierLTE
			queryFirst(&temp)
			res = res + getResHead("PowerAmplifierLTE") + common.StructToString(temp) + `->`
		case "get PowerAmplifierFddLTE":
			var temp PowerAmplifierFddLTE
			queryFirst(&temp)
			res = res + getResHead("PowerAmplifierFddLTE") + common.StructToString(temp) + `->`
		case "get LTEStatus":
			var temp LTEStatus
			queryFirst(&temp)
			var sn = ""
			rows, _ := mysql.DBer.DB().Query("SELECT SN FROM host_info where name='tdd'")
			for rows.Next() {
				rows.Scan(&sn)
			}
			var temp2 *lte.SyncInfoQueryACK
			temp22, _ := CommonFunc(sn, 0xF02D)
			temp2 = temp22.(*lte.SyncInfoQueryACK)
			temp.CurrentSynMode = fmt.Sprintf("%v", temp2.SyncMode)
			temp.CurrentSynStatus = fmt.Sprintf("%v", temp2.State)

			var temp3 *lte.ServingCellInfoQueryACK
			temp33, _ := CommonFunc(sn, 0xF027)
			temp3 = temp33.(*lte.ServingCellInfoQueryACK)
			temp.PhyCellId = fmt.Sprintf("%v", temp3.PCI)
			temp.EARFCNDL = fmt.Sprintf("%v", temp3.DLEarfcn)

			res = res + getResHead("LTEStatus") + common.StructToString(temp) + `->`
		case "get UnicomFddLTEStatus":
			var temp UnicomFddLTEStatus
			queryFirst(&temp)
			var sn = ""
			rows, _ := mysql.DBer.DB().Query("SELECT SN FROM host_info where name='fdd-u'")
			for rows.Next() {
				rows.Scan(&sn)
			}
			var temp2 *lte.SyncInfoQueryACK
			temp22, _ := CommonFunc(sn, 0xF02D)
			temp2 = temp22.(*lte.SyncInfoQueryACK)
			temp.CurrentSynMode = fmt.Sprintf("%v", temp2.SyncMode)
			temp.CurrentSynStatus = fmt.Sprintf("%v", temp2.State)

			var temp3 *lte.ServingCellInfoQueryACK
			temp33, _ := CommonFunc(sn, 0xF027)
			temp3 = temp33.(*lte.ServingCellInfoQueryACK)
			temp.PhyCellId = fmt.Sprintf("%v", temp3.PCI)
			temp.EARFCNDL = fmt.Sprintf("%v", temp3.DLEarfcn)

			res = res + getResHead("UnicomFddLTEStatus") + common.StructToString(temp) + `->`
		case "get TelecomFddLTEStatus":
			var temp TelecomFddLTEStatus
			queryFirst(&temp)
			var sn = ""
			rows, _ := mysql.DBer.DB().Query("SELECT SN FROM host_info where name='fdd-t'")
			for rows.Next() {
				rows.Scan(&sn)
			}
			var temp2 *lte.SyncInfoQueryACK
			temp22, _ := CommonFunc(sn, 0xF02D)
			temp2 = temp22.(*lte.SyncInfoQueryACK)
			temp.CurrentSynMode = fmt.Sprintf("%v", temp2.SyncMode)
			temp.CurrentSynStatus = fmt.Sprintf("%v", temp2.State)

			var temp3 *lte.ServingCellInfoQueryACK
			temp33, _ := CommonFunc(sn, 0xF027)
			temp3 = temp33.(*lte.ServingCellInfoQueryACK)
			temp.PhyCellId = fmt.Sprintf("%v", temp3.PCI)
			temp.EARFCNDL = fmt.Sprintf("%v", temp3.DLEarfcn)

			res = res + getResHead("TelecomFddLTEStatus") + common.StructToString(temp) + `->`
		case "get BackIfParam":
			var temp []BackIfParam
			queryAll(&temp)
			res = res + getResHead("BackIfParam") + common.StructToString(temp[0])
			for i := 1; i < len(temp); i++ {
				res = res + getOneLineHead("BackIfParam") + common.StructToString(temp[i])
			}
			res += `->`
		case "get BackIfModeNineParam":
			var temp BackIfModeNineParam
			queryFirst(&temp)
			res = res + getResHead("BackIfModeNineParam") + common.StructToString(temp) + `->`
		case "get BackIfModeTenParam":
			var temp BackIfModeTenParam
			queryFirst(&temp)
			res = res + getResHead("BackIfModeTenParam") + common.StructToString(temp) + `->`
		case "get BackIfModeElevenParam":
			var temp BackIfModeElevenParam
			queryFirst(&temp)
			res = res + getResHead("BackIfModeElevenParam") + common.StructToString(temp) + `->`
		case "get BackIfModeTwelveParam":
			var temp BackIfModeTwelveParam
			queryFirst(&temp)
			res = res + getResHead("BackIfModeTwelveParam") + common.StructToString(temp) + `->`
		case "get SerialNumberParam":
			var temp SerialNumberParam
			queryFirst(&temp)
			res = res + getResHead("SerialNumberParam") + common.StructToString(temp) + `->`
		case "get TacOrLacParam":
			var temp TacOrLacParam
			queryFirst(&temp)
			res = res + getResHead("TacOrLacParam") + common.StructToString(temp) + `->`
		case "get CellIdParam":
			var temp CellIdParam
			queryFirst(&temp)
			res = res + getResHead("CellIdParam") + common.StructToString(temp) + `->`
		case "get WANParameter":
			var temp WANParameter
			queryFirst(&temp)
			res = res + getResHead("WANParameter") + common.StructToString(temp) + `->`
		case "get NTPParameter":
			var temp NTPParameter
			queryFirst(&temp)
			res = res + getResHead("NTPParameter") + common.StructToString(temp) + `->`
		case "get CtrlPdParam":
			var temp CtrlPdParam
			queryFirst(&temp)
			res = res + getResHead("CtrlPdParam") + common.StructToString(temp) + `->`
		case "get RebootCtrl":
			var temp RebootCtrl
			queryFirst(&temp)

			var temp2 HostInfo
			temp2 = GetHostInfo("masterControl")
			temp.RebootEnable = temp2.RebootSW
			temp.RebootTime = temp2.RebootStartTime + "-" + temp2.RebootEndTime

			res = res + getResHead("RebootCtrl") + common.StructToString(temp) + `->`
		//-------------------------------获得state 0 or 1
		case "get NetworkDiagMgmt":
			var temp NetworkDiagMgmt
			queryFirst(&temp)
			res = res + getResHead("NetworkDiagMgmt") + common.StructToString(temp) + `->`
		case "get MetaData":
			str, err := ioutil.ReadFile("./dist/conf/MetaData.txt")
			if err != nil {
				res = "get MetaData failed."
			}
			res = string(str)
		case "get AlarmDictionary", "get WebConfig", "get DeviceInfo":
			res = responseHead(v, "203")
		case "GsmPowerOnOff":
			GSMSwitch()
			res = res + NoGetResHead("GsmPowerOnOff") + `->`
		case "LtePowerOnOff":
			TDDSwitch()
			res = res + NoGetResHead("LtePowerOnOff") + `->`
		case "FddLtePowerOnOff":
			FDDwitch()
			res = res + NoGetResHead("FddLtePowerOnOff") + `->`
		case "masterSwitchOff":
			MasterSwitch(0)
			res = res + NoGetResHead("masterSwitchOff") + `->`
		case "masterSwitchOn":
			MasterSwitch(1)
			res = res + NoGetResHead("masterSwitchOn") + `->`
		case "set PowerAmplifierGSM":
			var temp PowerAmplifierGSM
			re := SetFirst(args, &temp)
			res = res + setResHead("PowerAmplifierGSM", args) + re + `->`
		case "set PowerAmplifierLTE":
			var temp PowerAmplifierLTE
			re := SetFirst(args, &temp)
			res = res + setResHead("PowerAmplifierLTE", args) + re + `->`
		case "set PowerAmplifierFddLTE":
			var temp PowerAmplifierFddLTE
			re := SetFirst(args, &temp)
			res = res + setResHead("PowerAmplifierFddLTE", args) + re + `->`
		case "set PowerAmplifierGSMweb":
			var temp PowerAmplifierGSMweb
			re := SetFirst(args, &temp)
			res = res + setResHead("PowerAmplifierGSMweb", args) + re + `->`
		case "set PowerAmplifierLTEweb":
			var temp PowerAmplifierLTEweb
			re := SetFirst(args, &temp)
			res = res + setResHead("PowerAmplifierLTEweb", args) + re + `->`
		case "set PowerAmplifierFddLTEweb":
			var temp PowerAmplifierFddLTEweb
			re := SetFirst(args, &temp)
			res = res + setResHead("PowerAmplifierFddLTEweb", args) + re + `->`
		case "set BackIfParam":
			var temp BackIfParam
			re := SetSpecified(args, &temp)
			res = res + setResHead("BackIfParam", args) + re + `->`
		case "set ScanSelfOptParam":
			var temp ScanSelfOptParam
			re := SetFirst(args, &temp)
			res = res + setResHead("ScanSelfOptParam", args) + re + `->`
		case "set BackIfModeNineParam":
			var temp BackIfModeNineParam
			re := SetFirst(args, &temp)
			res = res + setResHead("BackIfModeNineParam", args) + re + `->`
		case "set BackIfModeTenParam":
			var temp BackIfModeTenParam
			re := SetFirst(args, &temp)
			res = res + setResHead("BackIfModeTenParam", args) + re + `->`
		case "set BackIfModeElevenParam":
			var temp BackIfModeElevenParam
			re := SetFirst(args, &temp)
			res = res + setResHead("BackIfModeElevenParam", args) + re + `->`
		case "set BackIfModeTwelveParam":
			var temp BackIfModeTwelveParam
			re := SetFirst(args, &temp)
			res = res + setResHead("BackIfModeTwelveParam", args) + re + `->`
		case "set SerialNumberParam":
			var temp SerialNumberParam
			re := SetFirst(args, &temp)
			res = res + setResHead("SerialNumberParam", args) + re + `->`
		case "set TacOrLacParam":
			var temp TacOrLacParam
			re := SetFirst(args, &temp)
			res = res + setResHead("TacOrLacParam", args) + re + `->`
		case "set CellIdParam":
			var temp CellIdParam
			re := SetFirst(args, &temp)
			res = res + setResHead("CellIdParam", args) + re + `->`
		case "set WANParameter":
			var temp WANParameter
			re := SetFirst(args, &temp)
			res = res + setResHead("WANParameter", args) + re + `->`
		case "set NTPParameter":
			var temp NTPParameter
			re := SetFirst(args, &temp)
			res = res + setResHead("NTPParameter", args) + re + `->`
		case "set CtrlPdParam":
			var temp CtrlPdParam
			re := SetFirst(args, &temp)
			res = res + setResHead("CtrlPdParam", args) + re + `->`
		case "set RebootCtrl":
			//var temp RebootCtrl
			//re := SetFirst(args, &temp)

			m, re := ArgsToMap(args)
			if _, ok := m["RebootEnable"]; ok {
				stmt, err := mysql.DBer.DB().Prepare("UPDATE host_info SET RebootSW =? where SN=?")
				if nil != err {
					common.Log.Info("更新rebootEnable失败")
				}
				_, err = stmt.Exec(m["RebootEnable"], "masterControl")
				if nil != err {
					common.Log.Info("更新rebootEnable失败")
				}
			}
			if _, ok := m["RebootTime"]; ok {
				time := strings.Split(m["RebootTime"], "-")
				stmt, err := mysql.DBer.DB().Prepare("UPDATE host_info SET RebootStartTime =? , RebootEndTime=? where SN=?")
				if nil != err {
					common.Log.Info("更新rebootEnable失败")
				}
				_, err = stmt.Exec(time[0], time[1], "masterControl")
				if nil != err {
					common.Log.Info("更新rebootEnable失败")
				}
				mysql.DBer.DB().Exec("UPDATE host_info SET CanReboot ='1' where SN='dashBoard'" )
			}

			res = res + setResHead("RebootCtrl", args) + re + `->`
		case "add BackIfParam":
			re, flag := addBackIfParam(args)
			if flag == 1 {
				//插入成功
				res = res + addSuccessHead("BackIfParam", args) + re + `->`
			} else if flag == 0 {
				//插入失败
				res = res + addEmptyIDHead("BackIfParam", args) + re + `->`
			} else {
				//传入ID为空
				res = res + addEmptyIDHead("BackIfParam", args) + re + `->`
			}
		case "delete BackIfParam":
			re := deleteBackIfParam(args)
			res = res + deleteHead("BackIfParam", args) + re + `->`
		case "networkDiagnosis DiagParam":
			var temp DiagParam
			re := SetFirst(args, &temp) //IP写进数据库
			res = res + DiagParamHead("DiagParam", args) + re + `->`

		case "get getFlowStatistics":
			res += "\r\nCMD    :    web get getLogFile......[OK]\r\nRESULT    :    ACK 000\r\nHINT    :    \r\nCONTENT:\r\ndate = \"TIME:20000104103028\tFLOW:0\nTIME:20000104110028\tFLOW:0\nTIME:20000104113028\tFLOW:0\nTIME:20000103173028\tFLOW:0\n\" \r\n\r\n\r\n->\r\n"
		case "web getFile":
			if args == "path=/OAM/software/web/web_page/IMEI.imei" {
				//IMSI数据采集
				quertSQL := "select No,imsi,imei,time,mode from imsi_imei;"
				rows, _ := mysql.DBer.DB().Query(quertSQL)
				var imsiStr string
				for rows.Next() {
					var No int64
					var imsi, imei, time, mode string
					rows.Scan(&No, &imsi, &imei, &time, &mode)
					imsiStr += fmt.Sprintf("NO:%v IMSI:%v IMEI:%v TIME:%v MODE:%v\n", No, imsi, imei, time, mode)
				}
				res += "\r\nCMD    :    web web command......[OK]\r\nRESULT    :    ACK 000\r\nHINT    :    \r\nCONTENT:\r\ndate = \" " + imsiStr + "\" \r\n\r\n\r\n->\r\n"
			} else if args == "path=/OAM/software/web/web_page/networkDiagnosis.txt" {
				//网络诊断
				var diag DiagParam
				var flag bool
				var t = time.Now().Format("20060102150405")
				//网络诊断实际代码
				queryFirst(&diag)
				IP := diag.DestIpAddress
				if IP == "0" { //初始IP为0，不返回任何结果
					var s = "nothing"
					res += "\r\nCMD    :    web web command......[OK]\r\nRESULT    :    ACK 000\r\nHINT    :    \r\nCONTENT:\r\ndate = \" " + s + "\" \r\n\r\n\r\n->\r\n"
					break
				}
				flag = isping(IP)
				//根据诊断情况修改NetworkDiagMgmt的state
				if flag == true { //ping通
					var s = fmt.Sprintf(pingSuccess, t, t, t, IP, t, IP, t, IP)
					res += "\r\nCMD    :    web web command......[OK]\r\nRESULT    :    ACK 000\r\nHINT    :    \r\nCONTENT:\r\ndate = \" " + s + "\" \r\n\r\n\r\n->\r\n"
				} else { //ping失败
					var s = fmt.Sprintf(pingFail, t, t, t, IP, t, IP, t, IP, t, IP, t, IP)
					res += "\r\nCMD    :    web web command......[OK]\r\nRESULT    :    ACK 000\r\nHINT    :    \r\nCONTENT:\r\ndate = \" " + s + "\" \r\n\r\n\r\n->\r\n"
				}
			} else if args == `path=/OAM/software/web/web_page/DiagnosticsInfo.csv` {
				//诊断信息
				var s = diagInfotoString()
				res += "\r\nCMD    :    web web command......[OK]\r\nRESULT    :    ACK 000\r\nHINT    :    \r\nCONTENT:\r\ndate = \"" + s + "me\" \r\n\r\n\r\n->\r\n"
			}
		case "reboot":
			res += fmt.Sprintf(defaulClires000, "reboot")
			GetAllHostInfo()
			for _, v := range HostInfoMap {
				if v.Name == "tdd" || v.Name == "fdd-u" || v.Name == "fdd-t" {
					RebootLTE(v.SN)
				}
				if v.Name == "gsm-m" || v.Name == "gsm-u" {
					RebootGSM(v.SN)
				}
			}
		default:
		}
	}
	return res
}

func RebootGSM(sn string) {
	stmt, err := mysql.DBer.DB().Prepare("UPDATE host_info SET StartTime =? where SN=?")
	if nil != err {
		log.Info("更新startTime失败")
	}
	_, err = stmt.Exec(time.Now().Unix(), sn)
	if nil != err {
		log.Info("更新startTime失败")
	}
	a := strings.Split(sn, "-")
	reboot := &gsm.Reboot{
		Mode: "1",
	}
	CommonFuncSet(a[0], "Reboot", reboot)
}

func RebootLTE(sn string) {
	stmt, err := mysql.DBer.DB().Prepare("UPDATE host_info SET StartTime =? where SN=?")
	if nil != err {
		log.Info(err)
	}
	_, err = stmt.Exec(time.Now().Unix(), sn)
	if nil != err {
		log.Info(err)
	}
	cfg := &lte.RebootCfg{
		ActiveMode: 0,
	}
	CommonFuncSet(sn, 0xF00B, cfg)
}
