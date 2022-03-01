# Unit Test in Go

## The characteristic of a test in Go:
1. The first and only param must be `t *testing.T`
2. The testing function must begin with the word `Test` with a capital letter.
3. The test call `t.Error` or `t.Fail` to indicate a failure
4. You can use `t.Log` to provide non-failing debug info.
5. Naming convention: file - `foo_test.go`, `main_test.go`.

## Command
1. `go test -v`
- -v: verbose 

```Go
package math

import "testing"

func TestAdd(t *testing.T){

    got := Add(4, 6)
    want := 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
```


## Deadlock in Database
Because there is Key constrain between accounts and transfer table
so we need to add keyword: FOR NO KEY UPDATE; in statement sql to prevent deadlock  

add txName to the context: 			
`ctx := context.WithValue(context.Background(), txKey, txtName)`
get txName from the context:
`txName := ctx.Value(txKey)`
