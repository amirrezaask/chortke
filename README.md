# Chortke
Chortke in persian means calculator, chortke is a simple calculator written in Golang, I wanted to know more about infix notation parsing and I was amazed by Rob Pike's ivy calculator.

# Syntax
Chortke can evaluate most mathematical expressions.

# Internal
Chortke uses `Shunting Yard` algorithm to convert infix expressions into postfix ones and then simply evaluate them.
## Example
```
2*2+1/3-1
```
