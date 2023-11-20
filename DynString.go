package main
import "errors"

type DynString struct {
    value string
}

func (d *DynString) Type() int {
    return DynTypeString
}

func (d *DynString) String() (string, error) {
    return d.value, nil
}

func (d *DynString) Int() (int, error) {
    return 0, errors.New("Cannot convert string to int")
}

func (d *DynString) Float() (float64, error) {
    return 0, errors.New("Cannot convert string to float")
}

