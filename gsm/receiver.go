package gsm

import (
	"dkay/fake-wlserver/common"
	"fmt"
	"net"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

var (
	//BaseTimeFormatStr ..
	BaseTimeFormatStr = "20060102150405"
	conn              *net.UDPConn
	log               = common.Log
	app               = &common.AppContext
	dev               = &common.AppContext.RealDev
	json              = jsoniter.ConfigCompatibleWithStandardLibrary
	mysql             = &common.AppContext.DB.Mysql
	memlist           = &common.AppContext.DB.MemoryList
)

var rec receiver

type receiver struct {
	Mu     sync.RWMutex
	revMap map[string]interface{}
}

//StartRCV ...
func StartRCV() {
	initUDP()
	fmt.Println("start rev gsm msg")
	for {
		buf := make([]byte, 4096, 4096)
		n, raddr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			continue
		}

		if 0 >= n {
			log.Warnln("received empty message buffer from", raddr.IP.String())
			continue
		}

		disp := &msgDispose{
			addr:   raddr,
			msgbuf: buf[0:n],
			msgL:   uint16(n),
		}
		log.Info("received message", string(buf[0:n]), " from ", raddr.IP.String())

		go disp.do()

	}
}

func initUDP() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+common.AppContext.Config.UDPCfg.GSM)
	common.CheckError(err)
	conn, err = net.ListenUDP("udp", udpAddr)
	common.CheckError(err)
}

func init() {
	rec.revMap = make(map[string]interface{})
}
