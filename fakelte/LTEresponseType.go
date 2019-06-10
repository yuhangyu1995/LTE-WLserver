package fakelte

import (
	"dkay/fake-wlserver/common"

	jsoniter "github.com/json-iterator/go"
)

var (
	log   =common.Log
	app   = &common.AppContext
	json  = jsoniter.ConfigCompatibleWithStandardLibrary
	mysql = &common.AppContext.DB.Mysql
	dev   = &common.AppContext.RealDev
)

var rFControl = RFControl{
	"1",
	"1",
	"fdd-u",
}

var clockSynTemp = ClockSynTemp{
	"1",
	"1",
	"-1",
	"2048",
	"0xFFFFFFFF",
	"0xFFFFFFFF",
	"0xFFFFFFFF",
	"0xFFFFFFFF",
	"fdd-u",
}

var deviceStatus = DeviceStatus{
	"1",
	"1",
	"1",
	"0",
	"-1",
	"362",
	"0001-01-01T00:00:00Z",
	"2000-01-01T00:06:27",
	"0 00:06:02",
	"fdd-u",
}

var mainCtrlParam = MainCtrlParam{
	"1",
	"0",
	"fdd-u",
}

var scanSelfOptParam = ScanSelfOptParam{
	"1",
	"0",
	"fdd-u",
}

var lAN = LAN{
	"1",
	"1",
	"192.168.1.92",
	"255.255.255.0",
	"192.168.1.2",
	"192.168.1.2",
	"192.168.1.2",
	"44:00:00:00:06:34",
	"0",
	"",
	"",
	"fdd-u",
}

var imsiRouteCtrlParam = ImsiRouteCtrlParam{
	"1",
	"192.168.1.2",
	"63352",
	"FA14C0001582",
	"fdd-u",
}

var pCCH = PCCH{
	"1",
	"32",
	"2",
	"fdd-u",
}

var bCCH = BCCH{
	"1",
	"2",
	"fdd-u",
}

var clockSyn = ClockSyn{
	"1",
	"1",
	"0",
	"0",
	"2000-01-01 00:04:58",
	"0.954051",
	"fdd-u",
}

var rFConfig = RFConfig{
	"1",
	"1",
	"25",
	"25",
	"18",
	"501",
	"0",
	"0",
	"350",
	"0",
	"0",
	"1",
	"1",
	"100",
	"18100",
	"fdd-u",
}

var eUTRAN = EUTRAN{
	"1",
	"1499999",
	"1",
	"fdd-u",
}

var fAPControl = FAPControl{
	"1",
	"1",
	"fdd-u",
}

var ePC = EPC{
	"1",
	"460011,",
	"37609",
	"",
	"0",
	"fdd-u",
}

var tacIncrease = TacIncrease{
	"1",
	"1800",
	"[30000,40000]",
	"9694",
	"fdd-u",
}

var sonConfigParam = SonConfigParam{
	"1",
	"1",
	"0",
	"0",
	"0",
	"1750,1800,1850,",
	"32",
	"32",
	"0",
	"211000..214500",
	"50",
	"4812",
	"0",
	"-6",
	"54",
	"99",
	"2",
	"192.168.197.100",
	"60008",
	"1",
	"1",
	"0",
	"5",
	"0",
	"0",
	"0",
	"0",
	"fdd-u",
}

var traceFunParam = TraceFunParam{
	"1",
	"0",
	"fdd-u",
}

var phyUlMonitorOpen = PhyUlMonitorOpen{
	"1",
	"0",
	"fdd-u",
}

var remoteMgmtParam = RemoteMgmtParam{
	"1",
	"1",
	"192.168.1.2",
	"fdd-u",
}

var masterSlaveLteIndParam = MasterSlaveLteIndParam{
	"1",
	"0",
	"fdd-u",
}

var pLMNListInfo = PLMNListInfo{
	"1",
	"1",
	"46011",
	"1",
	"fdd-u",
}

var networkTransCtrlParam = NetworkTransCtrlParam{
	"1",
	"0",
	"13",
	"13",
	"fdd-u",
}

var gSMNeighborCell = GSMNeighborCell{
	"1",
	"1",
	"46000",
	"34322",
	"51",
	"12976",
	"0",
	"68",
	"7",
	"1",
	"22",
	"8",
	"0",
	"30",
	"1",
	"1",
	"",
	"1",
	"1",
	"1",
	"",
	"-3",
	"fdd-u",
}

