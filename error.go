package gofilter

import "errors"

var (
	// filter error
	ErrFilterOffsetNegative   = errors.New("offset must be greater than 0")
	ErrFilterOffsetInvalid    = errors.New("invalid offset")
	ErrFilterLimitNegative    = errors.New("limit must be greater than 0")
	ErrFilterLimitInvalid     = errors.New("invalid limit")
	ErrFilterInvalidOrder     = errors.New("invalid order")
	ErrFilterInvalidOrderType = errors.New("invalid order type")
	ErrFilterInvalidOperator  = errors.New("invalid operator")
)
