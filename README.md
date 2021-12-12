# Chortke
Chortke is a pure functional language for calculations only.
It's not a general purpose language.
It's mostly designed to solve AdventOfCode like challenges using pure functional style code.

# Syntax
Chortke is a expression based.

## Example
```
let inc { x -> x + 1 }

inc 1 ;; 2

map [1 2 3] {x -> x*2 } ;; [ 2 4 6 ]

reduce [1 2 3] { acc x -> acc + x } ;; 6

```