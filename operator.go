package csv

type Operator interface {
	Operate(data []string) ([]string, error)
}
