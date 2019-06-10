package web

import (
	"dkay/fake-wlserver/fake"
	"dkay/fake-wlserver/fakegsm"
	"dkay/fake-wlserver/fakelte"
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

// User 用户类
type User struct {
	ID   string `json:"userId"`
	Name string `json:"userName"`
	Pwd  string `json:"pwd"`
	Role string `json:"role"`
}

func login(c *gin.Context) {
	h, _ := c.Get("hostInfo")
	v := h.(HostInfo)
	username := c.PostForm("username")
	password := c.PostForm("password")
	sqlstr := fmt.Sprintf("select pwd from `sys_user` where name='%s' and host_name='%s';", username, v.Name)
	var pwd string
	err := mysql.DBer.DB().QueryRow(sqlstr).Scan(&pwd)
	if nil == err {
		if password == pwd {
			if v.Name == "gsm-u" || v.Name == "gsm-m" {
				c.String(http.StatusOK, `{"resp":"1","content":"","hint":""}`)
			} else {
				c.String(200, "success|marketUser")
			}
		} else {
			if v.Name == "gsm-u" || v.Name == "gsm-m" {
				c.String(http.StatusOK, `{"resp":"0","content":"密码错误","hint":""}`)
			} else {
				c.String(200, "password_error|noRoot")
			}
		}
	}
}

func connectCliReq(c *gin.Context) {
	// var exit bool
	h, _ := c.Get("hostInfo")
	v := h.(HostInfo)
	if cmd, exit := c.GetPostForm("cmd"); exit {
		if v.Name == "main" {
			res := fake.GetResStringDashboard(cmd)
			c.String(http.StatusOK, res)
		}
		if v.Name == "tdd" || v.Name == "fdd-u" || v.Name == "fdd-t" {
			res := fakelte.GetResStringLTE(cmd, v.Name, v.SN)
			c.String(http.StatusOK, res)
		}
	} else {
		c.String(200, "请求参数错误")
	}
}

func generalReq(c *gin.Context) {
	// var exit bool
	if caseType, exit := c.GetPostForm("caseType"); exit {
		switch caseType {
		case "5":
			c.String(http.StatusOK, "16")
		case "7":
			c.String(http.StatusOK, "success")
		}
	} else {
		c.String(http.StatusOK, "请求参数错误")
	}
}

//MixInfo ...
type MixInfo struct {
	SN       string
	Devinfo  interface{} `json:"dev"`
	HostInfo HostInfo    `json:"host"`
}

func getOnlineDevHostInfo(c *gin.Context) {
	dev.Mu.RLock()
	defer dev.Mu.RUnlock()

	readHostInfo()
	var keys []string
	for k := range dev.DeviceMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	infos := make([]MixInfo, 0)
	infos = append(infos, MixInfo{
		SN:       "masterControl",
		HostInfo: gethostInfoBySN("masterControl"),
	})

	for _, v := range keys {
		_, in := dev.DeviceMap[v].GetBasicInfo()
		for _, d := range in {
			s := d["SN"].(string)
			infos = append(infos, MixInfo{
				SN:       s,
				Devinfo:  d,
				HostInfo: gethostInfoBySN(s),
			})

		}
	}

	c.JSON(http.StatusOK, responseJSON{Code: 2000, Data: gin.H{"eqlist": infos}})
}

func gethostInfoBySN(sn string) HostInfo {
	hostMap.Mu.RLock()
	defer hostMap.Mu.RUnlock()
	for _, v := range hostMap.DeviceMap {
		if sn == v.SN {
			return v
		}
	}
	return HostInfo{}
}

func oamapp(c *gin.Context) {
	param, _ := c.GetPostForm("params")
	h, _ := c.Get("hostInfo")
	v := h.(HostInfo)
	if cmd, exit := c.GetPostForm("cmd"); exit {
		res := fakegsm.GetResStringForGSM(cmd, param, v.Name, v.SN)
		c.String(http.StatusOK, res)
	} else {
		c.String(200, "请求参数错误")
	}
}
