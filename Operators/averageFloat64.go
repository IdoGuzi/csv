package operators

import (
	"strconv"
)

// AverageFloatOperator is Operator implementing calculation of average on float64s
// average - current average
// count - current number on float64s processed
// columnNumber - from which column of the row to take the float64
type AverageFloat64Operator struct {
	average      *float64
	count        *uint
	columnNumber uint
}

// AverageFloat64 create the Operator for calculating average of float64s
func AverageFloat64(columnNumber uint) AverageFloat64Operator {
	var count uint = 0
	var average float64 = 0
	operator := AverageFloat64Operator{
		count:        &count,
		average:      &average,
		columnNumber: columnNumber,
	}
	return operator
}

// Operate recalculate the new average after adding a new float64 to the calculation
// return - the new average, nil on success, nil and an error otherwise
func (afo AverageFloat64Operator) Operate(data []string) ([]string, error) {
	floatToAdd, err := strconv.ParseFloat(data[afo.columnNumber], 64)
	if err != nil {
		return nil, err
	}
	sum := (*afo.average)*float64(*afo.count) + floatToAdd
	*afo.count = (*afo.count) + 1
	*afo.average = sum / float64(*afo.count)
	return []string{strconv.FormatFloat(*afo.average, 'f', -1, 64)}, nil
}
