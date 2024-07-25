# Maps

unordered collection of key-value pairs

### Define map:
```
x := make(map[KEY-TYPE]VALUE-TYPE)
x["key"] = 10
fmt.Println(x["key"])    
```

If provided key not match with any map element the program return nothing (there is no error)


You can use syntax like this to check if there is an element:
```
name, ok := elementsMap["someKey"]
```
name represents value of element (if exists)
ok represents is lookup was succesfull

You can use shorter syntax to define maps

```
elements := map[string]string {
    "H": "Hydrogen",
    "He": "Helium",
    "C": "Carbon",
    "O": "Oxygen",
}
```

If we need to have nestes maps we can do something like this:

```
elements := map[string]map[string]string{
    "H": map[string]string{
        "name": "Hydrogen",
        "state": "gas",
    },
    "He": map[string]string{
        "name": "Helium",
        "state": "gas",
    },
    "C": map[string]string{
        "name": "Carbon",
        "state": "solid",
    },
}
```

And we can chceck if there is a Carbon: 

```
if el, ok := elements["C"]; ok {
    fmt.Println(el["name"], el["state"])
}
```

To remove key value pair from map:
`delete(nameOfMap, keyYouWantToRemove)`


If we are iterating through map we can:
```
for keyName, value := range nameOfMap {

}
```

Differences between maps and structs
 * Maps:
    - must have all keys with the same type and all values with the same types
    - are idexable
    - when we pass map somewhere we pass copy reference to this map (not strict map)

 * Structs:
    - types can be mixed
    - NOT idexable
    - when we pass structs somewhere we pass actuall copy of struct

