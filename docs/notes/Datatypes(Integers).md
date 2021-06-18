## DataType in Go:

In Go language, the type is divided into four categories which are as follows:
1. **Basic type**: Numbers, strings, and booleans.
1. **Aggregate type**: Array and structs.
1. **Reference type**: Pointers, slices, maps, functions, and channels.
1. **Interface type**

And In Go language, Numbers are divided into three sub-categories that are:
1. Integer
1. Float
1. Complex Number

## Lets see how to write unit test on adding function with 2 values(Integer) :

### Write unit test :
```
func TestAdder(t *testing.T) {
    sum := Add(2, 2)
    expected := 4

    if sum != expected {
        t.Errorf("expected '%d' but got '%d'", expected, sum)
    }
}
```

### Functionality of Add function:
```
func Add(x, y int) int {
    return x + y
}
```
