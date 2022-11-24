package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/handson module sentinel errors
var (
	ErrSample         = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrIncorrectDenom = sdkerrors.Register(ModuleName, 11, "wrong denom")
)
