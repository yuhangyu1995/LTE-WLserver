package lte

// EnbBaseInfoQuery msgID 0xF02B
type EnbBaseInfoQuery struct {
	InfoType uint32
}

// ServingCellInfoQuery msgID 0xF027
type ServingCellInfoQuery struct {
}

// SyncInfoQuery msgID 0xF02D
type SyncInfoQuery struct {
}

// CellStateQuery msgID 0xF02F
type CellStateQuery struct {
}

// RxTxGainQuery msgID 0xF031
type RxTxGainQuery struct {
}

// RedirectQuery msgID 0xF03F
type RedirectQuery struct {
}

// SelfActiveQuery msgID 0xF041
type SelfActiveQuery struct {
}
