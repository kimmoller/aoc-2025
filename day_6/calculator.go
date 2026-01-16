package main

import (
	"slices"
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

func ProblemsFromData(data []string, isAlien bool) ([]Problem, error) {
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

	if isAlien {
		return alienProblems(dataWithoutOperators, operators)
	}

	return normalProblems(dataWithoutOperators, operators)

}

func normalProblems(dataWithoutOperators []string, operators []string) ([]Problem, error) {
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

	problems := []Problem{}
	for key, value := range problemsWithoutOperators {
		problem := Problem{operator: operators[key], inputs: value}
		problems = append(problems, problem)
	}

	return problems, nil
}

func alienProblems(dataWithoutOperators []string, operators []string) ([]Problem, error) {
	// Split input data into data layers
	dataLayers := map[int][]string{}
	for i, data := range dataWithoutOperators {
		values := strings.Split(data, "")
		layer := []string{}
		for _, value := range values {
			layer = append(layer, value)
		}
		dataLayers[i] = layer
	}

	// Combine data layers into actual numbers
	numbers := []int{}
	for i := range dataLayers[0] {
		numberAsList := []string{}
		for j := 0; j < len(dataLayers); j++ {
			layer := dataLayers[j]
			digit := layer[i]
			if digit != " " {
				numberAsList = append(numberAsList, digit)
			}
		}
		// Mark the gap between numbers with a -1
		if len(numberAsList) == 0 {
			numbers = append(numbers, -1)
			continue
		}
		numStr := strings.Join(numberAsList, "")
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	// Add a -1 to indicate that the numbers have ended
	numbers = append(numbers, -1)

	problemsWithoutOperators := [][]int{}
	problem := []int{}
	for _, number := range numbers {
		if number == -1 {
			problemsWithoutOperators = append(problemsWithoutOperators, slices.Clone(problem))
			problem = []int{}
		} else {
			problem = append(problem, number)
		}
	}

	problems := []Problem{}
	for key, value := range problemsWithoutOperators {
		problem := Problem{operator: operators[key], inputs: value}
		problems = append(problems, problem)
	}

	return problems, nil
}
