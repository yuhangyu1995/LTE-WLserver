package lte

import (
	"bytes"
	"dkay/fake-wlserver/common"
	"encoding/binary"
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

//BaseTimeFormatStr ...
var BaseTimeFormatStr = "20060102150405"

// UEInfo msgID 0xF005
type UEInfo struct {
	InfoType    uint32 // 0:IMSI 1:IMEI 2:both
	IMSI        [maxIMEILen]uint8
	IMEI        [maxIMEILen]uint8
	RSSI        uint8
	TMSIPresent uint8
	TMSI        [5]uint8
	Res1        [3]uint8
	Year        uint16
	Month       uint8
	Day         uint8
	Hour        uint8
	Min         uint8
	Sec         uint8
	Res2        uint8
	SeqNum      uint32
}

func dealF005(d *msgDispose) {
	msg := new(UEInfo)
	err := readBufUseLittleEndian(d.msgbuf[msgHeaderL:], msg)
	if nil != err {
		log.Println(err.Error())
		return
	}

	imsi := bytetoString(msg.IMSI[:])
	imei := bytetoString(msg.IMEI[:])
	if "" == imei {
		imei = "0"
	}
	// tmsi := bytetoString(msg.TMSI[:])
	sn := bytetoString(d.head.SN[:])
	timeStr := fmt.Sprintf("%d%02d%02d%02d%02d%02d", msg.Year, msg.Month, msg.Day, msg.Hour, msg.Min, msg.Sec)
	t, _ := time.Parse(BaseTimeFormatStr, timeStr)
	t = t.In(time.Local)
	// time := t.Unix()

	var mode string
	if v, exits := dev.DeviceMap[sn]; exits {
		dev := v.(*DeviceInfo)
		if "46000" == dev.HeartBeat.PLMN || "46002" == dev.HeartBeat.PLMN || "46007" == dev.HeartBeat.PLMN || "46008" == dev.HeartBeat.PLMN {
			app.RunStat.LteScanCnt++
		} else if "46001" == dev.HeartBeat.PLMN || "46006" == dev.HeartBeat.PLMN || "46009" == dev.HeartBeat.PLMN {
			app.RunStat.FddLteUnicomScanCnt++
		} else if "46011" == dev.HeartBeat.PLMN || "46003" == dev.HeartBeat.PLMN || "46005" == dev.HeartBeat.PLMN {
			app.RunStat.FddLteTelecomScanCnt++
		}
		mode = fmt.Sprintf("L/%d-%d", dev.HeartBeat.DlEarfcn, dev.HeartBeat.PCI)
	}

	No := atomic.AddInt64(&app.RunStat.TotalScanIMSICnt, 1)
	insertSQL := fmt.Sprintf("INSERT INTO `imsi_imei` VALUES(%d,'%s','%s','%s','%s','%s');", No, sn, imsi, imei, t.Format(BaseTimeFormatStr), mode)
	_, err = mysql.Exec(insertSQL)
	if nil != err {
		log.Println(err)
		return
	}

	// ueinfo := common.UeInfoJSON{
	// 	SN:   sn,
	// 	Imsi: imsi,
	// 	Imei: imei,
	// 	Time: time,
	// }
	// memlist.Add(ueinfo)
}

// REMInfo msgID 0xF00A
type REMInfo struct {
	CellNum  uint16 // 1~8
	Flag     uint16 //
	CellInfo []CellInfo
}

func dealF00A(d *msgDispose) {
	msg := new(REMInfo)
	buff := bytes.NewBuffer(d.msgbuf[msgHeaderL:])
	var err error
	//read cell num
	if err = binary.Read(buff, binary.LittleEndian, &msg.CellNum); nil != err {
		log.Println(err.Error())
		return
	}
	if err = binary.Read(buff, binary.LittleEndian, &msg.Flag); nil != err {
		log.Println(err.Error())
		return
	}

	for i := msg.CellNum; i > 0; i-- {
		tempInfo := new(CellInfo)
		//baseInfo
		if err = binary.Read(buff, binary.LittleEndian, &tempInfo.BaseInfo); nil != err {
			log.Println(err.Error())
			return
		}

		//intra
		if err = binary.Read(buff, binary.LittleEndian, &tempInfo.IntraFreqNum); nil != err {
			log.Println(err.Error())
			return
		}
		for j := tempInfo.IntraFreqNum; j > 0; j-- {
			intraInfo := new(IntraFreqCellInfo)
			if err = binary.Read(buff, binary.LittleEndian, intraInfo); nil != err {
				log.Println(err.Error())
				return
			}
			tempInfo.IntraFreqCellInfo = append(tempInfo.IntraFreqCellInfo, *intraInfo)
		}

		//inter
		if err = binary.Read(buff, binary.LittleEndian, &tempInfo.InterFreqNum); nil != err {
			log.Println(err.Error())
			return
		}
		for k := tempInfo.InterFreqNum; k > 0; k-- {
			interInfo := new(InterFreqCellInfo)
			if err = binary.Read(buff, binary.LittleEndian, &interInfo.DlEarfcn); nil != err {
				log.Println(err.Error())
				return
			}
			if err = binary.Read(buff, binary.LittleEndian, &interInfo.Priorty); nil != err {
				log.Println(err.Error())
				return
			}
			if err = binary.Read(buff, binary.LittleEndian, &interInfo.QoffsetFreq); nil != err {
				log.Println(err.Error())
				return
			}
			if err = binary.Read(buff, binary.LittleEndian, &interInfo.MeasBandWidth); nil != err {
				log.Println(err.Error())
				return
			}
			if err = binary.Read(buff, binary.LittleEndian, &interInfo.NeighCellNum); nil != err {
				log.Println(err.Error())
				return
			}

			for l := interInfo.NeighCellNum; l > 0; l-- {
				neighcell := new(InterNeighCellInfo)
				if err = binary.Read(buff, binary.LittleEndian, neighcell); nil != err {
					log.Println(err.Error())
					return
				}
				interInfo.NeighcellInfo = append(interInfo.NeighcellInfo, *neighcell)

			}
			tempInfo.InterFreqCellInfo = append(tempInfo.InterFreqCellInfo, *interInfo)
		}
		msg.CellInfo = append(msg.CellInfo, *tempInfo)
	}
	fmt.Println(msg)
	t := time.Now()
	tableName := fmt.Sprintf("sweep_earfcn_info_%d_%02d", t.Year(), t.Month())
	createSQL := "CREATE TABLE IF NOT EXISTS `" + tableName + "`(" +
		" `id` int NOT NULL AUTO_INCREMENT," +
		" `sn` varchar(30) NOT NULL," +
		" `info` text NOT NULL," +
		" `time` bigint NOT NULL," +
		" PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;"
	_, err = mysql.Exec(createSQL)
	if nil != err {
		log.Println(err)
		return
	}
	sn := bytetoString(d.head.SN[:])
	msgjson, _ := json.MarshalToString(msg)
	insertSQL := fmt.Sprintf("INSERT INTO `%s` VALUES(%d,'%s','%s',%d);", tableName, 0, sn, msgjson, t.Unix())
	_, err = mysql.Exec(insertSQL)
	if nil != err {
		log.Println(err)
		return
	}

}

// EnbState msgID 0xF019
type EnbState struct {
	CellStateInd uint32
}

func dealF019(d *msgDispose) {
	msg := new(EnbState)
	err := readBufUseLittleEndian(d.msgbuf[msgHeaderL:], msg)
	if nil != err {
		log.Println(err.Error())
		return
	}

}

// HeartBeat msgID 0xF010
type HeartBeat struct {
	CellState uint16
	Band      uint16
	UlEarfcn  uint32
	DlEarfcn  uint32
	PLMN      [7]uint8
	Bandwidth uint8
	PCI       uint16
	TAC       uint16
}

func dealF010(d *msgDispose) {
	nowtime := time.Now()
	sn := bytetoString(d.head.SN[:])
	var lteDev *DeviceInfo
	srcBuf := string(d.msgbuf[msgHeaderL:])

	dev.Mu.Lock()
	defer dev.Mu.Unlock()

	//if exits change,not exits add
	if v, exits := dev.DeviceMap[sn]; exits {
		lteDev = v.(*DeviceInfo)
		if lteDev.Status == Online {
			lteDev.lasttime = time.Now()

			enb := &EnbBaseInfoQuery{
				InfoType: 6,
			}
			lteDev.SendMsg(0xF02B, enb)

			if lteDev.HeartBeat.heartSrcBuf == srcBuf {
				lteDev.SendMsg(0xF011, nil)
				return
			}
		}
	}

	msg := new(HeartBeat)
	err := readBufUseLittleEndian(d.msgbuf[msgHeaderL:], msg)
	if nil != err {
		log.Println(err.Error())
		return
	}
	plmn := bytetoString(msg.PLMN[:])

	if lteDev == nil || lteDev.Status == Offline {
		lteDev = new(DeviceInfo)
		lteDev.lasttime = nowtime
		lteDev.ENBBasic.SN = sn
		lteDev.ENBBasic.IP = d.addr
		lteDev.ENBBasic.Port = common.AppContext.Config.UDPCfg.LTE
		lteDev.ENBBasic.Mode = workMode[d.head.WorkMode]
		lteDev.HeartBeat.heartSrcBuf = srcBuf
		lteDev.Status = Online
		dev.DeviceMap[sn] = lteDev

		var upName string
		if "46000" == plmn {
			upName = "tdd"
		} else if "46001" == plmn {
			upName = "fdd-u"
		} else if "46011" == plmn {
			upName = "fdd-t"
		}

		upStr := fmt.Sprintf("update host_info set SN='%v' where Name='%v';", sn, upName)
		mysql.Exec(upStr)

		lteDev.SetBoardTime()
		enb := &EnbBaseInfoQuery{}
		for a := 0; a <= 6; a++ {
			enb.InfoType = uint32(a)
			lteDev.SendMsg(0xF02B, enb)
		}
	}

	lteDev.HeartBeat.CellState = cellState[msg.CellState]
	lteDev.HeartBeat.Band = msg.Band
	lteDev.HeartBeat.UlEarfcn = msg.UlEarfcn
	lteDev.HeartBeat.DlEarfcn = msg.DlEarfcn
	lteDev.HeartBeat.PLMN = plmn
	lteDev.HeartBeat.Bandwidth = msg.Bandwidth
	lteDev.HeartBeat.PCI = msg.PCI
	lteDev.HeartBeat.TAC = msg.TAC

	lteDev.SendMsg(0xF011, nil)
}

// Alarming msgID 0xF05B
type Alarming struct {
	AlarType uint32
	AlarFlag uint32
}

//CellInfo ...
type CellInfo struct {
	BaseInfo          CellBaseInfo
	IntraFreqNum      uint32
	IntraFreqCellInfo []IntraFreqCellInfo
	InterFreqNum      uint32
	InterFreqCellInfo []InterFreqCellInfo
}

// CellBaseInfo ...
type CellBaseInfo struct {
	DLEarfcn  uint32
	PCI       uint16
	TAC       uint16
	PLMN      uint16
	TDDSA     uint16
	CellID    uint32
	Priorty   uint32
	RSRP      uint8
	RSRQ      uint8
	Bandwidth uint8
	TDDSSP    uint8
}

// IntraFreqCellInfo ...
type IntraFreqCellInfo struct {
	DLEarfcn    uint32
	PCI         uint16
	QoffsetCell uint16
}

// InterFreqCellInfo ...
type InterFreqCellInfo struct {
	DlEarfcn      uint32
	Priorty       uint8
	QoffsetFreq   uint8
	MeasBandWidth uint16
	NeighCellNum  uint32
	NeighcellInfo []InterNeighCellInfo
}

// InterNeighCellInfo ...
type InterNeighCellInfo struct {
	PCI         uint16
	QoffsetCell uint16
}
