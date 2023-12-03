package types

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	Address     common.Address
	Name        string
	Symbol      string
	Decimal     common.Decimal
	Creator     common.Address
	TxHash      common.Hash
	BlockTime   time.Time
	BlockNumber int64
}
