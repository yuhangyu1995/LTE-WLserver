package lte

const (
	msgHeaderL                 uint16 = 32
	maxIMSILen                        = 17
	maxIMEILen                        = 17
	maxAddBlackAndWhiteListNum        = 10
	maxBlackAndWhiteListNum           = 100
	maxLocationBlackListNum           = 20
	maxQueryLocationNum               = 100
	maxDefaultArfcnNum                = 50
	maxAddRedirectListNum             = 20
)

// MsgHead UDP msg head struct
type MsgHead struct {
	FrameHeader uint32
	MsgID       uint16
	MsgLength   uint16
	WorkMode    uint16
	SysCode     uint16
	SN          [20]uint8
}
