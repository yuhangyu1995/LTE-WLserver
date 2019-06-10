package common

import (
	"bytes"
	"database/sql"

	"github.com/jinzhu/gorm"
	//_ ...
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type mysqlDB struct {
	sqlURL string
	DBer   *gorm.DB
}

func (m *mysqlDB) doInit(url string) {
	var err error
	if "" == url {
		log.Panic("empty mysqlurl")
	}
	m.DBer, err = gorm.Open("mysql", url)
	CheckError(err)
	err = m.DBer.DB().Ping()
	CheckError(err)
	//初始化用于零时存放数据的表
	m.InitTempDB()
}

func (m *mysqlDB) InitTempDB() {
	dropSQL := "DROP TABLE IF EXISTS `imsi_imei`;"
	m.Exec(dropSQL)
	createSQL := "CREATE TABLE `imsi_imei`  (" +
		" `No` bigint(0) NOT NULL," +
		" `sn` varchar(30) NOT NULL," +
		" `imsi` varchar(20) NOT NULL," +
		" `imei` varchar(20)," +
		" `time` varchar(30) NULL," +
		" `mode` varchar(30) NOT NULL" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
	_, err := m.Exec(createSQL)
	if nil != err {
		log.Panic(err)
	}
	AppContext.RunStat.GsmMobileScanCnt = 0
	AppContext.RunStat.GsmUnicomScanCnt = 0
	AppContext.RunStat.WcdmaScanCnt = 0
	AppContext.RunStat.LteScanCnt = 0
	AppContext.RunStat.FddLteTelecomScanCnt = 0
	AppContext.RunStat.FddLteUnicomScanCnt = 0
	AppContext.RunStat.TotalScanIMSICnt = 0
}

//GetUEInfo ...
func (m *mysqlDB) GetUEInfo() *bytes.Reader {
	// rows, err := m.DBer.Query("SELECT sn,imsi,imei,time from info_2018_09_20")
	// if nil != err {
	// 	fmt.Println(err.Error())
	// 	return nil
	// }

	// var re []UeInfoJSON

	// for rows.Next() {
	// 	var temp UeInfoJSON
	// 	rows.Scan(&temp.SN, &temp.Imsi, &temp.Imei, &temp.Time)

	// 	re = append(re, temp)
	// }
	// buf, err := json.Marshal(re)
	// if nil != err {
	// 	fmt.Println(err.Error())
	// 	return nil
	// }
	// return bytes.NewReader(buf)
	return nil
}

func (m *mysqlDB) Exec(sql string) (sql.Result, error) {
	return m.DBer.DB().Exec(sql)
}
