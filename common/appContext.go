package common

import (
	"io/ioutil"
	"sync"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	//AppContext ...
	AppContext appContext
)

//appContext ...
type appContext struct {
	DB        DataInterface
	Config    JSONConfig
	RunStat   AppRunStat
	RealDev   RealDeviceMap
	FakeDev   FakeDeviceMap
	StartTime time.Time
}

//LoadConfig ...
func loadConfig(file string) {
	// log.Info("read cfg from ./config.json")
	data, err := ioutil.ReadFile(file)
	CheckError(err)

	err = json.Unmarshal(data, &AppContext.Config)
	CheckError(err)
}

//InitApp ...
func InitApp() {
	AppContext.StartTime = time.Now()
	loadConfig("./config.json")

	initLogger(&AppContext.Config.LogCfg)

	AppContext.DB.doInit(&AppContext.Config.DbCfg)

	AppContext.startTickerTask()

	AppContext.RealDev.DeviceMap = make(map[string]RealDevice)
}

func (AppContext *appContext) startTickerTask() {
	go func() {
		t1 := time.NewTicker(time.Second * 15)
		t2 := time.NewTicker(time.Hour * 24)

		for {
			select {
			case <-t1.C:
				AppContext.RealDev.Mu.Lock()
				for _, v := range AppContext.RealDev.DeviceMap {
					v.CheckState()
				}
				AppContext.RealDev.Mu.Unlock()
			case <-t2.C:
				AppContext.RealDev.Mu.Lock()
				for _, v := range AppContext.RealDev.DeviceMap {
					v.SetBoardTime()
				}
				AppContext.RealDev.Mu.Unlock()
			default:
				nowTime := time.Now()
				runtime := nowTime.Sub(AppContext.StartTime)
				AppContext.RunStat.RunningTimeDay = int(runtime.Hours()) / 24
				AppContext.RunStat.RunningTimeHour = int(runtime.Hours()) % 24
				AppContext.RunStat.RunningTimeMin = int(runtime.Minutes()) % 60
				AppContext.RunStat.RunningTimeSec = int(runtime.Seconds()) % 60
				time.Sleep(time.Second * 1)
			}
		}
	}()
}

// DataInterface ...
type DataInterface struct {
	Mysql      mysqlDB
	MemoryList memoryList
}

func (d *DataInterface) doInit(c *DBConfigure) {
	d.Mysql.doInit(c.MysqlURL)
	d.MemoryList.doInit()
}

//DBConfigure ...
type DBConfigure struct {
	MysqlURL string `json:"mysqlurl"`
}

//UDPConfigure ...
type UDPConfigure struct {
	LTE string `json:"lte"`
	GSM string `json:"gsm"`
}

// HTTPConfigure ...
type HTTPConfigure struct {
	Flag        bool   `json:"flag"`
	UploadCycle int64  `json:"uploadCycle"`
	URL         string `json:"url"`
}

// LogConfigure ...
type LogConfigure struct {
	Logdir       string `json:"logdir"`
	LogFileName  string `json:"logfilename"`
	LogLevel     string `json:"loglevel"`
	MaxSaveHour  int    `json:"maxsaveHour"`
	RotationHour int    `json:"rotationHour"`
}

//FileConfigure 用于储存imsi文件的路径
type FileConfigure struct {
	File string `json:"filePath"`
}

//JSONConfig ...
type JSONConfig struct {
	DbCfg   DBConfigure      `json:"db"`
	UDPCfg  UDPConfigure     `json:"udp_listen_port"`
	FTPCfg  FTPConfigure     `json:"ftp"`
	HTTPCfg HTTPConfigure    `json:"http"`
	LogCfg  LogConfigure     `json:"log"`
	FielCfg FileConfigure    `json:"savefile"`
	Ver     VersionConfigure `json:"version"`
}

//RealDeviceMap ...
type RealDeviceMap struct {
	Mu        sync.RWMutex
	DeviceMap map[string]RealDevice
}

//RealDevice ...
type RealDevice interface {
	SendMsg(msgID interface{}, msg interface{}) (interface{}, int, error)
	CheckState()
	IsOnline() bool
	SetBoardTime()
	GetAck(transID interface{}) interface{}
	GetMsgEntity(msgID interface{}) interface{}
	GetBasicInfo() (string, []map[string]interface{})
}

//FakeDeviceMap ...
type FakeDeviceMap struct {
	Mu        sync.RWMutex
	DeviceMap map[string]RealDevice
}

//AppRunStat ...
type AppRunStat struct {
	FlowStatistics       int
	RunningTimeDay       int
	RunningTimeHour      int
	RunningTimeMin       int
	RunningTimeSec       int
	TotalScanIMSICnt     int64
	GsmMobileScanCnt     int
	GsmUnicomScanCnt     int
	WcdmaScanCnt         int
	LteScanCnt           int
	FddLteUnicomScanCnt  int
	FddLteTelecomScanCnt int
}

//VersionConfigure ...
type VersionConfigure struct {
	McCurrentVersion            string
	McBackupVersion             string
	GsmMobileCurrentVersion     string
	GsmUnicomCurrentVersion     string
	WcdmaCurrentVersion         string
	WcdmaBackupVersion          string
	TdLteCurrentVersion         string
	TdLteBackupVersion          string
	TdLte2300MCurrentVersion    string
	TdLte2300MBackupVersion     string
	TdLte2600MCurrentVersion    string
	TdLte2600MBackupVersion     string
	FddLteUnicomCurrentVersion  string
	FddLteUnicomBackupVersion   string
	FddLteTelecomCurrentVersion string
	FddLteTelecomBackupVersion  string
}
