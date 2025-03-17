package gofilter

type Operator string

var (
	Equal     Operator = "eq"
	NotEqual  Operator = "ne"
	Greater   Operator = "gt"
	GreaterEq Operator = "gte"
	Less      Operator = "lt"
	LessEq    Operator = "lte"
	Like      Operator = "like"
	In        Operator = "in"
	NotIn     Operator = "nin"
)

func StringToOperator(operator string) (Operator, error) {
	switch operator {
	case "eq":
		return Equal, nil
	case "ne":
		return NotEqual, nil
	case "gt":
		return Greater, nil
	case "gte":
		return GreaterEq, nil
	case "lt":
		return Less, nil
	case "lte":
		return LessEq, nil
	case "like":
		return Like, nil
	case "in":
		return In, nil
	case "nin":
		return NotIn, nil
	default:
		return Equal, ErrFilterInvalidOperator
	}
}
