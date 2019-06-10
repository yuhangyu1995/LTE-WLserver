package web

import (
	"dkay/fake-wlserver/common"
	"dkay/fake-wlserver/fake"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	log     = common.Log
	mysql   = &common.AppContext.DB.Mysql
	hostMap = HostMap{
		DeviceMap: make(map[string]HostInfo),
	}
	dev = &common.AppContext.RealDev
)

//HostMap ...
type HostMap struct {
	Mu        sync.RWMutex
	DeviceMap map[string]HostInfo
}

//responseJSON ...
type responseJSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

//Start ...
func Start() {
	gin.SetMode(gin.ReleaseMode)
	readHostInfo()
	//fakeLTE.ReadDBDataLTE()
	//fake.ReadDBData()
	hostMap.Mu.RLock()
	if 0 >= len(hostMap.DeviceMap) {
		log.Panic("please check host info")
	}
	router := routerHandler()
	for _, v := range hostMap.DeviceMap {
		server := &http.Server{
			Addr:    fmt.Sprintf("%s:%d", v.IP, v.Port),
			Handler: router,
		}
		go server.ListenAndServe()
	}
	hostMap.Mu.RUnlock()

	RebootTickerTask()

	confRouter := gin.Default()
	confRouter.Use(logger())

	confRouter.POST("/login", func(c *gin.Context) {
		var user User
		c.ShouldBind(&user)
		if "admin" == user.Name && "73acd9a5972130b75066c82595a1fae3" == user.Pwd {
			c.JSON(http.StatusOK, responseJSON{Code: 2000, Message: "success!"})
		} else {
			c.JSON(http.StatusOK, responseJSON{Code: 2001, Message: "ERROR Incorrect username or password"})
		}
	})

	confRouter.GET("/onlinedev", getOnlineDevHostInfo)
	confRouter.POST("/changehost", func(c *gin.Context) {
		var host MixInfo
		c.ShouldBind(&host)
		upstr := fmt.Sprintf("update host_info set IP='%v' where SN='%v';", host.HostInfo.IP, host.SN)
		_, err := mysql.Exec(upstr)
		if nil != err {
			c.JSON(http.StatusOK, responseJSON{Code: 2001, Message: "更新IP失败  " + err.Error()})
		} else {
			c.JSON(http.StatusOK, responseJSON{Code: 2000, Message: "更新IP成功"})
		}
	})

	confRouter.Static("/css", "dist/confDist/dist/css")
	confRouter.Static("/js", "dist/confDist/dist/js")
	confRouter.Static("/fonts", "dist/confDist/dist/fonts")

	confRouter.Any("/", func(c *gin.Context) {
		c.File("./dist/confDist/dist/index.html")
	})

	confRouter.Run(":8080")
}

//HostInfo ...
type HostInfo struct {
	ID              int `json:"-"`
	SN              string
	IP              string
	Port            int
	Name            string
	Alias           string
	Type            string
	PLMN            string
	RebootSW        string
	RebootStartTime string
	RebootEndTime   string
	StartTime       int64
	CanReboot       string
}

func readHostInfo() {
	rows, err := mysql.DBer.DB().Query("SELECT ID,SN,IP,Port,"+"Name,Alias,Type,PLMN,RebootSW,RebootStartTime,RebootEndTime,StartTime, CanReboot from host_info")
	if nil != err {
		log.Error(err)
	}
	for rows.Next() {
		var temp HostInfo
		rows.Scan(&temp.ID, &temp.SN, &temp.IP, &temp.Port, &temp.Name, &temp.Alias, &temp.Type, &temp.PLMN, &temp.RebootSW, &temp.RebootStartTime, &temp.RebootEndTime, &temp.StartTime, &temp.CanReboot)
		if 80 == temp.Port {
			if ok, _ := regexp.MatchString(`^((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))$`, temp.IP); !ok {
				continue
			}
		} else {
			if ok, _ := regexp.MatchString(`^((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\:\d+$`, temp.IP); !ok {
				continue
			}
		}
		hostMap.Mu.Lock()
		hostMap.DeviceMap[temp.IP] = temp
		hostMap.Mu.Unlock()
	}

}

