package compute

type AddOperator interface {
	Add () interface {}
}

type IntParams struct {
	P1 int
	P2 int
}

func (params *IntParams) Add() interface {} {
	return params.P1 + params.P2
}