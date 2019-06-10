package gsm

import (
	"net"
	"strconv"
	"strings"
	"unicode"
)

//MsgHeadLen 消息头长度
var MsgHeadLen = 5

type msgDispose struct {
	addr   *net.UDPAddr
	msgbuf []byte
	msgL   uint16
}

type myMsg struct {
	srcMsg   string
	fromAddr *net.UDPAddr
	Head     MsgHead
	Body     string
}

//MsgHead ...
type MsgHead struct {
	ID      string
	Type    string
	Len     int
	WLName  string
	BoardID string
}

func (m *msgDispose) do() {
	var msg myMsg
	srcmsg := string(m.msgbuf)
	stopFlag := false

	msgSlice := strings.FieldsFunc(srcmsg, func(s rune) bool {
		if stopFlag {
			return false
		}
		if unicode.IsSpace(s) {
			return true
		}
		if '<' == s || '[' == s {
			stopFlag = true
		}
		return false
	})

	var err error
	msg.srcMsg = srcmsg
	msg.fromAddr = m.addr
	len, err := strconv.Atoi(msgSlice[2])
	if nil != err {
		log.Warnln("received illed msg from", m.addr, "msg is :", string(m.msgbuf))
	}
	head := MsgHead{}
	head.ID = msgSlice[0]
	head.Type = msgSlice[1]
	head.Len = len
	head.WLName = msgSlice[3]
	head.BoardID = msgSlice[4]

	msg.Head = head
	msg.Body = msgSlice[5]
	if "CmdAck" == head.Type {
		msg.Head.Type = msgSlice[5]
		msg.Body = msgSlice[6]
	}
	msg.ProcessMsgByID()
}

func (my *myMsg) ProcessMsgByID() {
	switch my.Head.ID {
	case "101":
		dealHeadBeatMSG(my)
	case "103":
		dealOneUeInfo(my)
	case "104":
		dealCellPara(my)
	case "105":
		dealWeiLanCfg(my)
	case "109":
		dealCmdAck(my)
	case "116":
		dealStatusRPT(my)
	default:
	}
}
