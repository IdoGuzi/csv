package operators

import (
	"errors"
	"fmt"
)

// AverageIntegerOperator is Operator implementing calculation of average on integers
// average - current average
// count - current number on integers processed
// columnNumber - from which column of the row to take the integer
type AverageIntegerOperator struct {
	average      int
	count        uint
	columnNumber uint
}

// AverageInteger create the Operator for calculating average of integers
func AverageInteger(columnNumber uint) AverageIntegerOperator {
	operator := AverageIntegerOperator{
		count:        0,
		average:      0,
		columnNumber: columnNumber,
	}
	return operator
}

// Operate recalculate the new average after adding a new integer to the calculation
// return - the new average, nil on success, nil and an error otherwise
func (aio AverageIntegerOperator) Operate(data ...any) (any, error) {
	aio.count++
	intToAdd, ok := data[aio.columnNumber].(int)
	if !ok {
		return nil, errors.New(fmt.Sprintf("cannot convert %v to int", data[aio.columnNumber]))
	}
	aio.average = (aio.average + intToAdd) / int(aio.count)
	return aio.average, nil
}
