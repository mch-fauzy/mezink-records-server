package model

const (
	OperatorRange = "range"
)

type Filter struct {
	FilterFields []FilterField
}

type FilterField struct {
	Field    string
	Operator string `validate:"oneof=range"`
	Value    interface{}
}
