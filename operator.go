package csv

type Operator interface {
	Operate(args []any) (any, error)
}
