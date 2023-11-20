package main 
import "fmt"
import "io"
import "os"
import "bufio"
import "strconv"

type Stack struct {
    values []DynValue
}

func (s *Stack) Push(value DynValue) {
    s.values = append(s.values, value)
}

func (s *Stack) Pop() DynValue {
    if len(s.values) == 0 {
        panic("Stack underflow")
    }
    value := s.values[len(s.values) - 1]
    s.values = s.values[:len(s.values) - 1]
    return value
}

func runWord(word string, stack *Stack) {
    switch word {
    case "print":
        val, err := stack.Pop().String()
        if err != nil {
            panic(err)
        }
        fmt.Println(val)
    case "ifeq":
        elseBlock := stack.Pop()
        ifBlock := stack.Pop()
        a, err := stack.Pop().Int()
        if err != nil {
            panic(err)
        }
        b, err := stack.Pop().Int()
        if err != nil {
            panic(err)
        }
        if a == b {
            stack.Push(ifBlock)
        } else {
            stack.Push(elseBlock)
        }
    }
}

func main() {
    file, err := os.Open("test.baz")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    stack := Stack {
        values: make([]DynValue, 0),
    }

    reader := bufio.NewReader(file)
    word := ""
    num := false
    float := false
    str := false
    for {
        char, _, err := reader.ReadRune()
        if err == io.EOF {
            break
        }
        switch char {
        case ' ', '\n', '\t', '\r':
            if float {
                val, err := strconv.ParseFloat(word, 64)
                if err != nil {
                    panic(err)
                }
                stack.Push(&DynFloat {val})
                float = false
                num = false
            } else if num {
                val, err := strconv.Atoi(word)
                if err != nil {
                    panic(err)
                }
                stack.Push(&DynInt {val})
                num = false
            }

            if str {
                word += string(char)
                continue
            }
            runWord(word, &stack)
            word = ""
            continue
        case '"', '\'':
            str = !str
            if !str {
                stack.Push(&DynString {word})
                word = ""
            }
            continue
        case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
            if str {
                word += string(char)
                continue
            }
            num = true
            word += string(char)
            continue
        case '.':
            if str {
                word += string(char)
                continue
            }
            float = true
            word += string(char)
            continue
        case '+':
            runAdd(&stack)
            continue
        case '-':
            runSubtract(&stack)
            continue
        case '*':
            runMultiply(&stack)
            continue
        case '/':
            runDivide(&stack)
            continue
        case '%':
            runModulo(&stack)
            continue
        case '^':
            runPower(&stack)
            continue
        default:
            word += string(char)
        }
    }
}
