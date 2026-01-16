package main

type Link struct {
	output string
}

func NewLink(outputServer string) Link {
	return Link{output: outputServer}
}
