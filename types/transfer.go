package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Transfer struct {
	Index  uint16
	Token  common.Address
	From   common.Address
	To     common.Address
	Amount *big.Int
}