package main

import (
	"bufio"
)

func ReadGame(s *bufio.Scanner) []string {
	data := make([]string, 0, 32)

	s.Scan()
	size := s.Text()
	data = append(data, size)
	for i := 0; i < StrToInt(size); i++ {
		s.Scan()
		data = append(data, s.Text())
	}

	return data
}

func ReadStep(s *bufio.Scanner) []string {
	data := make([]string, 0, 1)

	s.Scan()
	data = append(data, s.Text())

	return data
}
