package common

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// const CodeType
const (
	DefaultCodespace = "common"

	CodeInternalError              sdk.CodeType = 60001
	CodeInvalidPaginateParam       sdk.CodeType = 60002
	CodeCreateAddrFromBech32Failed sdk.CodeType = 60003
	CodeMarshalJSONFailed          sdk.CodeType = 60004
	CodeUnMarshalJSONFailed        sdk.CodeType = 60005
)

type SDKError struct {
	Codespace string       `json:"codespace"`
	Code      sdk.CodeType `json:"code"`
	Message   string       `json:"message"`
}

func ParseSDKError(errMsg string) SDKError {
	var sdkErr SDKError
	err := json.Unmarshal([]byte(errMsg), &sdkErr)
	if err != nil {
		sdkErr = SDKError{
			Codespace: DefaultCodespace,
			Code:      CodeInternalError,
			Message:   "internal error",
		}
		return sdkErr
	}
	return sdkErr
}

// invalid paginate param
func ErrInvalidPaginateParam(page int, perPage int) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidPaginateParam, fmt.Sprintf("invalid params: page=%d or per_page=%d", page, perPage))
}

// invalid address
func ErrCreateAddrFromBech32Failed(addr string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeCreateAddrFromBech32Failed, fmt.Sprintf("invalid address：%s", addr))
}

// could not marshal result to JSON
func ErrMarshalJSONFailed(msg string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeMarshalJSONFailed, fmt.Sprintf("could not marshal result to JSON, %s", msg))
}

// could not unmarshal result to origin
func ErrUnMarshalJSONFailed(msg string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeUnMarshalJSONFailed, fmt.Sprintf("UnMarshal JSON failed, %s", msg))
}
