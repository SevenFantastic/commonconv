package weightconv

import "fmt"

type (
	Lb float64
	Kg float64
)

var (
	OneKG    Kg = 1
	TenKG    Kg = 10
	OneTonne Kg = 1000
)

func (k Kg) String() string { return fmt.Sprintf("%g Kg", k) }
func (l Lb) String() string { return fmt.Sprintf("%g Lb", l) }

func (k Kg) KToL() Lb { return Lb(k / 0.45) }
func (l Lb) LToK() Kg { return Kg(l * 0.45) }