var cDMA2000NeighborCell = CDMA2000NeighborCell{
	"1",
	"1",
	"0",
	"2",
	"26",
	"14",
	"2",
	"15",
	"0",
	"231",
	"37",
	"fdd-u",
}

var lTENeighborCell = LTENeighborCell{
	"1",
	"1",
	"46000",
	"1",
	"38850",
	"1",
	"-3",
	"0",
	"0",
	"0",
	"4528",
	"0x0FFFFFFF",
	"0x0FFFFFFF",
	"0",
	"1",
	"",
	"1",
	"0",
	"0",
	"4",
	"0",
	"51",
	"4",
	"0",
	"0",
	"1",
	"0",
	"fdd-u",
}

var targetImsiParam = TargetImsiParam{
	"1",
	"123451234512345",
	"fdd-u",
}

var packagee = Packagee{
	"1",
	"4MS_HSLF.12.28.1-R",
	"4MS_HSLF.12.28.0-R",
	"fdd-u",
}

var rebootControl = RebootControl{
	"1",
	"1",
	"02:00-05:00",
	"fdd-u",
}

type RFControl struct {
	Id         string
	RFTxStatus string
	Hostname   string
}

func (d *RFControl) TableName() string {
	return "lte_rFControl"
}

type ClockSynTemp struct {
	Id                       string
	CurrentSynMode           string
	CurrentSynStatus         string
	CurrentFreqCalibrate     string
	CurrentSyncCellFreq      string
	CurrentSyncCellPCI       string
	CurrentSyncCellBandWidth string
	CurrentSyncCellNTX       string
	Hostname                 string
}

func (d *ClockSynTemp) TableName() string {
	return "lte_clockSynTemp"
}

type DeviceStatus struct {
	Id                      string
	CellStatus              string
	ScaleStatus             string
	LastHmsConnectionStatus string
	LastTimeSynStatus       string
	UpTime                  string
	FirstUseDate            string
	CurrentLocalTime        string
	RunningTime             string
	Hostname                string
}

func (d *DeviceStatus) TableName() string {
	return "lte_deviceStatus"
}

type MainCtrlParam struct {
	Id           string
	MainCtrlType string
	Hostname     string
}

func (d *MainCtrlParam) TableName() string {
	return "lte_mainCtrlParam"
}

type ScanSelfOptParam struct {
	Id                string
	ScanSelfOptPeriod string
	Hostname          string
}

func (d *ScanSelfOptParam) TableName() string {
	return "lte_scanSelfOptParam"
}

type LAN struct {
	Id                 string
	AddressingType     string
	IPAddress          string
	SubnetMask         string
	DefaultGateway     string
	MainDNSServer      string
	AssistDNSServer    string
	MACAddress         string
	MACAddressOverride string
	PPPoEUsername      string
	PPPoEPassword      string
	Hostname           string
}

func (d *LAN) TableName() string {
	return "lte_lAN"
}

type ImsiRouteCtrlParam struct {
	Id        string
	IpAddress string
	Port      string
	CtjstId   string
	Hostname  string
}

func (d *ImsiRouteCtrlParam) TableName() string {
	return "lte_imsiRouteCtrlParam"
}

type PCCH struct {
	Id                 string
	DefaultPagingCycle string
	NB                 string
	Hostname           string
}

func (d *PCCH) TableName() string {
	return "lte_pcch"
}

type BCCH struct {
	Id                      string
	ModificationPeriodCoeff string
	Hostname                string
}

func (d *BCCH) TableName() string {
	return "lte_bcch"
}

type ClockSyn struct {
	Id                      string
	ClockSynMode            string
	FrameHeaderTimingOffset string
	FrameHeaderCalState     string
	FrameHeaderCalDate      string
	FrameHeaderRealOffset   string
	Hostname                string
}

func (d *ClockSyn) TableName() string {
	return "lte_clockSyn"
}

