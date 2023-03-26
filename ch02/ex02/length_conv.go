package main

import "fmt"

type Shaku float64
type Meter float64

func (s Shaku) String() string { return fmt.Sprintf("%.3g尺", s) }
func (m Meter) String() string { return fmt.Sprintf("%.3gメートル", m) }

func ShakuToMeter(s Shaku) Meter {
	return Meter(s * 3.3)
}

func MeterToShaku(m Meter) Shaku {
	return Shaku(m / 3.3)
}
