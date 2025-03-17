package gofilter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/core-stack/goutils/stringutils"
)

const MaxLimit = 2000

func BuildQueryFilter(limit, offset string, order []string, query map[string][]string) (*Filter, error) {
	var filter Filter
	// Limit
	if limitQuery := limit; limitQuery != "" {
		limit, err := strconv.Atoi(limitQuery)
		if err != nil {
			return nil, ErrFilterLimitInvalid
		}
		if limit < 0 {
			return nil, ErrFilterLimitNegative
		}
		filter.Limit = limit
	}

	// Offset
	if offsetQuery := offset; offsetQuery != "" {
		offset, err := strconv.Atoi(offsetQuery)
		if err != nil {
			return nil, ErrFilterOffsetInvalid
		}
		if offset < 0 {
			return nil, ErrFilterOffsetNegative
		}
		filter.Offset = offset
	}

	// Order
	if orderQuery := order; len(orderQuery) > 0 {
		order, err := BuildOrderFromQuery(orderQuery)
		if err != nil {
			return nil, err
		}
		filter.Order = order
	}

	// Where
	where, err := BuildWhereFromQuery(query)
	if err != nil {
		return nil, err
	}
	filter.Where = *where

	return &filter, nil
}

func BuildOrderFromQuery(orderQuery []string) (Order, error) {
	order := make(Order)
	for _, value := range orderQuery {
		parts := strings.Split(value, ":")
		if len(parts) != 2 {
			return nil, ErrFilterInvalidOrder
		}

		field := strings.ToLower(parts[0])
		orderType := OrderType(strings.ToLower(parts[1]))
		if orderType != OrderAsc && orderType != OrderDesc {
			return nil, ErrFilterInvalidOrderType
		}

		order[field] = orderType
	}

	return order, nil
}

func BuildWhereFromQuery(query map[string][]string) (*Where, error) {
	params := query
	where := Where{}
	for key, values := range params {
		// ignore limit, offset and order //
		if key == "limit" || key == "offset" || key == "order" {
			continue
		}

		value := values[0]
		split := strings.Split(value, ":")
		if len(split) < 2 {
			return nil, fmt.Errorf("invalid where format: %s", value)
		}

		operator, err := StringToOperator(split[0])
		if err != nil {
			return nil, err
		}
		cond := Condition{
			Field:    stringutils.ToSnakeCase(key),
			Operator: operator,
		}

		switch operator {
		case Equal, NotEqual, Greater, GreaterEq, Less, LessEq:
			cond.Value = split[1]
		case Like:
			cond.Value = "%" + split[1] + "%"
		case In, NotIn:
			cond.Value = strings.Split(split[1], ",")
		}
		where = append(where, cond)
	}
	return &where, nil
}
