package types

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	From        common.Address
	To          common.Address
	Tx
}
