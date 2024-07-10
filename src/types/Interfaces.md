# Interfaces
defines a set of methods that must be implemented (and their returning type)

```
type Shape interface {
    area() float64
}
```

We can do something like this:
```
func totalArea(circles []Circle, rects []Reactangle) flaote64 {
    var total float64
    for_, c := range circles {
        total := c.area()
    }
    for_, r := range rects {
        total := r.area()
    }
    return total
}
```

But in this case better will be interface:

```
func totalArea(shapes ...Shape) float64 {
    var total float64
    for_, s := range shapes {
        total := s.area()
    }
    return total
}
```

Invoke this func with:
```
fmt.Println(totalArea(&circles, &rects))
```
Pointers next to receivers name are important!