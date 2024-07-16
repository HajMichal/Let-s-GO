# Structs
are custom types

```
type Circle struct {
    x float64
    y float64
    r float64
}

var c Circle
c := new(Circle)
```

If you want to assign values to a struct:
```
z := Circle{x: 5.0, y: 3.5, r: 4.0}
```
or if we want pointer to the struct
```
z := &Circle{x: 5.0, y: 3.5, r: 4.0}
```

we can access to a field by `z.x`, `z.r`