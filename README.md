# Bazlang

This is a toy language written in Go. It is a pure concatenative stack-based language.

## Examples

Values are pushed onto the stack. Operators pop values off the stack and push the result.

The `println` operation pops and prints the top value on the stack.
```
"Hello, Sailor!" println
```

The above program will print:

```
Hello, Sailor!
```

The `+` operator pops two values off the stack, adds them, and pushes the result.
Unlike other postfix languages, mathmatical operators use values in the order they are written. So the following results in the same as the above:
```python
"Hello, " "Sailor!" + println
```

There is basic conditional logic using the `ifeq` operator. It takes 4 arguments that are as follows:
```python
<a> <b> <true> <false> ifeq
```
If `<a>` and `<b>` are equal, then the value of `<true>` is pushed. Otherwise, the value of `<false>` is pushed. For example:
```python
1 1 "equal" "not" ifeq # the top stack value is "equal"
1 2 "equal" "not" ifeq # the top stack value is "not"
```

There is basic looping capability using the `eval` operator. It takes 2 arguments that are as follows:
```python
<code: string> <count: int> eval
```
This can be used like so:
```python
0 "1 +" 10 eval # the top stack value is 10
```

The `eval` operator is especially useful when combined with the `dup` operator, that simply duplicates the top stack value:
```python
# prints and makes a stack of acending numbers 1 -> 9
0 "dup println dup 1 +" 10 eval
```

When writing long running `eval` statements, using quoted strings can get a little hard to navigate. In this case, parentheses can be used, which are effectively just multiline strings.
```python
# prints and makes a stack of acending numbers 1 -> 9
0 (
    # prints the current top
    dup println
    
    # push top + 1
    dup 1 +

# loop 10 times
) 10 eval
```

As a final example, here is, roughly, fizzbuzz implemented using the above features:
```python
## fizzbuzz from 0 -> 15

# Create a stack of numbers from 0 -> 15
15 (dup 1 -) 15 eval

(
    # duplicate the current number
    # because it needs to be checked twice (for fizz and buzz)
    dup
    # print the current number followed by a space
    dup print " " print

    3 % 0 "fizz" "" ifeq print
    5 % 0 "buzz" "" ifeq println

# loop 31 times
) 16 eval
```

The above program will print:
```
0 fizzbuzz
1
2
3 fizz
4
5 buzz
6 fizz
7
8
9 fizz
10 buzz
11
12 fizz
13
14
15 fizzbuzz
```
