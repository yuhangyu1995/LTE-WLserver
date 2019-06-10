package lte

// ServingCellCfg msgID 0xF003
// WIE: true
// reboot is saved: true
type ServingCellCfg struct {
	CfgInfo ServingCellInfo
}

// ServingCellInfo ...
type ServingCellInfo struct {
	ULEarfcn  uint32
	DLEarfcn  uint32
	PLMN      [7]uint8
	Bandwidth uint8
	Band      uint32
	PCI       uint16
	TAC       uint16
	CellID    uint32
	UEPMax    uint16
	ENBPMax   uint16
}

// CellSearchEarfcnCfg msgID 0xF009
// WIE: true
// reboot is saved: true
type CellSearchEarfcnCfg struct {
	IsWholeBand uint32 // 0:Enable 1:disable
	EarfcnNum   uint32 // 1~10
	List        [10]uint32
}

// CellSearchPortCfg msgID 0xF07D
// WIE: true
// reboot is saved: true
type CellSearchPortCfg struct {
	Mode uint32 // 0:RX 1:SNF
}

// RebootCfg msgID 0xF00B
// WIE: true
// reboot is saved: false
type RebootCfg struct {
	ActiveMode uint32 // 0:auto active 1:Non
}

// CellActive msgID 0xF00D
// WIE: true
// reboot is saved: fasle
type CellActive struct {
	Mode uint32 // 0:DeActive 1:active(TDD Non synchronization) 2:active & synchronization (only TDD)
}

// RxGainCfg msgID 0xF013
// WIE: true
// reboot is saved: optional
type RxGainCfg struct {
	RxGain uint32   // 0~127 db
	IsSave uint8    // 0:false 1:true
	Agc    uint8    // 0:rx 1:snf,only FDD
	Res    [2]uint8 //reserved
}

// PowerDereaseCfg msgID 0xF015
// WIE: true
// reboot is saved: optional
type PowerDereaseCfg struct {
	Txpower uint32   // 0x00~0xFF,step 0.25dB
	IsSave  uint8    // 0:false 1:true
	Res     [3]uint8 //reserved
}

// EnodeBIPCfg msgID oxF01B
// WIE: false
// reboot is saved: true
type EnodeBIPCfg struct {
	IP   [52]uint8 // eg:192.168.1.51#255.255.255.0#192.168.1.1
	Flag uint8     // 0:WAN 1:LAN
	Res  [3]uint8  //reserved
}

// EnodeBTimeCfg msgID oxF01F
type EnodeBTimeCfg struct {
	Time [20]uint8 // eg:192.168.1.51#255.255.255.0#192.168.1.1
}

// LMTIPCfg msgID 0xF025
// WIE: false
// reboot is saved: true
type LMTIPCfg struct {
	IP [32]uint8 // eg:192.168.1.53#3345
}

// REMModeCfg msgID 0xF023
// WIE: true
// reboot is saved: true
type REMModeCfg struct {
	Mode uint32 // 0:Air Interface 1:GPS，only TDD
}

// DelayPPS msgID 0xF029
// WIE: false
// reboot is saved: true
type DelayPPS struct {
	Value int32
}

// PowerOnCfg msgID 0xF03B
// WIE: fasle
// reboot is saved: true
type PowerOnCfg struct {
	PowerOn uint32 // 0:active 1:deactive
	Reboot  uint32 // 0:active 1:deactive
}

// TDDSFCfg msgID 0xF049
// WIE: false
// reboot is saved: true
type TDDSFCfg struct {
	SA  uint8 // 1:sa1 2:sa2
	SSP uint8 // 5: ssp5 7: ssp7
	Res [2]uint8
}

// GPSResetCfg msgID 0xF049
// WIE: true
// reboot is saved: false
type GPSResetCfg struct {
}

// SecondPLMNsCfg msgID 0xF060
// WIE: false
// reboot is saved: true
type SecondPLMNsCfg struct {
	Num  uint8 //1~5
	List [5][7]uint8
}

// CellParaSelfCfg msgID 0xF04F
// WIE: true
// reboot is saved: false
type CellParaSelfCfg struct {
	Band uint8
	Res  [3]uint8
}

// EarfcnSelfCfg msgID 0xF051
// WIE: true
// reboot is saved: true
type EarfcnSelfCfg struct {
	Mode  uint32 // 0:add 1:del
	Value uint32
}

// CellParaHotCfg msgID 0xF080
// WIE: true
// reboot is saved: false
type CellParaHotCfg struct {
	ULEarfcn uint32
	DLEarfcn uint32
	PLMN     [7]uint8
	Band     uint8
	CellID   uint32
	UePMax   uint32
}

// NTPServerIPCfg msgID 0xF075
// WIE: false
// reboot is saved: true
type NTPServerIPCfg struct {
	IP [20]uint8
}

