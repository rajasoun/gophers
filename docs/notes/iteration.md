# Iteration

For Loop is the single most loop used in Go Language and thus Go has kept the language simple by just implementing a single loop. It is a repetition structure that allows to write a loop i.e executed specific number of times.


## Syntax

```
for initialization; condition; post{
	// statements....
}
```

## Unit Test on Iteration 
> (Requirement: write a test for a function that repeats a character 5 times)
### Testing :
```
func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
```
### Unit Test-functionality :
```
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}
```
