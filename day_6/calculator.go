package main

import (
	"strconv"
	"strings"
)

const (
	ADD       = "+"
	SUBSTRACT = "-"
	MULTIPLY  = "*"
)

type Problem struct {
	inputs   []int
	operator string
}

type Calculator struct {
	problems []Problem
}

func NewCalculator(problems []Problem) *Calculator {
	return &Calculator{problems: problems}
}

func (c *Calculator) SumOfProblems() int {
	total := 0

	for _, problem := range c.problems {
		sum := 0
		for i, number := range problem.inputs {
			if problem.operator == ADD {
				sum += number
			}
			if problem.operator == SUBSTRACT {
				sum -= number
			}
			if problem.operator == MULTIPLY {
				if i == 0 {
					sum += number
				} else {
					sum *= number

				}
			}
		}
		total += sum
	}

	return total
}

func ProblemsFromData(data []string) ([]Problem, error) {
	problems := []Problem{}

	dataWithoutOperators := data[:len(data)-1]
	operatorString := data[len(data)-1]

	// Build array of operators
	dirtyOperators := strings.Split(operatorString, " ")
	operators := []string{}
	for _, operator := range dirtyOperators {
		if operator != "" {
			operators = append(operators, operator)
		}
	}

	// Build an array with all number inputs
	inputs := []string{}
	for _, dataString := range dataWithoutOperators {
		values := strings.Split(dataString, " ")
		for _, value := range values {
			if value != "" {
				inputs = append(inputs, value)
			}
		}
	}

	problemsWithoutOperators := map[int][]int{}
	for i, input := range inputs {
		key := i % len(operators)
		value, err := strconv.Atoi(input)
		if err != nil {
			return nil, err
		}
		problemsWithoutOperators[key] = append(problemsWithoutOperators[key], value)
	}

	for key, value := range problemsWithoutOperators {
		problem := Problem{operator: operators[key], inputs: value}
		problems = append(problems, problem)
	}

	return problems, nil
}
