package identity

import (
	"github.com/cosmos/cosmos-sdk/wire"
)

// Register concrete types on wire codec
func RegisterWire(cdc *wire.Codec) {
	cdc.RegisterConcrete(MsgCreateClaim{}, "identity/CreateClaim", nil)
	cdc.RegisterConcrete(MsgRevokeClaim{}, "identity/RevokeClaim", nil)
	cdc.RegisterConcrete(MsgRevokeClaim{}, "identity/AnswerClaim", nil)
}
