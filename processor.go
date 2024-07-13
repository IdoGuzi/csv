package csv

import (
	"encoding/csv"
	"io"
	"os"
)

// Read create a csv file @Processor
// filePath - is the path for csv file to process
// return - a new @Processor
func Read(filePath string) Processor {
	return Processor{inputFile: filePath, Operators: []Operator{}}
}

// Processor collect the needed data and operators for csv processing
// inputFile - path to csv file to process
// outputFile - path of output csv file
// Operators - an ordered list of operations used to process a csv file
type Processor struct {
	inputFile  string
	outputFile string
	Operators  []Operator
}

// Write set the path of the output csv file and start processing
func (p Processor) Write(filePath string) error {
	p.outputFile = filePath
	return p.read()
}

// With add an Operator to the Processor
// return the Processor to function chaining purposes
func (p Processor) With(operator Operator) Processor {
	p.Operators = append(p.Operators, operator)
	return p
}

// read is the actual input csv file and call the proccessing functions
// return error if anything went wrong
func (p Processor) read() error {
	inputFile, err := os.Open(p.inputFile)
	if err != nil {
		return err
	}
	outputFile, err := os.OpenFile(p.outputFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	defer outputFile.Close()
	reader := csv.NewReader(inputFile)
	writer := csv.NewWriter(outputFile)
	var record []string
	for {
		record, err = reader.Read()
		if err != nil {
			break
		}
		err = p.process(writer, record)
		if err != nil {
			return err
		}
	}
	if err != nil {
		if err != io.EOF {
			return err
		}
	}
	writer.Flush()
	return nil
}

func util(data ...any) []any {
	return data
}

// process call all the operators added to Processor in order on a given record
// and write the result to the output file
// writer - of the output csv file
// record - current csv row to process
// return - nil for success, error otherwise
func (p Processor) process(writer *csv.Writer, record []string) error {
	var err error
	var result []string = record
	//by order execute the operators
	for _, operator := range p.Operators {
		result, err = operator.Operate(result)
		if err != nil {
			return err
		}
		writer.Write(result)
	}
	return nil
}
