package lte

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

const (
	//Offline ...
	Offline = iota
	//Online ...
	Online
)

var transID int32 = 1

//GetTransID ...
func GetTransID() uint16 {
	atomic.AddInt32(&transID, 1)
	atomic.CompareAndSwapInt32(&transID, 0x8000, 0)
	return uint16(atomic.LoadInt32(&transID))
}

//DeviceInfo ...
type DeviceInfo struct {
	Status    uint8 //0:off-line 1:on-line
	lasttime  time.Time
	HeartBeat HeartBeatInfo
	ENBBasic  ENBBasicInfo
}

//HeartBeatInfo ...
type HeartBeatInfo struct {
	heartSrcBuf string
	CellState   string
	Band        uint16
	UlEarfcn    uint32
	DlEarfcn    uint32
	PLMN        string
	Bandwidth   uint8
	PCI         uint16
	TAC         uint16
}

//SendMsg ...
func (D *DeviceInfo) SendMsg(id, body interface{}) (interface{}, int, error) {
	msgbody := make([]byte, 32)
	if nil != body {
		buffer1 := new(bytes.Buffer)
		binary.Write(buffer1, binary.LittleEndian, body)
		msgbody = buffer1.Bytes()
	}

	msgID := uint16(id.(int))
	msgH := &MsgHead{
		FrameHeader: 0x5555AAAA,
		MsgID:       msgID,
	}
	msgH.MsgLength = msgHeaderL + uint16(len(msgbody))
	if "FDD" == D.ENBBasic.Mode {
		msgH.WorkMode = 0xFF00
	} else if "TDD" == D.ENBBasic.Mode {
		msgH.WorkMode = 0x00FF
	}
	msgH.SysCode = GetTransID()

	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, msgH)
	msg := buffer.Bytes()
	msg = append(msg, msgbody...)

	conn, err := net.Dial("udp", D.ENBBasic.IP+":"+D.ENBBasic.Port)
	defer conn.Close()
	if err != nil {
		return 0, 0, err
	}
	sendL, err := conn.Write(msg)
	return msgH.SysCode, sendL, err
}

//GetBasicInfo ...
func (D *DeviceInfo) GetBasicInfo() (string, []map[string]interface{}) {
	re := make([]map[string]interface{}, 0)
	infoMap := make(map[string]interface{})
	infoMap["IP"] = D.ENBBasic.IP
	infoMap["Mode"] = D.ENBBasic.Mode
	infoMap["SN"] = D.ENBBasic.SN
	infoMap["Band"] = fmt.Sprintf("%v", D.HeartBeat.Band)
	infoMap["Earfcn"] = fmt.Sprintf("%d/%d", D.HeartBeat.UlEarfcn, D.HeartBeat.DlEarfcn)
	infoMap["PLMN"] = D.HeartBeat.PLMN
	infoMap["PCI"] = fmt.Sprintf("%v", D.HeartBeat.PCI)
	infoMap["TAC"] = fmt.Sprintf("%v", D.HeartBeat.TAC)

	re = append(re, infoMap)
	return "LTE", re
}

//CheckState ...
func (D *DeviceInfo) CheckState() {
	nowtime := time.Now()
	if nowtime.Sub(D.lasttime) > time.Second*10 {
		D.Status = Offline
	}
}

//IsOnline ...
func (D *DeviceInfo) IsOnline() bool {
	return D.Status == Online
}

//SetBoardTime ...
func (D *DeviceInfo) SetBoardTime() {
	//fix time UTC 8 bug
	nowTime := []uint8(time.Now().UTC().Format("2006.01.02-15:04:05"))
	enbtime := new(EnodeBTimeCfg)
	for k := range nowTime {
		enbtime.Time[k] = nowTime[k]
	}
	D.SendMsg(0xF01F, enbtime)

}

//GetAck ...
func (D *DeviceInfo) GetAck(transID interface{}) interface{} {
	t1 := time.NewTicker(time.Second * 3)

	defer rec.Mu.RUnlock()
	defer t1.Stop()

	key := (transID.(uint16))
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
func (D *DeviceInfo) GetMsgEntity(msgID interface{}) interface{} {
	id := uint16(msgID.(int))
	return getMSGEntity(id)
}

var workMode = map[uint16]string{
	0xFF00: "FDD",
	0x00FF: "TDD",
}

var cellState = map[uint16]string{
	0: "IDLE",
	1: "Sweep/Synchronization",
	2: "Activating",
	3: "Activated",
	4: "Deactivation",
}

//ENBBasicInfo eNodeB Basic INfo
type ENBBasicInfo struct {
	IP             string
	Port           string
	SN             string
	Mode           string
	EquipmentModel string
	HWVersion      string
	SWVersion      string
	MAC            string
	Uboot          string
	Temperature    string
}

//CellCfgInfo Serving Cell cfg
type CellCfgInfo struct {
	ULEarfcn  uint32
	DLEArfcn  uint32
	PLMN      string
	Bandwidth uint8
	Band      uint32
	PCI       uint16
	TAC       uint16
	CellID    uint32
	UePMax    uint16
	ENBPMAX   uint16
}

//SyncInfo ...
type SyncInfo struct {
	Mode  uint16
	State uint16
}

//TddSAAndSSSP ...
type TddSAAndSSSP struct {
	SA      uint8
	SSP     uint8
	ULAlPha uint8
}
