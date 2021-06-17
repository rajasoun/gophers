## Dependency Injection

In software engineering,Dependency injection (DI) is the process of providing external dependency to a software component.
It is a technique in which an object receives other objects that it depends on. These other objects are called dependencies. And the "injection" refers to the passing of a dependency (a service) into the object (a client) that would use it.

### Purpose
1. It allows a client to remove all knowledge of a concrete implementation that it needs to use. This helps isolate the client from the impact of design changes and defects. 
1. Also,It promotes reusability, testability and maintainability.


## Code with Unit Testing in Go
1. Requirements:  Write Program with testing (Greets someone , also going to be testing the actual printing)
 a. If use "fmt.Printf" for prints to stdout ,so it is pretty hard to capture using the testing framework.
 b. So, we need to do is to "inject" the dependency of printing.
 c. Then,we can change the implementation to print to something.so we can test it. 

 Note: In "real life" you would inject in something that writes to stdout.

 > for printing ,we use "Printf" ,this fmt package func returns "Fprintf" and "Fprintf" passing in "os.Stdout".

 
 ```
 func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
    p := newPrinter()
    p.doPrintf(format, a)
    n, err = w.Write(p.buf)
    p.free()
    return
}
 ```

 And for io.Writer interface (using Writer to send our greeting somewhere.)
 ```
 type Writer interface {
    Write(p []byte) (n int, err error)
}
 ```
UnitTest 
```
func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Chris")

    got := buffer.String()
    want := "Hello, Chris"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```
Functionality :
 ```
 func Greet(writer io.Writer, name string) {
    fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
    Greet(os.Stdout, "Elodie")
}
 ```