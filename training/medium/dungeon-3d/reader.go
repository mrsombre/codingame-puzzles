package main

import (
	"bufio"
	"fmt"
	"io"
)

const scannerSize = 1000000

func NewScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, scannerSize), scannerSize)
	return scanner
}

func ReadGame(s *bufio.Scanner) []string {
	var result []string

	s.Scan()
	result = append(result, s.Text())
	s.Scan()
	result = append(result, s.Text())

	var ln int
	fmt.Sscan(s.Text(), &ln)

	for i := 0; i < ln; i++ {
		s.Scan()
		result = append(result, s.Text())
	}

	return result
}
