# [Advent of Code](https://en.wikipedia.org/wiki/Advent_of_Code)

Chapeau to Eric Wastl (http://was.tl/) (also for [vanilla js](http://vanilla-js.com/))

The aim is to - eventually - do all puzzles from all years. Much of the code is quick-and-dirty and I'm sometimes surprised to get the right answer. Occasionally I ask GPT-4 for help with a function; I've given up on Google's Bard, which has yet to yield any correct answers.

## 2023
Follow along day-by-day using Go.

Then go back to 2015 onwards finding solutions in other languages for each year, like Nim, Dart, Tcl, Ruby? ...

## Coding style

The first line of each solution should have a comment linking to the problem description page.

Solutions should have separate partOne() and partTwo() functions/procedures - no command line flags switching.

Favour readability, eg in Go, use `i += 1` rather than `i++`. Use Set and Stack libraries.

Discourage screen clutter, eg in Go use a `switch` rather than a nested cascade of `if else if else`.

Use one logging/trace system, eg in Go use `log.Println`, in Lua use a simple `log` library.

## Observations on language

- Go's regexp package seems a little weird and clunky, eg specifying named capture groups seems obliquely-supported, unless I'm missing something. Lua offers a simpler system. Why oh why oh why can't all languages just use one agreed regexp package?

- Lua 5.1's lack of bitwise operators hurts a bit, unless I use [the library that comes with LuaJIT](https://bitop.luajit.org/) or something like:

```Lua
OR, XOR, AND = 1, 3, 4

function bitoper(a, b, oper)
   local r, m, s = 0, 2^31
   repeat
      s,a,b = a+b+m, a%m, b%m
      r,m = r + m*oper%(s-a-b), m/2
   until m < 1
   return r
end

print(bitoper(6,3,OR))   --> 7
print(bitoper(6,3,XOR))  --> 5
print(bitoper(6,3,AND))  --> 2
```

- Go's verbosity grates a bit. When. for example, declaring a slice of three ints like `var dims []int = []int{0, 0, 0}`, why is the type `[]int` repeated? Why not just have `var dims []int = {0, 0, 0}`?

## Notable repos
- [Go](https://github.com/alexchao26/advent-of-code-go)
- [Nim](https://github.com/narimiran/advent_of_code_2015)
- [Racket](https://github.com/goderich/aoc2020/blob/master/day07.rkt)
- [Python](https://sharick.xyz/projects/advent-of-code)
- [Clojure](https://github.com/tschady/advent-of-code/tree/main)

## TODO
```go
package main

import (
    "fmt"
    "regexp"
)

var myExp = regexp.MustCompile(`(?P<first>\d+)\.(\d+).(?P<second>\d+)`)

func main() {
    match := myExp.FindStringSubmatch("1234.5678.9")
    result := make(map[string]string)
    for i, name := range myExp.SubexpNames() {
        if i != 0 && name != "" {
            result[name] = match[i]
        }
    }
    fmt.Printf("by name: %s %s\n", result["first"], result["second"])
}
```

