# Slices

Slices is something like array. The only difference between this and an array is the missing length between the brackets.

Slices are always associeted with some array by: 

`x := make([]float64, 5)`

`make` is a keyword here which creates a slice.

Slice can't be longer but can be smaller than associeted array.

This example creates a slice that is associeted with an underlying float64 array of length 5

We can add third parameter to make function:

`x := make([]float64, 5, 10)`

and number 10 represents capactiy of the underlying array.
So we should read this as: 
"A slice of length 5 with a capacity of 10"


Another way to create slice is:

`````
arr := [5]float64{1,2,3,4,5}
x := arr[1,3] // x returns [2,3]
``````

`x := arr[3:]` is the same as `x := arr[3:5]`

`x := arr[:3]` is the same as `x := arr[0:3]`

There are 2 built-in functions to assist with slices:
* append `slice2 := append(slice1, 4, 5)`
* copy `copy(slice1, slice2)` `// copy(dst, src)` items from slice2 (src) are copied into slice1 (dst) overwriting whatever is there.
If the length of slices are different the smaller one is taken 