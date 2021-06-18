## Pointers in GO
 Pointers in Go programming language is a variable which is used to store the memory address of another variable.The memory address is always found in hexadecimal format.

 ### Declaration and Initialization of Pointers
 1. "*" Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
      ```
      " var pointer_name *Data_Type"
      ```

1. "&" operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.

e.g: 
```
func main() {
    var creature string = "shark"
    var pointer *string = &creature

    fmt.Println("creature =", creature)
    fmt.Println("pointer =", pointer)

    fmt.Println("*pointer =", *pointer)
}
//Output
creature = shark
pointer = 0xc000010200
*pointer = shark
```

### Program with TDD
Requirement: Write a ATM program (deposit and withdraw function) with unit test

Program Functionality Code:
```
    type Bitcoin int
    type Wallet struct {
        balance Bitcoin
    }

    func (w *Wallet) Deposit(amount Bitcoin) {
        w.balance += amount
    }

    func (w *Wallet) Balance() Bitcoin {
        return w.balance
    }
```

Unit Testing Code :
```
    func TestWallet(t *testing.T) {

    wallet := Wallet{}

    wallet.Deposit(Bitcoin(10))

    got := wallet.Balance()

    want := Bitcoin(10)

    if got != want {
        t.Errorf("got %d want %d", got, want)
    }
}
```

## Errors
Errors are the way to signify failure when calling a function/method.

UnChecked Errors:
There is one scenario we have not tested. To find it, run the following in a terminal to install errcheck

> go get -u github.com/kisielk/errcheck
