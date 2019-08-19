# Integers

There are four distinct sizes of integers:

* Signed: `int8`, `int16`, `int32`, `int64`

* Unsigned: `uint8`, `uint16`, `uint32`, `uint64`

Where numeric value is a size in `bits`. 

The most used types would be `int` and `uint` where both of them would have either `32` or `64` bits, depending on the compiler type.

Signed numbers are repres ented in 2’s-complement form, in which the high-order bit is
reserved for the sign of the number and the range of values of an n-bit number is from `−2^(n−1)`
to `2^(n−1) −1`. Unsigned integers use the full range of bits for non-negative values and thus have
the range `0` to `2^n −1`. For instance, the range of int8 is −128 to 127, whereas the range of
uint8 is 0 to 255.

## Comparison operators

| operator | desc |
|----|----|
| `==` | equal to |
| `!=` | not equal to |
| `<`  | less than |
| `<=` | less than or equal to |
| `>`  | greater than |
| `>=` | greater than or equal to |

## Binary operators

| operator | desc |
|----|----|
| `&` | bitwise AND |
| `|` | bitwise OR |
| `^`  | bitwise XOR |
| `&^` | bit clear (AND NOT) |
| `<<`  | left shift |
| `>>` | right shift |