// RebootTimerCfg msgID 0xF086
// WIE: false
// reboot is saved: true
type RebootTimerCfg struct {
	Flag uint8 // 0:off 1:on
	Res  [3]uint8
	Time [12]int8 // 23:15:15 GMT
}

// GetIMEICfg msgID 0xF08A
// WIE: true
// reboot is saved: true
type GetIMEICfg struct {
	Enable uint8 //0:off 1:on
	Res    [3]uint8
}

// SelectFREQCfg msgID 0xF082
// WIE: false
// reboot is saved: true
type SelectFREQCfg struct {
	Num  uint32 // 1~15
	List [15]PinBandRelation
}

//PinBandRelation ...
type PinBandRelation struct {
	PinValue uint8
	BandVal1 uint8
	BandVal2 uint8
	BandVal3 uint8
}

// BlackAndWhiteListCfg msgID 0xF039
// WIE: true
// reboot is saved: true
type BlackAndWhiteListCfg struct {
	ControlMode uint8 // 0:Del to List 1:Add to List
	Num         uint8 //1~10
	Property    uint8 // 0: BlackList 1: WhiteList
	IMSIList    [10][maxIMSILen]uint8
	ClearType   uint8 // 0:inaction 1:clean blackList 2: clean whiteList 3: clean black and White List
	Res         [2]uint8
}

// TAURejectCauseCfg msgID 0xF057
// WIE: true
// reboot is saved: true
type TAURejectCauseCfg struct {
	Cause uint32 // 0:cause15 1:cause12 2:cause3 3:cause13 4:cause22
}

// RedirectInfoCfg msgID 0xF017
// WIE: true
// reboot is saved: true
type RedirectInfoCfg struct {
	Flag         uint32 // 0:on 1:off
	Earfcn       uint32
	RedirectType uint32 // 0:4G 1:3G 2:2G
}

// SysWorkModeCfg msgID 0xF002
// WIE: false
// reboot is saved: true
type SysWorkModeCfg struct {
	SysMode uint32 // 0:TDD 1:FDD 2:GPIO singe
}

// EnodeBUpdateSoftCfg msgID 0xF06F
// WIE: true
// reboot is saved: false
type EnodeBUpdateSoftCfg struct {
	UpdateType    uint8 // 0: EnodeB software 1:uboot 2: both
	EnbSoftName   [102]uint8
	IsSave        uint8 // 0: no 1: yes
	EnbSoftMD5    [36]uint8
	UbootSoftName [40]uint8
	UbootMD5      [36]uint8
	isCfgFtp      uint8 // 0:no 1:yes
	FtpServer     [16]uint8
	Res           [3]uint8
	FtpPort       uint32
	FtpName       [20]uint8
	FtpPassword   [10]uint8
	FtpFilePath   [66]uint8
}

// UploadIMSICfg msgID 0xF077
// WIE: true
// reboot is saved: true
type UploadIMSICfg struct {
	UploadMode  uint8 // 0: Default,Real-time,skip error 1:Real-Time,save error 2:FTP
	IsCfgFtp    uint8 // 0: no 1: yes
	Res         [2]uint8
	ReportStep  uint32 // min
	FtpServer   [16]uint8
	FtpPort     uint32
	FtpName     [20]uint8
	FtpPass     [10]uint8
	FtpFilePath [66]uint8
	FileNameIP  [16]uint8
}

// RedirectIMSIListCfg msgID 0xF08E
// WIE: true
// reboot is saved: false
type RedirectIMSIListCfg struct {
	Flag     uint8 // 0: save current configuration 1: clean
	AddNum   uint8
	ImsiList [maxAddRedirectListNum][maxIMSILen]uint8
	Res      [2]uint8
}

// FDDGPSResyncCfg msgID 0xF090
// WIE: true
// reboot is saved: false
type FDDGPSResyncCfg struct {
}

// ULAlphaCfg msgID 0xF092
// WIE: false
// reboot is saved: true
type ULAlphaCfg struct {
	ULAlpha uint8 //（0， 40， 50， 60， 70， 80， 90， 100）-> sib2 alpha（0， 1， 2， 3， 4， 5， 6，7）
	Res     [3]uint8
}

// GPSChoiceCfg msgID 0xF097
// WIE: false
// reboot is saved: true
type GPSChoiceCfg struct {
	Flag uint8 // 0:GPS 1:Beidou
	Res  [3]uint8
}

// EnbMeasUECfg msgID 0xF006
// WIE: true
// reboot is saved: true
type EnbMeasUECfg struct {
	Mode          uint8 // 0:Continuous 1:Periodic 2:management and control 4: redirect
	RedirectMode  uint8 // 0~4
	CapturePeriod uint16
	ControlMode   uint8 // 0:blackList 1:whiteList
	Res           [3]uint8
}
