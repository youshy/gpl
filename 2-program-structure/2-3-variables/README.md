# Variables

## General syntax

```golang
var name type = expression
```

In `go`, `type` or `= expression` can be ommited, but not both. If the `expression` is ommited, then the initial value is set to the `zero value` for the type, which is:

* `0` for `numbers`

* `false` for `booleans`

* `""` for `strings`

* `nil` for `interfaces` and reference types (`slice`, `pointer`, `map`, `channel`, `function`)

It is possible to declare and optionally initialize a set of variables in a single declaration:

```golang
var i, j, k int         // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string
```

A set of variables can also be initialized by calling a function that returns multiple values:

```golang
var f, err = os.Open(name) // returns a file and an error
```

## Short Variable Declarations

```golang
name := expression
```

Automatically assigns the type:

```
i, j := 100, 32 // both are int (use only if that helps readability)
var boiling float64 = 100 // float64
```

There's a big difference between `:=` and `=` - first one is a declaration, second is an assignment.

```golang
i, j = j, i // this swaps values of i and j
```


Same as `var`, `:=` declaration may be used for calls to functions:

```golang
f, err := os.Open(name)
if err != nil { // most used
  return err
}
```

A short variable declaration must declare at least one new variable:

```golang
// bad
f, err := os.Open(infile)

f, err := os.Create(outfile) // compile error: no new variables
```

```golang
// good
in, err := os.Open(infile)

out, err := os.Create(outfile) // err will be reassigned to already existing variable
```

## Pointers

Pointer is the address of the variable. It is the location at which a value is stored. Not every value has an address, but every variable does. Pointers allows us reading or updating the value of a variable indirectly, without using or even knowing the name of the variable (if it even has a name).

If a variable is declared var x int, the expression `&x` (‘address of x’) yields a pointer to an integer variable, that is, a value of type `*int`, which is pronounced ‘pointer to int.’

If this value is called `p`, we say ‘p points to x,’ or equivalently ‘p contains the address of x.’ The variable to which `p` points is written `*p`. The expression `*p` yields the value of that variable, an `int`, but since `*p` denotes a variable, it may also appear on the left-hand side of an assignment, in which case the assignment updates the variable:

```golang
x := 1
p := &x           // p, of type *int, points to x
fmt.Println(*p)   // "1"
*p = 2            // equivalent to x = 2
fmt.Println(x)    // "2"
```

Zero vaue for `pointer` of any type is `nil`. Test `p != nil` will be only true if `p` points to a variable. Pointers are comparable - two pointers are equal only if they point to the same variable or both are `nil`:

```golang
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
```

It is safe for a function to return the address of a local variable. In case below, `v` created by the call to `f` will remain in existence even after the call has returned, and `p` will still refer to it:

```golang
var p = f()

func f() *int {
  v := 1
  return &v
}
```

Each call will return a distict value:

```golang
fmt.Println(f() == f()) // false
```

Because a pointer containts the address of a variable, passing a pointer argument to a function gives the possibility to update the variable from a function:

```golang
func incr(p *int) int {
  *p++ // increments what p points to, but does not change it
  return *p
}

v := 1
incr(&v)                // v is now 2
fmt.Println(incr(&v))   // "3" (and v is 3 now)
```
