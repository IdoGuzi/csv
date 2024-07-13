package operators

import (
	"strconv"
)

// AverageIntegerOperator is Operator implementing calculation of average on integers
// average - current average
// count - current number on integers processed
// columnNumber - from which column of the row to take the integer
type AverageIntegerOperator struct {
	average      *int
	count        *uint
	columnNumber uint
}

// AverageInteger create the Operator for calculating average of integers
func AverageInteger(columnNumber uint) AverageIntegerOperator {
	var count uint = 0
	var average int = 0
	operator := AverageIntegerOperator{
		count:        &count,
		average:      &average,
		columnNumber: columnNumber,
	}
	return operator
}

// Operate recalculate the new average after adding a new integer to the calculation
// return - the new average, nil on success, nil and an error otherwise
func (aio AverageIntegerOperator) Operate(data []string) ([]string, error) {
	intToAdd, err := strconv.Atoi(data[aio.columnNumber])
	if err != nil {
		return nil, err
	}
	sum := (*aio.average)*int((*aio.count)) + intToAdd
	*aio.count = (*aio.count) + 1
	*aio.average = sum / int(*aio.count)
	return []string{strconv.Itoa(*aio.average)}, nil
}