func routerHandler() *gin.Engine {
	router := gin.Default()
	router.Use(logger())
	router.Use(getHostName())

	//req
	router.POST("/req/connectCliReq.asp", connectCliReq)
	router.POST("/req/generalReq.asp", generalReq)
	router.POST("/req/authReq.asp", login)

	//cgi
	//
	router.POST("/cgi-bin/oamapp.cgi", oamapp)
	router.POST("/cgi-bin/webapp.cgi", func(c *gin.Context) {
		c.String(http.StatusOK, `{"resp":"1","content":"/app/NODE/web/ui/","hint":""}`)
	})
	router.POST("/cgi-bin/login.cgi", login)

	//asp js css static
	router.Static("/css", "./dist/css")
	router.Any("/js/:name", func(c *gin.Context) {
		h, _ := c.Get("hostInfo")
		v := h.(HostInfo)
		name := c.Param("name")
		if "GSM" == v.Type {
			c.File("./dist/jsGsm/" + name)
		} else {
			c.File("./dist/js/" + name)
		}
	})
	router.Static("/img", "./dist/img")
	router.Static("/font", "./dist/font")
	router.Static("/conf", "./dist/conf")
	router.Static("/general", "./dist/general")

	router.Any("/menu/:name", func(c *gin.Context) {
		name := c.Param("name")
		if "menu.json" == name {
			c.File("./dist/menu/menu.json")
		} else {
			h, _ := c.Get("hostInfo")
			v := h.(HostInfo)
			if "dashBoard" == v.Type {
				c.File("./dist/menu/menu.xml")
			} else {
				c.File("./dist/menu/ltemenu.xml")
			}
		}
	})

	//主要页面
	router.StaticFile("/login.asp", "./dist/login.asp")
	router.StaticFile("/dashBoard.asp", "./dist/dashBoard.asp")
	router.StaticFile("/main.asp", "./dist/main.asp")
	router.StaticFile("/main.html", "./dist/main.html")

	router.Any("/", func(c *gin.Context) {
		h, _ := c.Get("hostInfo")
		v := h.(HostInfo)
		if "GSM" == v.Type {
			c.File("./dist/login.html")
		} else {
			r := fmt.Sprintf("/login.asp?r=%d", rand.Intn(100))
			c.Redirect(301, r)
		}
	})

	return router
}

func getHostName() gin.HandlerFunc {
	return func(c *gin.Context) {
		hostMap.Mu.RLock()
		if v, exit := hostMap.DeviceMap[c.Request.Host]; exit {
			dev.Mu.RLock()
			defer dev.Mu.RUnlock()
			sn := strings.Split(v.SN, "-")
			if ev, ex := dev.DeviceMap[sn[0]]; "masterControl" == v.SN || (ex && ev.IsOnline()) {
				c.Set("hostInfo", v)
			} else {
				c.String(500, "请求设备不存在或者离线，请检查")
				c.Abort()
			}
		} else {
			c.Abort()
			return
		}
		hostMap.Mu.RUnlock()
		// 处理请求
		c.Next()

	}
}

//logger 中间件，检查token
func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		//deal GIN Defaul logger time
		// 结束时间
		end := time.Now()

		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		log.Infof("| %3d | %13v | %15s | %s  %s |",
			statusCode,
			latency,
			clientIP,
			method, path,
		)
	}
}

func RebootTickerTask() {
	go func() {
		//t1 := time.NewTicker(time.Minute * 30)
		t1 := time.NewTicker(time.Second * 15)
		for {
			select {
			case <-t1.C:
				//readHostInfo()
				fake.GetAllHostInfo()
				for _, v := range fake.HostInfoMap {
					Ifreboot(v)
				}
			default:
				time.Sleep(time.Second * 1)
			}
		}
	}()
}

func Ifreboot(hostInfo fake.HostInfo) bool { //false不重启 true重启
	canreboot := hostInfo.CanReboot
	nowhour := time.Now().Hour()
	nowmin := time.Now().Minute()
	starttime := hostInfo.RebootStartTime
	timeArr := strings.Split(starttime, ":")
	starthour, _ := strconv.Atoi(timeArr[0])
	startmin, _ := strconv.Atoi(timeArr[1])
	endtime := hostInfo.RebootEndTime
	endhour, _ := strconv.Atoi(strings.Split(endtime, ":")[0])
	endmin, _ := strconv.Atoi(strings.Split(endtime, ":")[1])
	if (nowhour > endhour || (nowhour == endhour && nowmin >= endmin)) || (nowhour < starthour || (nowhour == starthour && nowmin <= startmin)) {
		mysql.DBer.DB().Exec("UPDATE host_info SET CanReboot ='1' where SN='" + hostInfo.SN + "'")
	}
	if hostInfo.RebootSW == "0" {
		return false
	} else {
		if (nowhour > starthour || (nowhour == starthour && nowmin >= startmin)) && canreboot == "1" && (nowhour < endhour || (nowhour == endhour && nowmin <= endmin)) {
			if hostInfo.Name == "tdd" || hostInfo.Name == "fdd-u" || hostInfo.Name == "fdd-t" {
				fake.RebootLTE(hostInfo.SN)
				mysql.DBer.DB().Exec("UPDATE host_info SET CanReboot ='0' where SN='" + hostInfo.SN + "'")
			}
			if hostInfo.Name == "gsm-m" || hostInfo.Name == "gsm-u" {
				fake.RebootGSM(hostInfo.SN)
				mysql.DBer.DB().Exec("UPDATE host_info SET CanReboot ='0' where SN='" + hostInfo.SN + "'")
			}
			if hostInfo.Name == "main" {
				for _, v := range fake.HostInfoMap {
					if v.Name == "tdd" || v.Name == "fdd-u" || v.Name == "fdd-t" {
						fake.RebootLTE(v.SN)
						mysql.DBer.DB().Exec("UPDATE host_info SET CanReboot ='0' where SN='" + v.SN + "'")
					}
					if v.Name == "gsm-m" || v.Name == "gsm-u" {
						fake.RebootGSM(v.SN)
						mysql.DBer.DB().Exec("UPDATE host_info SET CanReboot ='0' where SN='" + v.SN + "'")
					}
				}
				mysql.DBer.DB().Exec("UPDATE host_info SET CanReboot ='0' where SN='" + hostInfo.SN + "'")
				//系统重启
			}
			return true
		} else {
			return false
		}
	}
}
