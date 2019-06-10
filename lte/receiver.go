package lte

import (
	"dkay/fake-wlserver/common"
	"fmt"
	"net"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

var (
	app                    = &common.AppContext
	dev                    = &common.AppContext.RealDev
	json                   = jsoniter.ConfigCompatibleWithStandardLibrary
	mysql                  = &common.AppContext.DB.Mysql
	memlist                = &common.AppContext.DB.MemoryList
	checkError             = common.CheckError
	printError             = common.PrintError
	bytetoString           = common.BytetoString
	readBufUseLittleEndian = common.ReadBufUseLittleEndian
)

var rec receiver

type receiver struct {
	Mu     sync.RWMutex
	revMap map[uint16]interface{}
}

//StartRCV ...
func StartRCV() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":"+common.AppContext.Config.UDPCfg.LTE)
	common.CheckError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	common.CheckError(err)

	fmt.Println("start rev lte msg")
	for {
		buf := make([]byte, 1024, 2048)
		n, raddr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			continue
		}

		if 0 >= n {
			fmt.Println("recviver msg length error from ", raddr.IP)
			continue
		}

		disp := &msgDispose{
			addr:   raddr.IP.String(),
			msgbuf: buf[0:n],
			msgL:   uint16(n),
		}
		// fmt.Println("time:", time.Now(), "rev msg", hex.EncodeToString(buf[0:n]))

		go disp.do()

	}
}

func init() {
	rec.revMap = make(map[uint16]interface{})
	initHandFunc()
}

func initHandFunc() {
	AddHandleFunc(0xF004, dealCommonACK)
	AddHandleFunc(0xF005, dealF005)
	AddHandleFunc(0xF00A, dealF00A)
	AddHandleFunc(0xF00E, dealCommonACK)
	AddHandleFunc(0xF010, dealF010)
	AddHandleFunc(0xF014, dealCommonACK)
	AddHandleFunc(0xF016, dealCommonACK)
	AddHandleFunc(0xF019, dealF019)
	AddHandleFunc(0xF020, dealCommonACK)
	AddHandleFunc(0xF024, dealCommonACK)
	AddHandleFunc(0xF028, dealServingCellInfoQueryACK)
	AddHandleFunc(0xF02A, dealCommonACK)
	AddHandleFunc(0xF02C, DealEnbBaseInfoQueryACK)
	AddHandleFunc(0xF02E, dealSyncInfoQueryACK)
	AddHandleFunc(0xF032, dealRxTxPowerAck)
	AddHandleFunc(0xF03C, dealCommonACK)
	AddHandleFunc(0xF042, dealSelfActiveQueryACK)
	AddHandleFunc(0xF04A, dealCommonACK)
	AddHandleFunc(0xF04C, dealTddSSAndAplphaQueryACK)
	AddHandleFunc(0xF05D, dealGPSInfoQueryACK)
	AddHandleFunc(0xF066, dealRRCAccessQueryACK)
	AddHandleFunc(0xF06C, dealCommonACK)
	AddHandleFunc(0xF06E, dealCommonACK)
	AddHandleFunc(0xF074, dealGPSPPSQueryACK)
	AddHandleFunc(0xF081, dealCommonACK)
	AddHandleFunc(0xF076, dealCommonACK)
	AddHandleFunc(0xF093, dealCommonACK)
	AddHandleFunc(0xF09C, dealNtpStateACK)

}

func getMSGEntity(msgID uint16) interface{} {
	switch msgID {
	case 0xF003:
		return &ServingCellInfo{}
	case 0xF009:
		return &CellSearchEarfcnCfg{}
	case 0xF00B:
		return &RebootCfg{}
	case 0xF00D:
		return &CellActive{}
	case 0xF013:
		return &RxGainCfg{}
	case 0xF015:
		return &PowerDereaseCfg{}
	case 0xF023:
		return &REMModeCfg{}
	case 0xF029:
		return &DelayPPS{}
	case 0xF03B:
		return &PowerOnCfg{}
	case 0xF042:
		return &SelfActiveQueryACK{}
	case 0xF049:
		return &TDDSFCfg{}
	case 0xF075:
		return &NTPServerIPCfg{}
	case 0xF080:
		return &CellParaHotCfg{}
	case 0xF092:
		return &ULAlphaCfg{}
	default:
		return nil
	}
}
