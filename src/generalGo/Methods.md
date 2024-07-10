# Methods
Those are something like a functions of specific elements

For example we've got a struct:

```
type Rectangle struct {
    x, y float64
}
```

Method using this struct is:
```
func (rect *Rectangle) area() float64 {
    return rect.x * rect.y
}
```
`rect *Rectangle` is a receiver

We don't have to use * before using pointer value from struct in methods
Go do it by itself

And we want to invoke the method
`rect.area()`
E.g:
`fmt.Println(rect.area())`

We can define struct with struct inside:
```
type Person struct {
    Name string
}
type Android struct {
    Person
    Model string
}
func(p *Person) Talk() {
    fmt.Println("Say something")
}
```

and then we can:
```
android := new(Android)
android.Person.talk()
```
but we can allso invoke talk method instantly after android
`android.talk()`


We can assign a method to a specifc structs.
It helps keeps the code clean


## This topic is simmilar to :
-[Interfaces](../types/Interfaces.md) <br/>
-[Structs](../types/Structs.md)


