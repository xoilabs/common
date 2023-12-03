package types

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Swap struct {
	TxHash      common.Hash
	Pair        common.Address
	TxFrom      common.Address
	TxTo        common.Address
	Amount0     big.Int
	Amount1     big.Int
	BlockTime   time.Time
	BlockNumber int64
}
