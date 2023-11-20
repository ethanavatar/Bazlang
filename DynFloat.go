package main
import "fmt"

type DynFloat struct {
    value float64
}

func (d *DynFloat) Type() int {
    return DynTypeFloat
}

func (d *DynFloat) String() (string, error) {
    return fmt.Sprintf("%f", d.value), nil
}

func (d *DynFloat) Int() (int, error) {
    return int(d.value), nil
}

func (d *DynFloat) Float() (float64, error) {
    return d.value, nil
}

