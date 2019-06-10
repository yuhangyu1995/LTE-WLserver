package gsm

import (
	"encoding/xml"
	"fmt"
	"net"
	"strconv"
	"time"
)

//DeviceInfo ...
type DeviceInfo struct {
	Status     uint8
	lasttime   time.Time
	fromAddr   *net.UDPAddr
	BasicInfo  ENBBasicInfo
	HeartInfo  HeadBeatMSG
	StatusInfo StatusRPT
	Carrier    CellPara
}

//ENBBasicInfo ...
type ENBBasicInfo struct {
	IP   string
	Name string
	SN   string
}

//SendMsg ...
func (d *DeviceInfo) SendMsg(msgID, msgbody interface{}) (interface{}, int, error) {
	msgType := msgID.(string)
	msg := msgType
	if nil != msgbody {
		body, err := dealSendMsgBody(msgType, msgbody)
		if nil != err {
			return nil, 0, err
		}
		msg = msg + " " + body
	}
	recID := d.BasicInfo.SN + getMsgType(msgType)
	sendL, err := conn.WriteToUDP([]byte(msg), d.fromAddr)
	return recID, sendL, err
}

//CheckState ...
func (d *DeviceInfo) CheckState() {
	nowtime := time.Now()
	if nowtime.Sub(d.lasttime) > time.Second*10 {
		d.Status = Offline
	}
}

//IsOnline ...
func (d *DeviceInfo) IsOnline() bool {
	return d.Status == Online
}

//SetBoardTime ...
func (d *DeviceInfo) SetBoardTime() {
	nowtime := time.Now().Unix()
	d.SendMsg("SetTime", nowtime)
}

//GetAck ...
func (d *DeviceInfo) GetAck(transID interface{}) interface{} {
	t1 := time.NewTicker(time.Second * 3)

	defer rec.Mu.RUnlock()
	defer t1.Stop()

	key := transID.(string)
	for {
		rec.Mu.RLock()
		if v, exits := rec.revMap[key]; exits {
			if nil != v {
				rec.revMap[key] = nil
				return v
			}
		}
		select {
		case <-t1.C:
			return nil
		default:
			time.Sleep(time.Millisecond * 100)
		}
		rec.Mu.RUnlock()
	}
}

//GetMsgEntity ...
func (d *DeviceInfo) GetMsgEntity(msgID interface{}) interface{} {
	switch msgID.(string) {
	case "SetCellPara":
		return new(CellPara)
	case "Reboot":
		return new(Reboot)
	default:
		return nil
	}
}

func dealSendMsgBody(msgType, msgbody interface{}) (string, error) {
	switch msgType {
	case "SetCellPara":
		re, err := xml.MarshalIndent(msgbody, "", "")
		return `<?xml version='1.0' encoding = 'UTF-8' standalone = 'no' ?>` + string(re), err
	case "SetTime":
		str := strconv.FormatInt(msgbody.(int64), 10)
		return str, nil
	case "Reboot":
		body := msgbody.(*Reboot)
		return body.Mode, nil
	default:
		return "", nil
	}
}

//GetBasicInfo ...
func (d *DeviceInfo) GetBasicInfo() (string, []map[string]interface{}) {
	re := make([]map[string]interface{}, 0)
	for i, v := range d.Carrier.Cell.Item {
		infoMap := make(map[string]interface{})

		infoMap["IP"] = d.BasicInfo.IP
		infoMap["Mode"] = "GSM"
		infoMap["SN"] = fmt.Sprintf("%s-%d", d.BasicInfo.SN, i)
		infoMap["Arfcn"] = fmt.Sprintf("%v", v.ArfcnList.Arfcn)
		infoMap["Mcc"] = v.Mcc
		infoMap["Mnc"] = v.Mnc
		infoMap["Lai"] = v.Lai
		re = append(re, infoMap)
	}

	return "LTE", re
}
