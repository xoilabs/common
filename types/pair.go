package types

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Pair struct {
	Address     common.Address
	Token0      common.Address
	Token1      common.Address
	Creator     common.Address
	Tx
}
