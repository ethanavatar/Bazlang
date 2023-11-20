package main
import "errors"

func runAdd(stack *Stack) {
    a := stack.Pop()
    switch a.Type() {
    case DynTypeInt:
        b := stack.Pop()
        if b.Type() != DynTypeInt {
            panic(errors.New("Cannot add int to non-int"))
        }
        val, err := a.Int()
        if err != nil {
            panic(err)
        }

        result := b.(*DynInt).value + val
        stack.Push(&DynInt {result})
    case DynTypeFloat:
        b := stack.Pop()
        if b.Type() != DynTypeFloat {
            panic(errors.New("Cannot add float to non-float"))
        }
        val, err := a.Float()
        if err != nil {
            panic(err)
        }

        result := b.(*DynFloat).value + val
        stack.Push(&DynFloat {result})
    case DynTypeString:
        b := stack.Pop()
        if b.Type() != DynTypeString {
            panic(errors.New("Cannot add string to non-string"))
        }
        val, err := a.String()
        if err != nil {
            panic(err)
        }

        result := b.(*DynString).value + val
        stack.Push(&DynString {result})
    }
}

func runSubtract(stack *Stack) {
    a := stack.Pop()
    switch a.Type() {
    case DynTypeInt:
        b := stack.Pop()
        if b.Type() != DynTypeInt {
            panic(errors.New("Cannot subtract int by non-int"))
        }
        val, err := a.Int()
        if err != nil {
            panic(err)
        }

        result := b.(*DynInt).value - val
        stack.Push(&DynInt {result})
    case DynTypeFloat:
        b := stack.Pop()
        if b.Type() != DynTypeFloat {
            panic(errors.New("Cannot subtract float by non-float"))
        }
        val, err := a.Float()
        if err != nil {
            panic(err)
        }

        result := b.(*DynFloat).value - val
        stack.Push(&DynFloat {result})
    case DynTypeString:
        panic(errors.New("Cannot apply subtraction to string"))
    }
}

func runMultiply(stack *Stack) {
    a := stack.Pop()
    switch a.Type() {
    case DynTypeInt:
        b := stack.Pop()
        if b.Type() != DynTypeInt {
            panic(errors.New("Cannot multiply int by non-int"))
        }
        val, err := a.Int()
        if err != nil {
            panic(err)
        }

        result := b.(*DynInt).value * val
        stack.Push(&DynInt {result})
    case DynTypeFloat:
        b := stack.Pop()
        if b.Type() != DynTypeFloat {
            panic(errors.New("Cannot multiply float by non-float"))
        }
        val, err := a.Float()
        if err != nil {
            panic(err)
        }

        result := b.(*DynFloat).value * val
        stack.Push(&DynFloat {result})
    case DynTypeString:
        b := stack.Pop()
        if b.Type() != DynTypeInt {
            panic(errors.New("Cannot multiply string by non-int"))
        }
        val, err := a.Int()
        if err != nil {
            panic(err)
        }
        
        result := ""
        for i := 0; i < val; i++ {
            result += b.(*DynString).value
        }
        stack.Push(&DynString {result})
    }
}

func runDivide(stack *Stack) {
    a := stack.Pop()
    switch a.Type() {
    case DynTypeInt:
        b := stack.Pop()
        if b.Type() != DynTypeInt {
            panic(errors.New("Cannot divide int by non-int"))
        }
        val, err := a.Int()
        if err != nil {
            panic(err)
        }

        result := b.(*DynInt).value / val
        stack.Push(&DynInt {result})
    case DynTypeFloat:
        b := stack.Pop()
        if b.Type() != DynTypeFloat {
            panic(errors.New("Cannot divide float by non-float"))
        }
        val, err := a.Float()
        if err != nil {
            panic(err)
        }

        result := b.(*DynFloat).value / val
        stack.Push(&DynFloat {result})
    case DynTypeString:
        panic(errors.New("Cannot apply division by string"))
    }
}

func runModulo(stack *Stack) {
    a := stack.Pop()
    switch a.Type() {
    case DynTypeInt:
        b := stack.Pop()
        if b.Type() != DynTypeInt {
            panic(errors.New("Cannot modulo int by non-int"))
        }
        val, err := a.Int()
        if err != nil {
            panic(err)
        }

        result := b.(*DynInt).value % val
        stack.Push(&DynInt {result})
    case DynTypeFloat:
        panic(errors.New("Cannot apply modulo to float"))
    case DynTypeString:
        panic(errors.New("Cannot apply modulo to string"))
    }
}

func runPower(stack *Stack) {
    a := stack.Pop()
    switch a.Type() {
    case DynTypeInt:
        b := stack.Pop()
        if b.Type() != DynTypeInt {
            panic(errors.New("Cannot power int by non-int"))
        }
        val, err := a.Int()
        if err != nil {
            panic(err)
        }

        result := b.(*DynInt).value ^ val
        stack.Push(&DynInt {result})
    case DynTypeFloat:
        panic(errors.New("Cannot apply exponent to float"))
    case DynTypeString:
        panic(errors.New("Cannot apply exponent to string"))
    }
}

