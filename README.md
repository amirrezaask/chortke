# Chortke
Chortke is a pure functional language for calculations only.
It's not a general purpose language.
It's mostly designed to solve AdventOfCode like challenges using pure functional style code.

# Syntax
Chortke is a expression based.

## Example
```
reduce + 1 2 3 4 5
15

3-2
1

2*3
6
reduce * 1 2 3 4 5 6
reduce {$acc, $x: $x*$acc} 1 2 3 4 5 6
720

map {x:x*2} 1 2 3 4 5
Map {$x:$x*2} 1 2 3 4 5

2 4 6 8 10

x = reduce + 1 2 3 4 5

```

## Formal Syntax
If line starts with a number it's an arithmitic
If line starts with a symbol name it's a call
only map and reduce are supported now

