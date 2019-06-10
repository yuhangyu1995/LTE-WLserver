package gsm

import (
	"encoding/xml"
	"fmt"
	"sync/atomic"
	"time"
)

var (
	gsmCMCC = ""
	gsmCUCC = ""
)

const (
	//Offline ...
	Offline = iota
	//Online ...
	Online
)

func getMsgType(msgID string) string {
	switch msgID {
	case "GetCellPara":
		return "GetCellParaRsp"
	case "GetWeilanCfg":
		return "GetWeilanCfgRsp"
	default:
		return msgID
	}
}

func saveAckMsg(msgID string, msg interface{}) {
	rec.Mu.Lock()
	defer rec.Mu.Unlock()
	rec.revMap[msgID] = msg
}

func stringToStruct(src string, dist interface{}) {
	dealKeyValue(src, dist)
}

func dealKeyValue(src string, dist interface{}) error {
	jsonstr := stringToJSON(src)
	err := json.UnmarshalFromString(jsonstr, dist)
	return err
}

func stringToJSON(src string) string {
	flag := false
	res := `{"`
	for _, v := range src {
		if flag {
			res += `,"`
			flag = false
		} else if '[' == v {
			res += `":"`
		} else if ']' == v {
			flag = true
			res += `"`
		} else {
			res += string(v)
		}
	}
	return res + `}`
}

//HeadBeatMSG ID:101
type HeadBeatMSG struct {
	srcmsg    string
	Time      uint64 `json:"TIME,string"`
	Version   string `json:"VERSION"`
	BuildDate string `json:"BUILD_DATE"`
	Temp      int    `json:"TEMP,string"`
	GPS       bool   `json:"GPS"`
	Status    string `json:"STATUS"`
}

func dealHeadBeatMSG(msg *myMsg) {
	dev.Mu.Lock()
	defer dev.Mu.Unlock()

	var gsmdev *DeviceInfo
	nowTime := time.Now()

	if v, exits := dev.DeviceMap[msg.Head.BoardID]; exits {
		gsmdev = v.(*DeviceInfo)
		if gsmdev.Status == Online {
			gsmdev.lasttime = nowTime
			if gsmdev.HeartInfo.srcmsg == msg.srcMsg {
				return
			}
		}
	}

	if gsmdev == nil || gsmdev.Status == Offline {
		gsmdev = new(DeviceInfo)
		gsmdev.lasttime = nowTime
		gsmdev.fromAddr = msg.fromAddr
		gsmdev.BasicInfo.IP = msg.fromAddr.IP.String()
		gsmdev.BasicInfo.Name = msg.Head.WLName
		gsmdev.BasicInfo.SN = msg.Head.BoardID
		gsmdev.HeartInfo.srcmsg = msg.srcMsg
		gsmdev.Status = Online

		upStr := fmt.Sprintf("update host_info set SN='%v' where Name='%v';", msg.Head.BoardID+"-"+"0", "gsm-m")
		mysql.Exec(upStr)
		upStr = fmt.Sprintf("update host_info set SN='%v' where Name='%v';", msg.Head.BoardID+"-"+"1", "gsm-u")
		mysql.Exec(upStr)

		gsmdev.SetBoardTime()
		gsmdev.SendMsg("GetCellPara", nil)
		dev.DeviceMap[msg.Head.BoardID] = gsmdev
	}

	temp := HeadBeatMSG{}
	stringToStruct(msg.Body, &temp)
	gsmdev.HeartInfo = temp
}

//UEList ID:102
type UEList struct {
}

//OneUeInfo ID:103
type OneUeInfo struct {
	ID     string `json:"ID"`
	Time   int64  `json:"time,string"`
	TaType string `json:"taType"`
	Rsrp   int    `json:"rsrp,string"`
	UlCQI  int    `json:"ulCqi,string"`
	UlRssi int    `json:"ulRssi,string"`
	Imsi   string `json:"imsi"`
	Imei   string `json:"imei"`
}

