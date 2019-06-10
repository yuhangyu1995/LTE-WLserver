package fake

//-----------------------------------------------------DashBoard 结构体初始值
var diagParam = DiagParam{
	"",
}

var powerAmplifierLTEweb = PowerAmplifierLTEweb{
	"0",
	"1050",
}

var pADisplayLTE = PADisplayLTE{
	"1177",
	"-27.7",
	"11",
	"-6.2",
	"450",
	"-2.2",
	"0",
	"",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
}

var powerAmplifierFddLTEweb = PowerAmplifierFddLTEweb{
	"0",
	"1000",
}

var mcDiagParam = McDiagParam{
	"W500-1807115",
	"1",
	"1800",
	"180",
	"",
	"",
	"",
}

var scanSelfOptParam = ScanSelfOptParam{
	"60",
	"1",
	"5",
}

var sysFlashImsiMgmt = SysFlashImsiMgmt{
	"0",
	"3",
	"5",
}

var gPSPosition = GPSPosition{
	"0",
	"",
	"",
}

//---------------------------------main页面结构体初值

var powerAmplifierGSM = PowerAmplifierGSM{
	"sy",
	"V5001.F.23",
	"0",
	"1",
	"118161320",
	"1245150",
	"30736448",
	"0",
	"33",
	"1100",
	"0",
	"50",
	"480",
	"380",
	"90",
	"0",
}

var powerAmplifierGSMweb = PowerAmplifierGSMweb{
	"0",
}

var pADisplayGSM = PADisplayGSM{
	"1802",
	"-8.8",
	"18",
	"-3.4",
	"469",
	"6.4",
	"0",
	"",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
}

var powerAmplifierLTE = PowerAmplifierLTE{
	"RA-1900",
	"V5001.0.22",
	"0",
	"1",
	"83033845",
	"786370",
	"29556714",
	"0",
	"32",
	"1050",
	"2200",
	"-200",
	"460",
	"380",
	"90",
	"0",
}

var powerAmplifierFddLTE = PowerAmplifierFddLTE{
	"RA-1900",
	"V5001.0.22",
	"0",
	"1",
	"42073657",
	"720875",
	"28573646",
	"0",
	"31",
	"1000",
	"2200",
	"-200",
	"480",
	"380",
	"90",
	"0",
}

var pADisplayFddLTE = PADisplayFddLTE{
	"641",
	"-45.5",
	"10",
	"-2.1",
	"435",
	"-5.0",
	"0",
	"",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
	"0",
}

var lTEStatus = LTEStatus{
	"1",
	"38950",
	"503",
	"1",
	"-1",
	"",
}

var unicomFddLTEStatus = UnicomFddLTEStatus{
	"1",
	"575",
	"502",
	"1",
	"-1",
	"",
}

var backIfParam = BackIfParam{
	"1",
	"0",
	"1",
	"120",
	"180",
	"ftp://192.168.3.246:21",
	"fykj",
	"fykj.2018",
	"0000",
	"0",
	"",
	"",
}

var backIfModeNineParam = BackIfModeNineParam{
	"",
	"",
	"",
	"",
	"",
	"",
	"0",
	"22",
	"",
	"50",
	"���ӿƼ���ѧ����",
	"",
	"9",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
}

var backIfModeTenParam = BackIfModeTenParam{
	"",
	"",
	"",
	"",
}

var backIfModeElevenParam = BackIfModeElevenParam{
	"0",
	"",
	"",
	"",
	"103.931765",
	"30.760025",
	"",
}

var backIfModeTwelveParam = BackIfModeTwelveParam{
	"",
	"",
}

var serialNumberParam = SerialNumberParam{
	"1807115",
	"180711520",
	"180711522",
	"",
	"180711542",
	"180711543",
	"180711544",
}

var cellIdParam = CellIdParam{
	"",
	"",
	"",
	"",
	"",
	"",
}

var tacOrLacParam = TacOrLacParam{
	"",
	"",
	"",
	"",
	"",
	"",
}

var wANParameter = WANParameter{
	"DHCP",
	"",
	"255.255.255.0",
	"",
	"",
	"",
	"",
	"12:32:92:00:00:00",
	"",
	"0",
	"20",
}

var nTPParameter = NTPParameter{
	"1",
	"120.25.108.11",
	"",
	"-1",
}

