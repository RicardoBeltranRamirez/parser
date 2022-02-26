# Parser 

A tiny library that checks whether brackets are balanced or not. It supports ``()[]{}``, any other character is ignored by default. E.x.:

- ``({})`` -> true
- ``([a]())`` -> true
- ``([)]`` -> false

# How to use it

- Import in your project ``github.com/RicardoBeltranRamirez/parser``.
- Then, you can check with  ``parser.AreBracketsBalanced(val)`` where ``val`` is a string value.