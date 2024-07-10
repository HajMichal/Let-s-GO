# Functions

You can name return type like this:
```
func f2() (r int) {
    r = 1
    return
}
```

You can refer multiple types which will be returned:
```
func f2() (int, string) {
    return 5, "SomeString"
}
```

And then you can destructure:
```
    someNumber, randomString := f2()
```

We cam define variadic function like this:
```
func add(args ...int) int {}
```
... is ellipsis (zero or more)

And when you call this function you can pass many arguments
```
fmt.Println(add(1,2,3,4,5))
```

## Defer
`defer funcName()` -> schedules call to be run after the function completes

E.g usage:
```
f, _ := os.Open(filename)
defer f.Close()
```

## Panic, Recover

panic() is a runtime error which can be stopped by recover()
recover() stops the panic an returns the value that was passed to the call to panic
```
func main() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("Panic")
}
```

## Pointers
are arguemnts which can change variable value from the original function in sub-function
```
func zero(xPtr *int) {
    *xPtr = 0
}
func main() {
    someValue := 5
    zero(&someValue)
    fmt.Println(someValue) // output is 0
}
```

Asterisk(*) is representing pointer type and value in sub-function, 
Operator(&) is finding the address of a variable

Another way to get a pointer is to use built-in function named new 
Insted of 
```
    someValue := 5
    zero(&someValue)
```
We can:
```
    someValue := new(int)
    zero(someValue)
```
New takes a type as arguemtn