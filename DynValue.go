package main

const (
    DynTypeString = iota
    DynTypeInt = iota
    DynTypeFloat = iota
)

type DynValue interface {
    Type() int
    String() (string, error)
    Int() (int, error)
    Float() (float64, error)
}

