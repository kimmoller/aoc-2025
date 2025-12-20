package main

func PassowrdFromDial(data []string, countRotations bool) (*int, error) {
	dial := NewDial(50, 0, 99)

	for _, input := range data {
		err := dial.Turn(input)
		if err != nil {
			return nil, err
		}
	}

	if countRotations {
		sum := dial.atZero + dial.crossedZero
		return &sum, nil
	}

	return &dial.atZero, nil
}
