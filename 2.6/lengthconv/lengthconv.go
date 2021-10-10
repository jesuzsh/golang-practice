package lengthconv

import "fmt"

type Foot float64
type Meter float64

func (f Foot) String() string  { return fmt.Sprintf("%g Feet", f) }
func (m Meter) String() string { return fmt.Sprintf("%g Meter(s)", m) }
