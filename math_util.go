package godemo

import (
)

func BigdecimalMulty(m *int64) {
	for i := 0; i < 8; i++ {
		m = m * i
	}
	return m
}
