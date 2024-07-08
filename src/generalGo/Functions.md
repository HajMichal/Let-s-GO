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