func dealOneUeInfo(msg *myMsg) {
	temp := new(OneUeInfo)
	stringToStruct(msg.Body, temp)

	sn := msg.Head.BoardID
	tm := time.Unix(temp.Time, 0)

	if "000000000000000" == temp.Imsi {
		temp.Imsi = "0"
	}

	var mode string
	mnc := temp.Imsi[3:5]
	if "00" == mnc || "02" == mnc || "07" == mnc || "08" == mnc {
		mode = gsmCMCC
		app.RunStat.GsmMobileScanCnt++
	} else if "01" == mnc || "06" == mnc || "09" == mnc {
		mode = gsmCUCC
		app.RunStat.GsmUnicomScanCnt++
	}

	No := atomic.AddInt64(&app.RunStat.TotalScanIMSICnt, 1)
	insertSQL := fmt.Sprintf("INSERT INTO `imsi_imei` VALUES(%d,'%s','%s','%s','%s','%s');", No, sn, temp.Imsi, temp.Imei, tm.Format(BaseTimeFormatStr), mode)
	_, err := mysql.Exec(insertSQL)
	if nil != err {
		log.Println(err)
		return
	}

	// ueinfo := common.UeInfoJSON{
	// 	SN:   sn,
	// 	Imsi: temp.Imsi,
	// 	Imei: temp.Imei,
	// 	Time: temp.Time,
	// }
	// memlist.Add(ueinfo)
}

//ArfcnList ...
type ArfcnList struct {
	ItemCnt int `xml:"itemCnt,attr"`
	Arfcn   int `xml:",chardata"`
}

//Item ...
type Item struct {
	ArfcnList   ArfcnList `xml:"arfcnList>arfcn"`
	Mcc         int       `xml:"mcc,career"`
	Mnc         int       `xml:"mnc,career"`
	Lai         int       `xml:"lai,career"`
	CellID      int       `xml:"sib3CellId,career"`
	Bsic        int       `xml:"bsic,career"`
	Cro         int       `xml:"cro,career"`
	RxLevAccMin int       `xml:"rxLevAccMin,career"`
	RestlctHyst int       `xml:"reselctHyst,career"`
	NbFreq      int       `xml:"nbFreq,career"`
}

//Cell CelLPara xml item
type Cell struct {
	ItemCnt int    `xml:"itemCnt,attr"`
	Item    []Item `xml:"item"`
}

//CellPara ID:104
type CellPara struct {
	XMLName xml.Name `xml:"content" json:"-"`
	RxGain  int      `xml:"rxGain,career" `
	TxPwr   int      `xml:"txPwr,career"`
	Cell    Cell     `xml:"cell"`
}

func dealCellPara(msg *myMsg) {
	cellPara := new(CellPara)
	err := xml.Unmarshal([]byte(msg.Body), cellPara)
	if nil != err {
		fmt.Println(err)
		return
	}
	for _, cell := range cellPara.Cell.Item {
		if 0 == cell.Mnc || 2 == cell.Mnc || 7 == cell.Mnc || 8 == cell.Mnc {
			gsmCMCC = fmt.Sprintf("G/%d-0,0", cell.ArfcnList.Arfcn)
		} else if 1 == cell.Mnc || 6 == cell.Mnc || 9 == cell.Mnc {
			gsmCUCC = fmt.Sprintf("G/%d-0,0", cell.ArfcnList.Arfcn)
		}
	}

	dev.Mu.Lock()
	defer dev.Mu.Unlock()

	if v, exits := dev.DeviceMap[msg.Head.BoardID]; exits {
		gsmdev := v.(*DeviceInfo)
		gsmdev.Carrier = *cellPara
	}

	saveAckMsg(msg.Head.BoardID+msg.Head.Type, cellPara)
}

func (c *CellPara) getMsgBody() {

}

