package main
import "fmt"

type DynInt struct {
    value int
}

func (d *DynInt) Type() int {
    return DynTypeInt
}

func (d *DynInt) String() (string, error) {
    return fmt.Sprintf("%d", d.value), nil
}

func (d *DynInt) Int() (int, error) {
    return d.value, nil
}

func (d *DynInt) Float() (float64, error) {
    return float64(d.value), nil
}