var ctrlPdParam = CtrlPdParam{
	"1",
	"837",
	"132.232.15.83:51122",
	"0",
}

var networkDiagMgmt = NetworkDiagMgmt{
	"0",
}

var telecomFddLTEStatus = TelecomFddLTEStatus{
	"1",
	"100",
	"501",
	"1",
	"-1",
	"",
}

//var rebootCtrl=RebootCtrl{
//	"0",
//	"02:00-05:00",
//}

var diagInfo = DiagInfo{
	"1807115",
	"20000101080017",
	"G/32-0,0",
	"0",
	"0000000000",
	"0000000000",
	"-9.0",
	"-3.4",
	"",
	"",
	"",
	"",
	"",
}

type PowerAmplifierGSMweb struct {
	PASwitch string
}

func (d *PowerAmplifierGSMweb) TableName() string {
	return "zhukong_Power_Amplifier_GSMweb"
}

//主键？
type PADisplayGSM struct {
	InputPowerVoltage         string
	InputPowerValue           string
	OutputPowerVoltage        string
	OutputPowerValue          string
	ReflexPowerVoltage        string
	ReflexPowerValue          string
	UlOutputPowerVoltage      string
	UlOutputPowerValue        string
	GainAlarm                 string
	InputPowerOverLimitAlarm  string
	TemperatureOverLimitAlarm string
	OutputPowerOverLimitAlarm string
	ReflexPowerAlarm          string
	ElectricityAlarm          string
	PowerAmplificationAlarm   string
	GridVoltageAlarm          string
	SelfExcitationAlarm       string
	AlcAlarm                  string
}

func (d *PADisplayGSM) TableName() string {
	return "zhukong_pa_display_gsm"
}

type PowerAmplifierLTEweb struct {
	PASwitch string
	DlATT    string
}

func (d *PowerAmplifierLTEweb) TableName() string {
	return "zhukong_Power_Amplifier_LTEweb"
}

//主键？
type PADisplayLTE struct {
	InputPowerVoltage         string
	InputPowerValue           string
	OutputPowerVoltage        string
	OutputPowerValue          string
	ReflexPowerVoltage        string
	ReflexPowerValue          string
	UlOutputPowerVoltage      string
	UlOutputPowerValue        string
	GainAlarm                 string
	InputPowerOverLimitAlarm  string
	TemperatureOverLimitAlarm string
	OutputPowerOverLimitAlarm string
	ReflexPowerAlarm          string
	ElectricityAlarm          string
	PowerAmplificationAlarm   string
	GridVoltageAlarm          string
	SelfExcitationAlarm       string
	AlcAlarm                  string
}

func (d *PADisplayLTE) TableName() string {
	return "zhukong_PA_Display_LTE"
}

type PowerAmplifierFddLTEweb struct {
	PASwitch string
	DlATT    string
}

func (d *PowerAmplifierFddLTEweb) TableName() string {
	return "zhukong_Power_Amplifier_Fdd_LTEweb"
}

//主键？

type PADisplayFddLTE struct {
	InputPowerVoltage         string
	InputPowerValue           string
	OutputPowerVoltage        string
	OutputPowerValue          string
	ReflexPowerVoltage        string
	ReflexPowerValue          string
	UlOutputPowerVoltage      string
	UlOutputPowerValue        string
	GainAlarm                 string
	InputPowerOverLimitAlarm  string
	TemperatureOverLimitAlarm string
	OutputPowerOverLimitAlarm string
	ReflexPowerAlarm          string
	ElectricityAlarm          string
	PowerAmplificationAlarm   string
	GridVoltageAlarm          string
	SelfExcitationAlarm       string
	AlcAlarm                  string
}

func (d *PADisplayFddLTE) TableName() string {
	return "zhukong_PA_Display_Fdd_LTE"
}

type McDiagParam struct {
	UploadFileSn       string
	UploadImeiSwitch   string
	UploadImeiPeriod   string
	UploadImeiTimeOut  string
	UploadImeiURL      string
	UploadImeiUserName string
	UploadImeiPassword string
}

func (d *McDiagParam) TableName() string {
	return "zhukong_Mc_Diag_Param"
}

type ScanSelfOptParam struct {
	ScanSelfOptPeriod       string
	ScanSelfOptDeviceSwitch string
	ScanSelfOptDeviceInd    string
}

func (d *ScanSelfOptParam) TableName() string {
	return "zhukong_Scan_Self_Opt_Param"
}

