package main

type Link struct {
	input  string
	output string
}

func NewLink(inputServer, outputServer string) *Link {
	return &Link{input: inputServer, output: outputServer}
}
