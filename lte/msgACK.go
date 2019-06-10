package lte

import (
	"log"
)

//CommonACK ...
type CommonACK struct {
	Result uint32 // 0:OK, >0 errno
}

func saveAckMsg(d *msgDispose, msg interface{}) {
	err := readBufUseLittleEndian(d.msgbuf[msgHeaderL:], msg)
	if nil != err {
		log.Println(err.Error())
		return
	}
	rec.Mu.Lock()
	defer rec.Mu.Unlock()
	rec.revMap[d.head.SysCode] = msg
}

//CommonACK ...
func dealCommonACK(d *msgDispose) {
	msg := new(CommonACK)
	saveAckMsg(d, msg)
}

//MsgACK ...
type MsgACK struct {
	Result uint32
}

//BlackAndWhiteListACK ...
type BlackAndWhiteListACK struct {
	Result     uint8
	IgnoreNum  uint8
	IgnoreList [10][maxIMSILen]uint8
}

// UpdateSoftACK ...
type UpdateSoftACK struct {
	VerType uint8 // 0: EnbSoft 1: uboot
	Status  uint8
	Res     [2]uint8
}

// EnbBaseInfoQueryACK msgID 0xF02C
type EnbBaseInfoQueryACK struct {
	InfoType uint32
	Info     [100]uint8
}

//DealEnbBaseInfoQueryACK ...
func DealEnbBaseInfoQueryACK(d *msgDispose) {
	msg := new(EnbBaseInfoQueryACK)
	err := readBufUseLittleEndian(d.msgbuf[msgHeaderL:], msg)
	if nil != err {
		log.Println(err.Error())
		return
	}

	dev.Mu.Lock()
	defer dev.Mu.Unlock()

	sn := bytetoString(d.head.SN[:])
	if v, exits := dev.DeviceMap[sn]; exits {
		dev := v.(*DeviceInfo)
		switch msg.InfoType {
		case 0:
			dev.ENBBasic.EquipmentModel = bytetoString(msg.Info[:])
		case 1:
			dev.ENBBasic.HWVersion = bytetoString(msg.Info[:])
		case 2:
			dev.ENBBasic.SWVersion = bytetoString(msg.Info[:])
		case 3:
			dev.ENBBasic.SN = bytetoString(msg.Info[:])
		case 4:
			dev.ENBBasic.MAC = bytetoString(msg.Info[:])
		case 5:
			dev.ENBBasic.Uboot = bytetoString(msg.Info[:])
		case 6:
			dev.ENBBasic.Temperature = bytetoString(msg.Info[:])
		default:
		}
	}
}

// ServingCellInfoQueryACK msgID 0xF028
type ServingCellInfoQueryACK ServingCellInfo

//dealServingCellInfoQueryACK ...
func dealServingCellInfoQueryACK(d *msgDispose) {
	msg := new(ServingCellInfoQueryACK)
	err := readBufUseLittleEndian(d.msgbuf[msgHeaderL:], msg)
	if nil != err {
		log.Println(err.Error())
		return
	}
	rec.Mu.Lock()
	defer rec.Mu.Unlock()
	rec.revMap[d.head.SysCode] = msg
}

// SyncInfoQueryACK msgID 0xF02E
type SyncInfoQueryACK struct {
	SyncMode uint16 // 0:Air interface 1: GPS
	State    uint16 //
}

func dealSyncInfoQueryACK(d *msgDispose) {
	msg := new(SyncInfoQueryACK)
	saveAckMsg(d, msg)
}

// CellStateQueryACK msgID 0xF030
type CellStateQueryACK struct {
	State uint32
}

// RXTxGainQueryACK msgID 0xF032
type RXTxGainQueryACK struct {
	GainFromReg   uint8
	GainFromMib   uint8
	PowerDFromReg uint8
	PowerDFromMib uint8
	AGCFlag       uint8
	SnfFromReg    uint8
	SnfFromMin    uint8
	Res           uint8
}

//dealRxTxPowerAck ...
func dealRxTxPowerAck(d *msgDispose) {
	msg := new(RXTxGainQueryACK)
	saveAckMsg(d, msg)
}

// RedirectQueryACK msgID 0xF040
type RedirectQueryACK struct {
	Flag   uint32 // 0:open 1:close
	Earfcn uint32
	Mode   uint32 // 0:4G 1:3G 2:2G
}

// SelfActiveQueryACK msgID 0xF042
type SelfActiveQueryACK struct {
	State uint32
}

//dealSelfActiveQueryACK ...
func dealSelfActiveQueryACK(d *msgDispose) {
	msg := new(SelfActiveQueryACK)
	saveAckMsg(d, msg)
}

// TddSSAndAplphaQueryACK msgID 0xF04C
type TddSSAndAplphaQueryACK struct {
	SA      uint8
	SSP     uint8
	ULAlpha uint8
}

//dealGPSPPSQueryACK ...
func dealTddSSAndAplphaQueryACK(d *msgDispose) {
	msg := new(TddSSAndAplphaQueryACK)
	saveAckMsg(d, msg)
}

// GPSInfo msgID 0xF05D
type GPSInfo struct {
	Longitude float64
	Latitude  float64
	Altitude  float64
	RateOfPro uint32
}

func dealGPSInfoQueryACK(d *msgDispose) {
	msg := new(GPSInfo)
	saveAckMsg(d, msg)
}

// RRCAccessQueryACK msgID 0xF066
type RRCAccessQueryACK struct {
	RrcReqNum uint32
	RrcCmpNum uint32
	Res       uint32
	ImsiNum   uint32
}

func dealRRCAccessQueryACK(d *msgDispose) {
	msg := new(RRCAccessQueryACK)
	saveAckMsg(d, msg)
}

// GPSPPSQueryACK msgID 0xF074
type GPSPPSQueryACK struct {
	Value int32
}

//dealGPSPPSQueryACK ...
func dealGPSPPSQueryACK(d *msgDispose) {
	msg := new(GPSPPSQueryACK)
	saveAckMsg(d, msg)
}

// UEInflACk msgID 0xF096
type UEInflACk struct {
	SeqNum uint32
}

//NtpStateACK msgID 0xF09C
type NtpStateACK struct {
	Value uint8 // 0:Not sync 1:sync Succ
}

func dealNtpStateACK(d *msgDispose) {
	msg := new(NtpStateACK)
	saveAckMsg(d, msg)
}