type SysFlashImsiMgmt struct {
	Enable       string
	ImsiPassCnt  string
	ImsiGroupCnt string
}

func (d *SysFlashImsiMgmt) TableName() string {
	return "zhukong_Sys_Flash_Imsi_Mgmt"
}

type GPSPosition struct {
	Status   string
	Logitude string
	Latiude  string
}

func (d *GPSPosition) TableName() string {
	return "zhukong_GPS_Position"
}

type PowerAmplifierGSM struct {
	ProductModel         string
	SlaveSoftInfo        string
	PASwitch             string
	ProtectSwitch        string
	InputPower           string
	OutputPower          string
	ReflexPower          string
	UlOutputPower        string
	Temperature          string
	DlATT                string
	UlATT                string
	InputPowerOverLimit  string
	OutputPowerOverLimit string
	ReflexPowerOverLimit string
	TemperatureOverLimit string
	PAStatus             string
}

func (d *PowerAmplifierGSM) TableName() string {
	return "zhukong_Power_Amplifier_GSM"
}

type PowerAmplifierLTE struct {
	ProductModel         string
	SlaveSoftInfo        string
	PASwitch             string
	ProtectSwitch        string
	InputPower           string
	OutputPower          string
	ReflexPower          string
	UlOutputPower        string
	Temperature          string
	DlATT                string
	UlATT                string
	InputPowerOverLimit  string
	OutputPowerOverLimit string
	ReflexPowerOverLimit string
	TemperatureOverLimit string
	PAStatus             string
}

func (d *PowerAmplifierLTE) TableName() string {
	return "zhukong_Power_Amplifier_LTE"
}

type PowerAmplifierFddLTE struct {
	ProductModel         string
	SlaveSoftInfo        string
	PASwitch             string
	ProtectSwitch        string
	InputPower           string
	OutputPower          string
	ReflexPower          string
	UlOutputPower        string
	Temperature          string
	DlATT                string
	UlATT                string
	InputPowerOverLimit  string
	OutputPowerOverLimit string
	ReflexPowerOverLimit string
	TemperatureOverLimit string
	PAStatus             string
}

func (d *PowerAmplifierFddLTE) TableName() string {
	return "zhukong_Power_Amplifier_Fdd_LTE"
}

type LTEStatus struct {
	CellStatus         string
	EARFCNDL           string
	PhyCellId          string
	CurrentSynMode     string
	CurrentSynStatus   string
	CurrentSyncCellPCI string
}

func (d *LTEStatus) TableName() string {
	return "zhukong_LTE_Status"
}

type UnicomFddLTEStatus struct {
	CellStatus         string
	EARFCNDL           string
	PhyCellId          string
	CurrentSynMode     string
	CurrentSynStatus   string
	CurrentSyncCellPCI string
}

func (d *UnicomFddLTEStatus) TableName() string {
	return "zhukong_Unicom_Fdd_LTE_Status"
}

type TelecomFddLTEStatus struct {
	CellStatus         string
	EARFCNDL           string
	PhyCellId          string
	CurrentSynMode     string
	CurrentSynStatus   string
	CurrentSyncCellPCI string
}

func (d *TelecomFddLTEStatus) TableName() string {
	return "zhukong_Telecom_Fdd_LTE_Status"
}

type BackIfParam struct {
	BackIfID             string
	BackIfMode           string
	UploadImeiSwitch     string
	UploadImeiPeriod     string
	UploadImeiTimeOut    string
	UploadImeiURL        string
	UploadImeiUserName   string
	UploadImeiPassword   string
	UploadImeiAreaCode   string
	FtpCommand           string
	UploadImeiUnicomURL  string
	UploadImeiTelecomURL string
}

func (d *BackIfParam) TableName() string {
	return "zhukong_Back_If_Param"
}

type BackIfModeNineParam struct {
	PartyAName          string
	DataMFRSName        string
	OrganizationCode    string
	PlaceCode           string
	PlaceLogitude       string
	PlaceLatiude        string
	BasicInfoEnable     string
	DeviceName          string
	DevicePlace         string
	CollectionRadius    string
	SiteName            string
	SiteAddress         string
	SiteServerType      string
	SiteBusiNature      string
	ManufacturerName    string
	ManufacturerAddress string
	Contacts            string
	ContactNumber       string
	Email               string
}

