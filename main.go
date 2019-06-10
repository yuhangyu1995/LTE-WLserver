package main

import (
	"dkay/fake-wlserver/common"
	"dkay/fake-wlserver/gsm"
	"dkay/fake-wlserver/lte"
	"dkay/fake-wlserver/web"
	"dkay/fake-wlserver/fake"
)

func main() {
	common.Log.Infoln("Server start...")
	go gsm.StartRCV()
	go lte.StartRCV()
	go fake.DiagnosisInfoRefresh()
	web.Start()

}

func init() {
	common.InitApp()
}