type RFConfig struct {
	Id                         string
	FrequencyBandIndicator     string
	DLBandwidth                string
	ULBandwidth                string
	ReferenceSignalPower       string
	PhyCellId                  string
	PSCHPowerOffset            string
	SSCHPowerOffset            string
	PBCHPowerOffset            string
	PCFICHPowerOffset          string
	PHICHPowerOffset           string
	AdditionalSpectrumEmission string
	DuplexingMode              string
	EARFCNDL                   string
	EARFCNUL                   string
	Hostname                   string
}

func (d *RFConfig) TableName() string {
	return "lte_rFConfig"
}

type EUTRAN struct {
	Id           string
	CellIdentity string
	EnbType      string
	Hostname     string
}

func (d *EUTRAN) TableName() string {
	return "lte_eUTRAN"
}

type FAPControl struct {
	Id         string
	AdminState string
	Hostname   string
}

func (d *FAPControl) TableName() string {
	return "lte_fAPControl"
}

type EPC struct {
	Id            string
	PLMNList      string
	TAC           string
	EAID          string
	NNSFSupported string
	Hostname      string
}

func (d *EPC) TableName() string {
	return "lte_ePC"
}

type TacIncrease struct {
	Id                string
	TacIncreasePeriod string
	TacIncreaseScope  string
	TacAcerNetwork    string
	Hostname          string
}

func (d *TacIncrease) TableName() string {
	return "lte_tacIncrease"
}

type SonConfigParam struct {
	Id                         string
	SonWorkMode                string
	ManualSynCellEnable        string
	DistributedARFCNEnable     string
	PowerEnable                string
	CandidateARFCNList         string
	MaxTDNeighbourCellNum      string
	MaxLTENeighbourCellNum     string
	PbchSearchNum              string
	FreqBandList               string
	SearchFreqStep             string
	PssThreshold               string
	SssThreshold               string
	CellRsSnrThreshold         string
	TempPCIList                string
	RoutinePCIList             string
	PCIReconfigRRCActUeNumThre string
	CentralizedIP              string
	CentralizedPort            string
	CentralizedID              string
	DistributedANREnable       string
	DistributedPCIEnable       string
	TempPCIValidTimer          string
	RootSeqConfigEnable        string
	PrachConfigEnable          string
	PreambleFormat             string
	PCIAlgMode                 string
	Hostname                   string
}

func (d *SonConfigParam) TableName() string {
	return "lte_sonConfigParam"
}

type TraceFunParam struct {
	Id        string
	TraceMode string
	Hostname  string
}

func (d *TraceFunParam) TableName() string {
	return "lte_traceFunParam"
}

type PhyUlMonitorOpen struct {
	Id            string
	UeMonitorOpen string
	Hostname      string
}

func (d *PhyUlMonitorOpen) TableName() string {
	return "lte_phyUlMonitorOpen"
}

type RemoteMgmtParam struct {
	Id        string
	Enable    string
	IpAddress string
	Hostname  string
}

func (d *RemoteMgmtParam) TableName() string {
	return "lte_remoteMgmtParam"
}

type MasterSlaveLteIndParam struct {
	Id                string
	MasterSlaveLteInd string
	Hostname          string
}

func (d *MasterSlaveLteIndParam) TableName() string {
	return "lte_masterSlaveLteIndParam"
}

type PLMNListInfo struct {
	Id                         string
	PLMNListInfoId             string
	PLMNID                     string
	CellReservedForOperatorUse string
	Hostname                   string
}

func (d *PLMNListInfo) TableName() string {
	return "lte_pLMNListInfo"
}

type NetworkTransCtrlParam struct {
	Id                 string
	NetworkTransEnable string
	AttachRejectCause  string
	TAURejectCause     string
	Hostname           string
}

func (d *NetworkTransCtrlParam) TableName() string {
	return "lte_networkTransCtrlParam"
}

type GSMNeighborCell struct {
	Id                        string
	GSMNeighborCellId         string
	PLMNID                    string
	LAC                       string
	BSIC                      string
	CellID                    string
	BandIndicator             string
	BCCHARFCN                 string
	CellReselectionPriority   string
	NccPermitted              string
	ThreshXHigh               string
	ThreshXLow                string
	QRxLevMin                 string
	PMaxGERAN                 string
	Choice                    string
	ExplicitListOfARFCNsNum   string
	ExplicitListOfARFCNs      string
	ArfcnSpacing              string
	NumberOfFollowingARFCNs   string
	VariableBitMapOfARFCNsNum string
	VariableBitMapOfARFCNs    string
	QOffset                   string
	Hostname                  string
}

