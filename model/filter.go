package model

import "github.com/anchamber/genetics-api/proto"

type Filter struct {
	Key      string
	Operator Operator
	Value    interface{}
}

type Operator int32

const (
	EQ         Operator = Operator(proto.Operator_EQ)
	GREATER    Operator = Operator(proto.Operator_GREATER)
	SMALLER    Operator = Operator(proto.Operator_SMALLER)
	GREATER_EQ Operator = Operator(proto.Operator_GREATER_EQ)
	SMALLER_EQ Operator = Operator(proto.Operator_SMALLER_EQ)
	CONTAINS   Operator = Operator(proto.Operator_CONTAINS)
)

func NewFilterFromProto(filter *proto.Filter) *Filter {
	transformed := &Filter{
		Key:      filter.Key,
		Operator: Operator(filter.Operator),
	}
	if _, ok := filter.GetValue().(*proto.Filter_S); ok {
		transformed.Value = filter.GetS()
	} else if _, ok := filter.GetValue().(*proto.Filter_I); ok {
		transformed.Value = filter.GetI()
	} else if _, ok := filter.GetValue().(*proto.Filter_D); ok {
		transformed.Value = filter.GetD()
	} else if _, ok := filter.GetValue().(*proto.Filter_B); ok {
		transformed.Value = filter.GetB()
	}

	return transformed
}
