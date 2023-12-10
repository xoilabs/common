package types

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Tx struct {
	TxHash common.Hash
	BlockTime   time.Time
	BlockNumber int64
}