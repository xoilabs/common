package types

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Swap struct {
	Index       uint16
	Pair        common.Address
	Sender    common.Address
	Recipient common.Address
	Amount0     big.Int
	Amount1     big.Int
	Tx
}