func (d *GSMNeighborCell) TableName() string {
	return "lte_gSMNeighborCell"
}

type CDMA2000NeighborCell struct {
	Id                      string
	CDMA2000NeighborCellId  string
	BandClass               string
	CellReselectionPriority string
	ThreshXHigh             string
	ThreshXLow              string
	TReselectionCDMA2000    string
	SearchWindowSize        string
	PreRegistrationAllowed  string
	PhysCellIdCDMA2000      string
	ARFCN                   string
	Hostname                string
}

func (d *CDMA2000NeighborCell) TableName() string {
	return "lte_cDMA2000NeighborCell"
}

type LTENeighborCell struct {
	Id                        string
	LTENeighborCellId         string
	PLMNID                    string
	CellID                    string
	EUTRACarrierARFCN         string
	PhyCellID                 string
	QOffset                   string
	CIO                       string
	RSTxPower                 string
	NoHo                      string
	TAC                       string
	PhyCellIdStart            string
	PhyCellIdRange            string
	AccessMode                string
	CsgId                     string
	FemtoName                 string
	NoX2                      string
	NoRemove                  string
	HandoverFailureRatio      string
	PreambleFormat            string
	RootSequenceIndex         string
	PrachConfigIndex          string
	ZeroCorrelationZoneConfig string
	HighSpeedFlag             string
	FreqOffset                string
	SubFrameAssignment        string
	EnbType                   string
	Hostname                  string
}

func (d *LTENeighborCell) TableName() string {
	return "lte_lTENeighborCell"
}

type TargetImsiParam struct {
	Id             string
	TargetImsiList string
	Hostname       string
}

func (d *TargetImsiParam) TableName() string {
	return "lte_targetImsiParam"
}

type Packagee struct {
	Id             string
	CurrentVersion string
	BackupVersion  string
	Hostname       string
}

func (d *Packagee) TableName() string {
	return "lte_packagee"
}

type RebootControl struct {
	Id           string
	RebootEnable string
	RebootTime   string
	Hostname     string
}

func (d *RebootControl) TableName() string {
	return "lte_rebootControl"
}

var deviceInfo = DeviceInfo{
	"1",
	"InternetGatewayDevice:1.0",
	"",
	"",
	"HSPSN_LF_1",
	"FDD-LTE",
	"LTE_FDD",
	"FA14C0001582",
	"HSPSN_LF_1V2",
	"4MS_HSLF.12.28.1-R",
	"",
	"4MS_HSLF.12.28.1-R",
	"",
	"1.3",
	"fdd-u",
}

type DeviceInfo struct {
	Id                        string
	DeviceSummary             string
	Manufacturer              string
	ManufacturerOUI           string
	ModelName                 string
	Description               string
	ProductClass              string
	SerialNumber              string
	HardwareVersion           string
	SoftwareVersion           string
	AdditionalHardwareVersion string
	AdditionalSoftwareVersion string
	ProvisioningCode          string
	SpecVersion               string
	Hostname                  string
}

func (d *DeviceInfo) TableName() string {
	return "lte_deviceInfo"
}

type LTENeighborCellInUse struct {
	Id                        string
	LTENeighborCellId         string
	PLMNID                    string
	CellID                    string
	EUTRACarrierARFCN         string
	PhyCellID                 string
	QOffset                   string
	CIO                       string
	RSTxPower                 string
	NoHo                      string
	TAC                       string
	PhyCellIdStart            string
	PhyCellIdRange            string
	AccessMode                string
	CsgId                     string
	FemtoName                 string
	NoX2                      string
	NoRemove                  string
	HandoverFailureRatio      string
	PreambleFormat            string
	RootSequenceIndex         string
	PrachConfigIndex          string
	ZeroCorrelationZoneConfig string
	HighSpeedFlag             string
	FreqOffset                string
	SubFrameAssignment        string
	EnbType                   string
	Hostname                  string
}

func (d *LTENeighborCellInUse) TableName() string {
	return "lte_Neighbor_CellInUse"
}

