package types

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Swap struct {
	Index       uint16
	TxHash      common.Hash
	Pair        common.Address
	TxSender    common.Address
	TxRecipient common.Address
	Amount0     big.Int
	Amount1     big.Int
	BlockTime   time.Time
}