//WeiLanCfg ID:105
type WeiLanCfg struct {
	XMLName                 xml.Name `xml:"content" json:"-"`
	LocalIP                 string   `xml:"localIp" json:"-"`
	LocalPort               string   `xml:"localPort" json:"-"`
	SvrIP                   string   `xml:"svrIp" json:"-"`
	SvrPort                 string   `xml:"svrPort" json:"-"`
	DefaultGw               string   `xml:"defaultGw" json:"-"`
	ListenPortToRrm         string   `xml:"listenPortToRrm" json:"-"`
	SendPortToRrm           string   `xml:"sendPortToRrm" json:"-"`
	SntpEn                  string   `xml:"sntpEn" json:"-"`
	SntpSvrIP               string   `xml:"sntpSvrIp" json:"-"`
	HostSearial             string   `xml:"hostSearial"`
	HeartBeatTmrSecCnt      string   `xml:"heartBeatTmrSecCnt" json:"-"`
	InterRatBand            string   `xml:"interRatBand"`
	UeRejCause              string   `xml:"ueRejCause" json:"-"`
	UeRedirectionRat        string   `xml:"ueRedirectionRat"`
	RedirectionFcn          string   `xml:"redirectionFcn"`
	ImeiTry                 string   `xml:"imeiTry"`
	ActionAfterStart        string   `xml:"actionAfterStart"`
	UeInfoRptRealTime       string   `xml:"ueInfoRptRealTime"`
	RptTmrThreshInSec       string   `xml:"rptTmrThreshInSec"`
	RptCntThresh            string   `xml:"rptCntThresh"`
	TacUpdateInMin          string   `xml:"tacUpdateInMin"`
	TacMinVal               string   `xml:"tacMinVal"`
	TacMaxVal               string   `xml:"tacMaxVal"`
	UeLogEn                 string   `xml:"ueLogEn"`
	UeLogFileCnt            string   `xml:"ueLogFileCnt"`
	UeLogSaveTmrThreshInSec string   `xml:"ueLogSaveTmrThreshInSec"`
	UeLogSaveCntThresh      string   `xml:"ueLogSaveCntThresh"`
	RsrpRptEn               string   `xml:"rsrpRptEn"`
	CellRipEn               string   `xml:"cellRipEn"`
	RebootIfSyncLoss        string   `xml:"rebootIfSyncLoss"`
	RebootIfNoUeInMin       string   `xml:"rebootIfNoUeInMin"`
	DisableGpsAnt           string   `xml:"disableGpsAnt"`
	RptStateEn              string   `xml:"rptStateEn"`
	AutoResyncEn            string   `xml:"autoResyncEn"`
	AutoResyncHour          string   `xml:"autoResyncHour"`
	AutoSvsEn               string   `xml:"autoSvsEn"`
	McCtrlFlag              string   `xml:"mcCtrlFlag"`
	DuplexMode              string   `xml:"duplexMode"`
}

func dealWeiLanCfg(msg *myMsg) {
	cellPara := new(WeiLanCfg)
	err := xml.Unmarshal([]byte(msg.Body), cellPara)
	if nil != err {
		fmt.Println(err)
		return
	}
	saveAckMsg(msg.Head.BoardID+msg.Head.Type, cellPara)
}

//ErrorIndi ID:108
type ErrorIndi struct {
	Info string `json:"INFO"`
}

func dealErrorIndi(msg *myMsg) {

}

//CmdACK ID:109
type CmdACK struct {
	Result string `json:"RESULT"`
	Info   string `json:"INFO"`
}

func dealCmdAck(msg *myMsg) {
	temp := new(CmdACK)
	stringToStruct(msg.Body, temp)
	saveAckMsg(msg.Head.BoardID+msg.Head.Type, temp)
}

//StatusRPT ID:116
type StatusRPT struct {
	Mode     int   `json:"MODE,string"`
	Status   int   `json:"STATUS,string"`
	CPUTemp  int   `json:"CPUTEMP,string"`
	GPS      uint8 `json:"GPS,string"`
	SyncTime uint8 `json:"SYNCTIME,string"`
}

func dealStatusRPT(msg *myMsg) {
	dev.Mu.Lock()
	defer dev.Mu.Unlock()

	var gsmdev *DeviceInfo
	if v, exits := dev.DeviceMap[msg.Head.BoardID]; exits {
		gsmdev = v.(*DeviceInfo)
		temp := StatusRPT{}
		stringToStruct(msg.Body, &temp)
		gsmdev.StatusInfo = temp
	}

}

//Reboot ...
type Reboot struct {
	Mode string `json:"rebootMode,number"`
}
