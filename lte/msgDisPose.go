package lte

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

var msgTempBuf = make(map[string]map[uint16]*msgDispose)

type msgDispose struct {
	addr   string
	msgbuf []byte
	msgL   uint16
	head   *MsgHead
}

func (d *msgDispose) do() {
	err := d.initMsgHead()
	if nil != err {
		fmt.Println(err, "msg is", hex.EncodeToString(d.msgbuf))
		return
	}

	//debug
	if 0xf00c == d.head.MsgID {
		fmt.Println("watch msg", hex.EncodeToString(d.msgbuf))
	}

	//init tempBUf
	if _, exits := msgTempBuf[d.addr]; !exits {
		msgTempBuf[d.addr] = make(map[uint16]*msgDispose)
	}

	//processing message segmentation
	//u16 Syscode,the highest bit equals 1 for the message is not complete
	//the highest bit equals 0 for message completion
	highBit := d.head.SysCode & 0x8000
	low7Bit := d.head.SysCode & 0x7FFF
	switch highBit {
	case 0x8000:
		msgTempBuf[d.addr][low7Bit] = d
		go cleanTempBuf(d.addr, low7Bit)
	default:
		if v, exits := msgTempBuf[d.addr][low7Bit]; exits {
			v.head.MsgLength += d.head.MsgLength
			v.msgbuf = append(v.msgbuf, d.msgbuf[msgHeaderL:]...)
			DealMsg(v)
		} else {
			DealMsg(d)
		}
	}
}

func (d *msgDispose) initMsgHead() error {
	var err error
	buf := bytes.NewBuffer(d.msgbuf)
	d.head = new(MsgHead)

	if 0xaa == d.msgbuf[0] {
		err = binary.Read(buf, binary.LittleEndian, d.head)
	} else if 0xbb == d.msgbuf[0] {
		err = binary.Read(buf, binary.BigEndian, d.head)
	} else {
		err = errors.New("unknown msg")
	}

	if d.head.MsgLength != d.msgL {
		err = errors.New("msg length errer")
	}
	return err
}

func cleanTempBuf(k1 string, k2 uint16) {
	time.Sleep(time.Second * 3)
	delete(msgTempBuf[k1], k2)
}