func (d *BackIfModeNineParam) TableName() string {
	return "zhukong_Back_If_Mode_Nine_Param"
}

type BackIfModeTenParam struct {
	MFRSNumber    string
	BaseStationID string
	PlaceLogitude string
	PlaceLatiude  string
}

func (d *BackIfModeTenParam) TableName() string {
	return "zhukong_Back_If_Mode_Ten_Param"
}

type BackIfModeElevenParam struct {
	BasicInfoEnable         string
	UploadBasicInfoURL      string
	UploadBasicInfoUserName string
	UploadBasicInfoPassword string
	DeviceLogitude          string
	DeviceLatiude           string
	DeviceRemarks           string
}

func (d *BackIfModeElevenParam) TableName() string {
	return "zhukong_Back_If_Mode_Eleven_Param"
}

type BackIfModeTwelveParam struct {
	CityCode   string
	DeviceName string
}

func (d *BackIfModeTwelveParam) TableName() string {
	return "zhukong_Back_If_Mode_Twelve_Param"
}

type SerialNumberParam struct {
	McSN            string
	GsmMobileSN     string
	GsmUnicomSN     string
	WcdmaSN         string
	LteSN           string
	UnicomFddLteSN  string
	TelecomFddLteSN string
}

func (d *SerialNumberParam) TableName() string {
	return "zhukong_Serial_Number_Param"
}

type TacOrLacParam struct {
	GsmMobileLac     string
	GsmUnicomLac     string
	WcdmaLac         string
	LteTac           string
	UnicomFddLteTac  string
	TelecomFddLteTac string
}

func (d *TacOrLacParam) TableName() string {
	return "zhukong_Tac_Or_Lac_Param"
}

type CellIdParam struct {
	GsmMobileCellId     string
	GsmUnicomCellId     string
	WcdmaCellId         string
	LteCellId           string
	UnicomFddLteCellId  string
	TelecomFddLteCellId string
}

func (d *CellIdParam) TableName() string {
	return "zhukong_Cell_Id_Param"
}

type WANParameter struct {
	AddressType            string
	IPAddress              string
	SubnetMask             string
	DefaultGateway         string
	StaticRouteSubnetMask  string
	DNSServer              string
	AssistDNSServer        string
	MACAddress             string
	WANWorkIPAddress       string
	DhcpFailedRebootEnable string
	DhcpFailedRebootPeriod string
}

func (d *WANParameter) TableName() string {
	return "zhukong_WAN_Parameter"
}

type NTPParameter struct {
	NTPEnable          string
	NTPServerIPAddress string
	LocalTimeZone      string
	LastTimeSyncState  string
}

func (d *NTPParameter) TableName() string {
	return "zhukong_NTP_Parameter"
}

type CtrlPdParam struct {
	Enable       string
	DeviceState  string
	BackServerIp string
	LocalPort    string
}

func (d *CtrlPdParam) TableName() string {
	return "zhukong_Ctrl_Pd_Param"
}

type NetworkDiagMgmt struct {
	NetworkDiagState string
}

func (d *NetworkDiagMgmt) TableName() string {
	return "zhukong_Network_Diag_Mgmt"
}

type DiagParam struct {
	DestIpAddress string
}

func (d *DiagParam) TableName() string {
	return "zhukong_Diag_Param"
}

type DiagInfo struct {
	Id string
	Time string
	Type string
	Num string
	CurAlarm string
	HisAlarm string
	Ipwr string
	Opwr string
	Rpwr string
	Tem string
	Syn string
	UnSynCnt string
	Last string
}

func (d *DiagInfo) TableName() string {
	return "zhukong_Diag_info"
}

var rebootctrl= RebootCtrl{
	"0",
	"02:00-05:00",
}

type RebootCtrl struct {
	RebootEnable string
	RebootTime string
}

func (d *RebootCtrl) TableName() string {
	return "zhukong_Reboot_Ctrl"
}
//HostInfo ...
type HostInfo struct {
	ID    int `json:"-"`
	SN    string
	IP    string
	Port  int
	Name  string
	Alias string
	Type  string
	PLMN  string
	RebootSW string
	RebootStartTime string
	RebootEndTime string
	StartTime int64
	CanReboot string
}

var HostInfoMap = make(map[string]HostInfo)

