package gofilter

type Condition struct {
	Field    string
	Operator Operator
	Value    any
}

type Where []Condition

type OrderType string

const (
	OrderAsc  OrderType = "asc"
	OrderDesc OrderType = "desc"
)

type Order map[string]OrderType

type Filter struct {
	Where   Where
	Limit   int
	Offset  int
	Order   Order
	Select  []string
	Join    []string
	Preload []string
}

func (f *Filter) AddWhere(condition Condition) {
	f.Where = append(f.Where, condition)
}
