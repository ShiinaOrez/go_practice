package tempconv

import (
    "fmt"
)


type C float64
type F float64
type K float64

const (
    AbsoluteZeroC    C=-273.15
    FreezingC        C=0
    BoilingC         C=100
)

func (c C) String() string {
    return fmt.Sprintf("%g°C",c)
}

func (f F) String() string {
    return fmt.Sprintf("%g°F",f)
}

func (k K) String() string {
    return fmt.Sprintf("%g°K",k)
}