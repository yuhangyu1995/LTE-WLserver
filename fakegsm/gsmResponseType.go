package fakegsm

//键值对，对应页面一个结构
type GsmWorkState struct {
	RunningTime string
	CellState string
	Arfcn string
	CurrentClockSyncMode string
	CurrentClockSyncState string
	SignalSwitch string
	CellId string
	TimeSyncState string
	Mcc string
	Mnc string
	Ncc string
	Bcc string
	Lac string
	Ba1List string
	Ba2List string
	ForbitLacList string
}
type GsmBasicSetting struct {
	LanIpAddress string
	OutputDataIp string
	OutputDataPort string
	Switch string
	IpAddress string
	Port string
	Type string
	Mcc string
	Mnc string
	Lac string
	CellId string
	Ncc string
	Bcc string
	LacHoppingPeriod string
	LacRange string
	Switch2 string
	Band string
	SonBand string
	InterferenceTh string
	AvailableArfcnList string
	ForbitSonArfcnList string
	SignalSwitch string
	BcchTxPowerDbm string
	Arfcn string
	MaxTxPower string
	MinTxPower string
	Period string
	Switch3 string
	IpAddress2 string
}
type RebootCtrlInfo struct {
	RebootSwitch string
	RebootTimeRangesStart string
	RebootTimeRangesEnd string
}

var gsmCarrierInfoSet=GSMCarrierInfoSet{
	"1",
	"0",
	"95",
	"1",
	"1",
	"0",
	"1",
	"gsm-u",
}
type GSMCarrierInfoSet struct {
	Id 		string
	CarrierId         string
	CarrierArfcn        string
	BcchInd             string
	TrxType        string
	EdgeCapable           string
	CarrierState          string
	HostName string
}

func (d *GSMCarrierInfoSet) TableName() string {
	return "gsm_carrier_info_set"
}
//------------------------------------------------------------------------------------
var gsmCurrentAlarmInfo=CurrentAlarmInfo{
	"1",
	"1",
	"10007",
	"fail",
	"no cell found",
	"2019-10-09",
	"gsm-u",
}

type CurrentAlarmInfo struct {
	Id              string
	InstanceId      string
	AlarmId         string
	AlarmCnName     string
	AlarmRaiseCause string
	AlarmRaiseTime  string
	HostName string
}

func (d *CurrentAlarmInfo) TableName() string {
	return "gsm_current_alarm_info"
}
//-------------------------------------------------------------------------------------
var gsmHistoryAlarmInfo=HistoryAlarmInfo{
	"1",
	"10007",
	"侦听邻区失败",
	"no ncell record",
	"1970-01-01 00:00:39",
	"system reboot",
	"1970-01-01 00:00:24",
	"gsm-u",
}

type HistoryAlarmInfo struct {
	InstanceId string
	AlarmId string
	AlarmCnName string
	AlarmRaiseCause string
	AlarmRaiseTime string
	AlarmClearCause string
	AlarmClearTime string
	HostName string
}

func (d *HistoryAlarmInfo) TableName() string {
	return "gsm_history_alarm_info"
}
//--------------------------------------------------------
func ReadDBDataGSM() {
	//mysql.DBer.CreateTable(&gsmCarrierInfoSet)
	//mysql.DBer.Create(&gsmCarrierInfoSet)

	//mysql.DBer.CreateTable(&gsmCurrentAlarmInfo)
	//mysql.DBer.Create(&gsmCurrentAlarmInfo)

	//mysql.DBer.CreateTable(&gsmHistoryAlarmInfo)
	//mysql.DBer.Create(&gsmHistoryAlarmInfo)

}