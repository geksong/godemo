package utils

import ()

type UserMes struct {
	Name string
	Aag  int
}

// int mulity
func BigdecimalMulty(m int) int {
	for i := 1; i < 8; i++ {
		m = m * i
	}
	return m
}

// swap two int number
func SwapTwoInt(m1 int, m2 int) (int, int) {
	m1, m2 = m2, m1
	return m1, m2
}