//数据库初始化
func ReadDBData() {
	//mysql.DBer.CreateTable(&powerAmplifierLTEweb)
	//mysql.DBer.Create(&powerAmplifierLTEweb)
	//
	//mysql.DBer.CreateTable(&pADisplayLTE)
	//mysql.DBer.Create(&pADisplayLTE)
	//
	//mysql.DBer.CreateTable(&powerAmplifierFddLTEweb)
	//mysql.DBer.Create(&powerAmplifierFddLTEweb)
	//
	//mysql.DBer.CreateTable(&mcDiagParam)
	//mysql.DBer.Create(&mcDiagParam)
	//
	//mysql.DBer.CreateTable(&deviceRunningState)
	//mysql.DBer.Create(&deviceRunningState)
	//
	//mysql.DBer.CreateTable(&scanSelfOptParam)
	//mysql.DBer.Create(&scanSelfOptParam)
	//
	//mysql.DBer.CreateTable(&sysFlashImsiMgmt)
	//mysql.DBer.Create(&sysFlashImsiMgmt)
	//
	//mysql.DBer.CreateTable(&gPSPosition)
	//mysql.DBer.Create(&gPSPosition)
	//
	//mysql.DBer.CreateTable(&powerAmplifierGSM)
	//mysql.DBer.Create(&powerAmplifierGSM)
	//
	//mysql.DBer.CreateTable(&powerAmplifierGSMweb)
	//mysql.DBer.Create(&powerAmplifierGSMweb)
	//
	//mysql.DBer.CreateTable(&pADisplayGSM)
	//mysql.DBer.Create(&pADisplayGSM)
	//
	//mysql.DBer.CreateTable(&powerAmplifierLTE)
	//mysql.DBer.Create(&powerAmplifierLTE)
	//
	//mysql.DBer.CreateTable(&powerAmplifierFddLTE)
	//mysql.DBer.Create(&powerAmplifierFddLTE)
	//
	//mysql.DBer.CreateTable(&pADisplayFddLTE)
	//mysql.DBer.Create(&pADisplayFddLTE)
	//
	//mysql.DBer.CreateTable(&lTEStatus)
	//mysql.DBer.Create(&lTEStatus)
	//
	//mysql.DBer.CreateTable(&unicomFddLTEStatus)
	//mysql.DBer.Create(&unicomFddLTEStatus)
	//
	//mysql.DBer.CreateTable(&backIfParam)
	//mysql.DBer.Create(&backIfParam)
	//
	//mysql.DBer.CreateTable(&backIfModeNineParam)
	//mysql.DBer.Create(&backIfModeNineParam)
	//
	//mysql.DBer.CreateTable(&backIfModeTenParam)
	//mysql.DBer.Create(&backIfModeTenParam)
	//
	//mysql.DBer.CreateTable(&backIfModeElevenParam)
	//mysql.DBer.Create(&backIfModeElevenParam)
	//
	//mysql.DBer.CreateTable(&backIfModeTwelveParam)
	//mysql.DBer.Create(&backIfModeTwelveParam)
	//
	//mysql.DBer.CreateTable(&serialNumberParam)
	//mysql.DBer.Create(&serialNumberParam)
	//
	//mysql.DBer.CreateTable(&cellIdParam)
	//mysql.DBer.Create(&cellIdParam)
	//
	//mysql.DBer.CreateTable(&tacOrLacParam)
	//mysql.DBer.Create(&tacOrLacParam)
	//
	//mysql.DBer.CreateTable(&wANParameter)
	//mysql.DBer.Create(&wANParameter)
	//
	//mysql.DBer.CreateTable(&nTPParameter)
	//mysql.DBer.Create(&nTPParameter)
	//
	//mysql.DBer.CreateTable(&ctrlPdParam)
	//mysql.DBer.Create(&ctrlPdParam)
	//
	//mysql.DBer.CreateTable(&networkDiagMgmt)
	//mysql.DBer.Create(&networkDiagMgmt)
	//
	//mysql.DBer.CreateTable(&telecomFddLTEStatus)
	//mysql.DBer.Create(&telecomFddLTEStatus)

	//mysql.DBer.CreateTable(&diagParam)
	//mysql.DBer.Create(&diagParam)

	//mysql.DBer.CreateTable(&diagInfo)
	//mysql.DBer.Create(&diagInfo)

	//mysql.DBer.CreateTable(&rebootctrl)
	//mysql.DBer.Create(&rebootctrl)

}