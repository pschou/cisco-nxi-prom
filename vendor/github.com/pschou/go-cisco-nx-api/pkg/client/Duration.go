package client

import (
	"errors"
	"fmt"
)

type durationUnit struct {
	U   byte
	Val Duration
}

var durationUnits = []durationUnit{
	{'D', 24 * 3600e9},
	{'H', 3600e9},
	{'M', 60e9},
	{'S', 1e9},
}

type Duration uint64

func (d *Duration) UnmarshalText(text []byte) (err error) {
	*d, err = ParseDuration(string(text))
	return
}
func (d Duration) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d Duration) String() (ret string) {
	if d < 1e9 {
		return "PT0S"
	}
	ret = "P"
	for i, u := range durationUnits {
		if d >= u.Val {
			st := d / u.Val

			d -= u.Val * st
			ret = ret + fmt.Sprintf("%d%c", st, u.U)
		}
		if i == 0 {
			ret = ret + "T"
		}
	}
	return
}

func ParseDuration(d string) (durationVal Duration, err error) {
	if d[0] != 'P' && !(d[0] >= '0' && d[0] <= '9') {
		return 0, errors.New("cisco time: must start with 'P' or a number, found " + d)
	}
	if d[0] == 'P' {
		if len(d) < 4 {
			return 0, errors.New("cisco time: incomplete / no value")
		}
		val := 0
		unit := 0
		for i := 1; i < len(d); i++ {
			switch {
			case d[i] >= '0' && d[i] <= '9':
				val = val*10 + int(d[i]) - '0'
			case d[i] == 'T':
				if val > 0 {
					return 0, errors.New("cisco time: malformed time section")
				}
				if unit <= 1 {
					unit = 1
				}
			default:
				for unit < len(durationUnits) && d[i] != durationUnits[unit].U {
					unit++
				}
				if unit == len(durationUnits) {
					return 0, errors.New("cisco time: found unusable duration: " + string(d))
				}
				if d[i] == durationUnits[unit].U {
					durationVal += durationUnits[unit].Val * Duration(val)
					val = 0
					unit++
				} else {
					return 0, errors.New("cisco time: found extra numeric " + string(d[i]) + " in " + string(d))
				}
			}
		}
	} else {
		if len(d) != 8 || d[2] != ':' || d[5] != ':' {
			return 0, errors.New("cisco time: must be in the format 00:00:00, found " + d)
		}
		val := 0
		for i := 0; i < 9; i += 3 {
			if !(d[i] >= '0' && d[i] <= '9') || !(d[i+1] >= '0' && d[i+1] <= '9') {
				return 0, errors.New("cisco time: must be in the format 00:00:00, found " + d)
			}
			val = val*60 + 10*(int(d[i])-'0') + int(d[i+1]-'0')
		}
		durationVal = Duration(val) * 1e9
	}
	return
}
