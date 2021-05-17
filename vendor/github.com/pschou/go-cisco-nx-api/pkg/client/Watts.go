package client

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Watts float32

func (w *Watts) UnmarshalText(text []byte) (err error) {
	in := strings.TrimSpace(string(text))
	if in == "N/A" {
		*w = Watts(math.NaN())
		return
	}
	if strings.HasSuffix(in, " W") {
		var val float64
		val, err = strconv.ParseFloat(strings.TrimSpace(in[:len(in)-2]), 32)
		*w = Watts(val)
		return
	}
	err = errors.New("parsing watt: unknown field " + string(text))
	return
}

func (w Watts) MarshalText() (text []byte, err error) {
	return []byte(w.String()), nil
}

func (w Watts) String() string {
	if math.IsNaN(float64(w)) {
		return "N/A"
	}
	return fmt.Sprintf("%g W", float32(w))
}
