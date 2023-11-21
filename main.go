package main 
import "fmt"
import "io"
import "os"
import "bufio"
import "strconv"
import "strings"

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

func (s *Stack) Peek() DynValue {
    if len(s.values) == 0 {
        panic("Stack underflow")
    }
    return s.values[len(s.values) - 1]
}

func runWord(word string, stack *Stack) {
    switch word {
    case "print":
        val, err := stack.Pop().String()
        if err != nil {
            panic(err)
        }
        fmt.Print(val)
    case "println":
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
    case "eval":
        count, err := stack.Pop().Int()
        if err != nil {
            panic(err)
        }
        if count < 0 {
            panic("Cannot eval negative number of times")
        }

        block, err := stack.Pop().String()
        if err != nil {
            panic(err)
        }

        for i := 0; i < count; i++ {
            reader := bufio.NewReader(strings.NewReader(block))
            runProgram(reader, stack)
        }
    case "dup":
        stack.Push(stack.Peek())
    }
}

func runProgram(reader *bufio.Reader, stack *Stack) {
    word := ""
    for {
        char, _, err := reader.ReadRune()
        if err == io.EOF {
            runWord(word, stack)
            break
        }
        switch char {
        case '#':
            for {
                char, _, err := reader.ReadRune()
                if err == io.EOF || char == '\n' {
                    break
                }
            }
            continue
        case ' ', '\n', '\t', '\r':
            runWord(word, stack)
            word = ""
            continue
        case '(':
            depth := 1
            for {
                char, _, err := reader.ReadRune()
                if err == io.EOF {
                    panic("Unmatched parenthesis")
                }
                if char == '(' {
                    depth++
                } else if char == ')' {
                    depth--
                    if depth == 0 {
                        break
                    }
                }
                word += string(char)
            }
            stack.Push(&DynString {word})
            continue
        case '"', '\'':
            for {
                char, _, err := reader.ReadRune()
                if err == io.EOF {
                    panic("Unmatched string delimiter")
                }
                if char == '"' || char == '\'' {
                    break
                }
                word += string(char)
            }
            stack.Push(&DynString {word})
            continue
        case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
            word += string(char)
            float := false
            for {
                char, _, err := reader.ReadRune()
                if err == io.EOF {
                    break
                }
                if char == '.' {
                    float = true
                    word += string(char)
                    continue
                }
                if char < '0' || char > '9' {
                    reader.UnreadRune()
                    break
                }
                word += string(char)
            }
            if float {
                value, err := strconv.ParseFloat(word, 64)
                if err != nil {
                    panic(err)
                }
                stack.Push(&DynFloat {value})
            } else {
                value, err := strconv.ParseInt(word, 10, 64)
                if err != nil {
                    panic(err)
                }
                stack.Push(&DynInt {int(value)})
            }
            continue
        case '+':
            runAdd(stack)
            continue
        case '-':
            runSubtract(stack)
            continue
        case '*':
            runMultiply(stack)
            continue
        case '/':
            runDivide(stack)
            continue
        case '%':
            runModulo(stack)
            continue
        case '^':
            runPower(stack)
            continue
        default:
            word += string(char)
        }
    }
}

func main() {

    if len(os.Args) < 2 {
        fmt.Println("Usage: baz <scriptPath>")
        return
    }

    program := os.Args[1]

    file, err := os.Open(program)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    stack := Stack {
        values: make([]DynValue, 0),
    }

    reader := bufio.NewReader(file)
    runProgram(reader, &stack)
}
