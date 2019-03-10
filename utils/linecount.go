package utils

import (
	"bytes"
)

func LineCounter(data []byte) float64 {
	return float64(bytes.Count(data, []byte("\n")))
}
