package types

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// const uint32
const (
	DefaultCodespace = "backend"

	CodeNewErrorsMergedFailed         uint32 = 62001
	CodeProductIsRequired             uint32 = 62002
	CodeAddressIsEmpty                uint32 = 62003
	CodeOrderStatusMustBeOpenOrClosed uint32 = 62004
	CodeAddressAndProductRequired     uint32 = 62005
	CodeGetChainHeightFailed          uint32 = 62006
	CodeGetBlockTxHashesFailed        uint32 = 62007
	CodeOrderSideMustBuyOrSell        uint32 = 62008
	CodeProductDoesNotExist           uint32 = 62009
	CodeBackendPluginNotEnabled       uint32 = 62010
	CodeRecoverPanicGoroutineFailed   uint32 = 62011
	CodeUnknownBackendEndpoint        uint32 = 62012
	CodeGetCandlesFailed              uint32 = 62013
	CodeGetCandlesByMarketFailed      uint32 = 62014
	CodeGetTickerByProductsFailed     uint32 = 62015
	CodeParamNotCorrect				  uint32 = 62016
	CodeNoKlinesFunctionFound		  uint32 = 62017
	CodeMarketkeeperNotInitialized	  uint32 = 62018
)

// invalid param side, must be buy or sell
func ErrOrderSideParamMustBuyOrSell(side string) sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeOrderSideMustBuyOrSell, fmt.Sprintf("Side should not be %s", side))
}

// product is required
func ErrProductIsRequired() sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeProductIsRequired, "invalid params: product is required")
}

// product does not exist
func ErrProductDoesNotExist(product string) sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeProductDoesNotExist, fmt.Sprintf("product %s does not exist", product))
}

func ErrBackendPluginNotEnabled(message string) sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeBackendPluginNotEnabled, message)
}

func ErrParamNotCorrect(message string) sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeParamNotCorrect, message)
}

func ErrNoKlinesFunctionFound(message string) sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeNoKlinesFunctionFound, message)
}

func ErrMarketkeeperNotInitialized(message string) sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeMarketkeeperNotInitialized, message)
}

func ErrUnknownBackendEndpoint(message string) sdk.Error {
	return sdkerrors.New(DefaultCodespace, CodeUnknownBackendEndpoint, message)
}