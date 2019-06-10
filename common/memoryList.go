package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync/atomic"
	"time"
	"unsafe"
)

type memoryList struct {
	lock     int32
	priority int32
	list     []UeInfoJSON
}

//UeInfoJSON ...
type UeInfoJSON struct {
	SN   string `json:"dc"`
	Imsi string `json:"is"`
	Imei string `json:"ie"`
	Time int64  `json:"t"`
}


func (m *memoryList) doInit() {
	m.list = make([]UeInfoJSON, 0)
}

func (m *memoryList) Add(ue UeInfoJSON) {
	for {
		if priority := atomic.LoadInt32(&m.priority); 0 == priority {
			if atomic.CompareAndSwapInt32(&m.lock, 0, 1) {
				m.list = append(m.list, ue)
				atomic.StoreInt32(&m.lock, 0)
				return
			}
		}
	}
}

func (m *memoryList) doPost(url string) {
	atomic.StoreInt32(&m.priority, 1)
	//restore default priority&lock
	defer func() {
		m.list = append(m.list[:0], m.list[:0]...)
		atomic.StoreInt32(&m.lock, 0)
		atomic.StoreInt32(&m.priority, 0)
	}()

	for {
		if atomic.CompareAndSwapInt32(&m.lock, 0, 1) {
			if 0 == len(m.list) {
				return
			}
			buff, err := json.Marshal(m.list)
			fmt.Println(string(buff))
			if nil != err {
				fmt.Println(err.Error())
				return
			}

			reader := bytes.NewReader(buff)
			req, err := http.NewRequest("POST", url, reader)
			if nil != err {
				fmt.Println(err.Error())
				return
			}

			req.Header.Set("Content-Type", "application/json;charset=UTF-8")
			//use http DefaultClient fix Repeated Client
			http.DefaultClient.Timeout = time.Second * 5
			resp, err := http.DefaultClient.Do(req)

			//Close resp.Body to prevent possible resource leaks
			if nil != resp {
				defer resp.Body.Close()
			}

			if nil != err {
				fmt.Println(err.Error())
				return
			}

			respBytes, err := ioutil.ReadAll(resp.Body)
			if nil != err {
				fmt.Println(err.Error())
				return
			}

			//byte数组直接转成string，优化内存
			str := (*string)(unsafe.Pointer(&respBytes))
			fmt.Println(*str)

			return
		}
	}
}
