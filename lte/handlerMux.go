package lte

import (
	"fmt"
	"sync"
)

//Handler ...
type Handler interface {
	DealMsg(*msgDispose)
}

//HandlerFunc ...
type HandlerFunc func(*msgDispose)

// DealMsg calls f(h).
func (f HandlerFunc) DealMsg(h *msgDispose) {
	f(h)
}

//handlerMux ...
type handlerMux struct {
	mu sync.RWMutex
	m  map[uint16]muxEntry
}

type muxEntry struct {
	h Handler
}

// DefaultHandlerMux is the default ServeMux used by Serve.
var DefaultHandlerMux = &defaultHandlerMux

var defaultHandlerMux handlerMux

// Handle registers the handler for the given k.
func (mux *handlerMux) Handle(k uint16, handler Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if (k&0xFF00 != 0xF000) || (k&0x00FF) < 0x01 || (k&0x00FF) > 0x9E {
		panic("invalid MsgID")
	}
	if handler == nil {
		panic("nil handler")
	}
	if _, exist := mux.m[k]; exist {
		panic(fmt.Sprintf("multiple registrations for %#X", k))
	}

	if mux.m == nil {
		mux.m = make(map[uint16]muxEntry)
	}
	mux.m[k] = muxEntry{h: handler}
}

// HandleFunc registers the handler function for the given pattern.
func (mux *handlerMux) HandleFunc(k uint16, handler func(*msgDispose)) {
	mux.Handle(k, HandlerFunc(handler))
}

func (mux *handlerMux) matchMsg(k uint16, d *msgDispose) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()
	if v, exist := mux.m[k]; exist {
		if nil != v.h {
			v.h.DealMsg(d)
		} else {
			fmt.Printf("msg ID %#xDeal handler is nil,msg buf is %#x\n", k, d.msgbuf)
		}
	} else {
		fmt.Printf("unknown msg ID for %#x,msg buf is %#x\n", k, d.msgbuf)
	}
}

//AddHandleFunc ...
func AddHandleFunc(k uint16, handler func(*msgDispose)) {
	DefaultHandlerMux.HandleFunc(k, handler)
}

//DealMsg ...
func DealMsg(d *msgDispose) {
	DefaultHandlerMux.matchMsg(d.head.MsgID, d)
}
