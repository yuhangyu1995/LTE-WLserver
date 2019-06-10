package common

import (
	"fmt"
	"time"
)

func httpUpload(http *HTTPConfigure) {
	t1 := time.NewTicker(time.Second * time.Duration(http.UploadCycle))
	for {
		select {
		case <-t1.C:
			AppContext.DB.MemoryList.doPost(http.URL)
			fmt.Println(time.Now())
		default:
			time.Sleep(time.Millisecond * 100)
		}

	}
}
