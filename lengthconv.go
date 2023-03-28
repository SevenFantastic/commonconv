package commonconv

import "fmt"

type (
	Foot  float64
	Meter float64
)

var (
	ZeroMeterF Meter = 0
	OneKMF     Meter = 1000
	TenKMF     Meter = 10000
)

func (i Foot) String() string  { return fmt.Sprintf("%g foot", i) }
func (m Meter) String() string { return fmt.Sprintf("%g meter", m) }

func (m Meter) MToF() Foot { return Foot(m / 0.3048) }
func (f Foot) FToM() Meter { return Meter(f * 0.3048) }
