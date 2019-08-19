# Assignments

```golang
x = 1                         // named variable
*p = true                     // indirect variable
person.name = "bob"           // struct field
count[x] = count[x] * scale   // array or slice or map element
```

## Tuples

Aka - `Assign more than one at once`

```golang
// swapping values

x, y = y, x
a[i], a[j] = a[j], a[i]

// Computing greates common divisor of two integers
func gcd(x, y int) int {
  for y != 0 {
    x, y = y, x%y
  }
  return x
}

v, ok = m[key]              // map lookup
v, ok = x.(T)               // type assertion
v, ok = <-ch                // channel receive

_, err = io.Copy(dst, src)  // discard byte count
_, ok = x.(T)               // check type but discard result
```

## Assignability

```golang
// implicit 
medals := []string{"gold", "silver", "bronze"}

// each element
medals[0] = "gold"
medals[1] = "silver"
medals[2] = "bronze"
```