var lteNeighborCellInUse = LTENeighborCellInUse{
	"1",
	"1",
	"46000",
	"4533",
	"38400",
	"153",
	"-3",
	"0",
	"200",
	"0",
	"23435",
	"0x0FFFFFFF",
	"0x0FFFFFFF",
	"0",
	"0",
	"",
	"1",
	"0",
	"0",
	"4",
	"0",
	"51",
	"4",
	"0",
	"0",
	"1",
	"0",
	"fdd-u",
}

func ReadDBDataLTE() {
	//mysql.DBer.CreateTable(&lteNeighborCellInUse)
	//mysql.DBer.Create(&lteNeighborCellInUse)

	//mysql.DBer.CreateTable(&deviceInfo)
	//mysql.DBer.Create(&deviceInfo)

	//mysql.DBer.CreateTable(&rFControl)
	//mysql.DBer.Create(&rFControl)
	//
	//mysql.DBer.CreateTable(&clockSynTemp)
	//mysql.DBer.Create(&clockSynTemp)
	//
	//mysql.DBer.CreateTable(&deviceStatus)
	//mysql.DBer.Create(&deviceStatus)
	//
	//mysql.DBer.CreateTable(&mainCtrlParam)
	//mysql.DBer.Create(&mainCtrlParam)
	//
	//mysql.DBer.CreateTable(&scanSelfOptParam)
	//mysql.DBer.Create(&scanSelfOptParam)
	//
	//mysql.DBer.CreateTable(&lAN)
	//mysql.DBer.Create(&lAN)
	//
	//mysql.DBer.CreateTable(&imsiRouteCtrlParam)
	//mysql.DBer.Create(&imsiRouteCtrlParam)
	//
	//mysql.DBer.CreateTable(&pCCH)
	//mysql.DBer.Create(&pCCH)
	//
	//mysql.DBer.CreateTable(&bCCH)
	//mysql.DBer.Create(&bCCH)
	//
	//mysql.DBer.CreateTable(&clockSyn)
	//mysql.DBer.Create(&clockSyn)
	//
	//mysql.DBer.CreateTable(&rFConfig)
	//mysql.DBer.Create(&rFConfig)
	//
	//mysql.DBer.CreateTable(&eUTRAN)
	//mysql.DBer.Create(&eUTRAN)
	//
	//
	//mysql.DBer.CreateTable(&fAPControl)
	//mysql.DBer.Create(&fAPControl)
	//
	//mysql.DBer.CreateTable(&ePC)
	//mysql.DBer.Create(&ePC)
	//
	//mysql.DBer.CreateTable(&tacIncrease)
	//mysql.DBer.Create(&tacIncrease)
	//
	//mysql.DBer.CreateTable(&sonConfigParam)
	//mysql.DBer.Create(&sonConfigParam)
	//
	//mysql.DBer.CreateTable(&traceFunParam)
	//mysql.DBer.Create(&traceFunParam)
	//
	//mysql.DBer.CreateTable(&phyUlMonitorOpen)
	//mysql.DBer.Create(&phyUlMonitorOpen)
	//
	//
	//mysql.DBer.CreateTable(&remoteMgmtParam)
	//mysql.DBer.Create(&remoteMgmtParam)
	//
	//mysql.DBer.CreateTable(&masterSlaveLteIndParam)
	//mysql.DBer.Create(&masterSlaveLteIndParam)
	//
	//mysql.DBer.CreateTable(&pLMNListInfo)
	//mysql.DBer.Create(&pLMNListInfo)
	//mysql.DBer.CreateTable(&networkTransCtrlParam)
	//mysql.DBer.Create(&networkTransCtrlParam)
	//
	//mysql.DBer.CreateTable(&gSMNeighborCell)
	//mysql.DBer.Create(&gSMNeighborCell)
	//
	//mysql.DBer.CreateTable(&cDMA2000NeighborCell)
	//mysql.DBer.Create(&cDMA2000NeighborCell)
	//mysql.DBer.CreateTable(&lTENeighborCell)
	//mysql.DBer.Create(&lTENeighborCell)
	//
	//mysql.DBer.CreateTable(&targetImsiParam)
	//mysql.DBer.Create(&targetImsiParam)
	//
	//mysql.DBer.CreateTable(&packagee)
	//mysql.DBer.Create(&packagee)
	//mysql.DBer.CreateTable(&rebootControl)
	//mysql.DBer.Create(&rebootControl)

}
