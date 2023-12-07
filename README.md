# [Advent of Code](https://en.wikipedia.org/wiki/Advent_of_Code)

Chapeau to Eric Wastl (http://was.tl/) (also for [vanilla js](http://vanilla-js.com/))

The aim is to - eventually - do all puzzles from all years. Much of the code is quick-and-dirty and I'm sometimes surprised to get the right answer. Occasionally I ask GPT-4 for help with a function; I've given up on Google's Bard, which has yet to yield any correct answers.

## Coding style

The first line of each solution should have a comment linking to the problem description page.

Solutions should have separate partOne() and partTwo() functions/procedures - no command line flags switching.

Favour readability, eg in Go, use `i += 1` rather than `i++`. Use self-written helper libraries that apply to all days. Only simple regular expressions.

Discourage screen clutter, eg in Go use a `switch` rather than a nested cascade of `if else if else`.

Use one logging/trace system, eg in Go use `log.Println`, in Lua use a simple `log` library.

## Lessons

1. Remove all irrelevant characters from input before processing it. (2023 day 1 part 1). Ignore all the fluff and just get the values (2023 day 5)
2. Use maps (2023 day 3)
3. Brute force can be done in reverse (2023 day 5)

## Go helper library

A Lodash-style Go library based on Go 1.18+ Generics (map, filter, contains, find...) https://github.com/samber/lo

## Observations on language

- Go's regexp package seems a little weird and clunky, eg specifying named capture groups seems obliquely-supported, unless I'm missing something. Lua offers a simpler, non-POSIX, system. (*Unlike several other scripting languages, **Lua does not use POSIX regular expressions (regexp) for pattern matching**. The main reason for this is size: A typical implementation of POSIX  regexp takes more than 4,000 lines of code. This is bigger than all Lua  standard libraries together*).

- Lua 5.1's lack of bitwise operators hurts a bit, unless I use 5.4 or [the library that comes with LuaJIT](https://bitop.luajit.org/), or something like:

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

- Go's verbosity grates a bit (but then it is a 'modern C'). When. for example, declaring a slice of three ints like `var dims []int = []int{0, 0, 0}`, why is the type `[]int` repeated? Why not just have `var dims []int = {0, 0, 0}`? Go is 'designed for concurrency and scalability' so it's not really a great fit for AoC.
- The most popular choice of language for AOC seems to be [Python](https://www.python.org/). However, I don't know Python and don't really want to. The nearest I'd be willing to go in that direction would be [Nim](https://nim-lang.org/). Other well-suited languages include [Perl](https://www.perl.org/), [Ruby](https://www.ruby-lang.org/en/), Dart, Tcl, ...
- AoC likes a language that is compact without being unreadable (so, not Javascript, SED, and certainly no code golf), not long winded (so, not Java or C and it's modern ilk), that has good out-of-the-box support for string manipulation, and can occasionally do MD5 hashes and bit-twiddling. It doesn't have to be fast. This [this blog entry](https://www.benkraft.org/2017/12/26/advent-of-code/) for a discussion of several languages (TLDR he didn't like Perl).

## Notable repos
- [Go](https://github.com/alexchao26/advent-of-code-go)
- [Nim](https://github.com/narimiran/advent_of_code_2015)
- [Racket](https://github.com/goderich/aoc2020/blob/master/day07.rkt)
- [Python](https://sharick.xyz/projects/advent-of-code)
- [Clojure](https://github.com/tschady/advent-of-code/tree/main)
- [F#](https://github.com/CameronAavik/AdventOfCode)
- [AWK](https://github.com/phillbush/aoc)

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

