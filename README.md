# [Advent of Code](https://en.wikipedia.org/wiki/Advent_of_Code)

Chapeau to Eric Wastl (http://was.tl/) (also for [vanilla js](http://vanilla-js.com/))

The aim is to - eventually - do all puzzles from all years. Some of the code is quick-and-dirty and I'm sometimes surprised to get the right answer. Occasionally I ask GPT-4 for help with a low-level function or algorithm in case I've missed something simple; I've given up on Google's Bard, which has yet to yield any correct answers. Occasionally, using when a brute force solution is taking too long, I have a look at the solutions mega thread on reddit to get ideas. 

## Coding style

The first line of each solution should have a comment linking to the problem description page.

Solutions should have separate partOne() and partTwo() functions/procedures - no command line flags switching.

Favour readability over being clever, for example in Go, use `i += 1` rather than `i++`. Code golfing is an ingenious art form, but not for me.

Discourage screen clutter, for example in Go use a `switch` rather than a nested cascade of `if else if else`. (Except that, the Go debugger gets a bit jumpy around switch statements.)

Keep all the code self-contained; no helper functions or libraries unless there is a very good reason.

Use one logging/trace system, for example in Go use `log.Println`, in Lua use a simple `log` library.

## Lessons

1. Remove all irrelevant characters from input before processing it. (2023 day 1 part 1). We mindful of any trailing newlines. Ignore all the fluff and just get the values (2023 day 5)
2. Use maps (2023 day 3) with numeric keys (Lua), especially when key is multi-part.
3. Brute force can be done in reverse (2023 day 5)
4. Avoid regular expressions, especially in Go. Like someone said: "if you solve a problem with regular expressions, you now have two problems". `fmt.Sscanf()` can go a long way.
5. Seeing the word 'infinite' in the description means you should be using maps, not a fixed size grid.

## Observations on language

- Go's regexp package seems a little weird and clunky, eg specifying named capture groups seems obliquely-supported, unless I'm missing something. Lua offers a simpler, non-POSIX, system. (*Unlike several other scripting languages, **Lua does not use POSIX regular expressions (regexp) for pattern matching**. The main reason for this is size: A typical implementation of POSIX  regexp takes more than 4,000 lines of code. This is bigger than all Lua  standard libraries together*).

- In Go, it's often not clear (to me) if I'm dealing with: (1) the object, (2) a copy of the object, (3) a pointer to the object.

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

- Go's verbosity grates a bit (but then it is a 'modern C'). When. for example, declaring a slice of three ints like `var dims []int = []int{0, 0, 0}`, why is the type `[]int` repeated? Why not just have `var dims []int = {0, 0, 0}`? Go is 'designed for concurrency and scalability' so it's not an ideal fit for AoC, but it's great at just getting the job done.
- The most popular choice of language for AoC seems to be [Python](https://www.python.org/). However, I don't know Python and don't really want to. The nearest I'd be willing to go in that direction would be [Nim](https://nim-lang.org/). Other well-suited languages include [Perl](https://www.perl.org/), [Ruby](https://www.ruby-lang.org/en/), Dart, Tcl, ...
- AoC likes a language that is compact without being unreadable (so, not Javascript, SED, and certainly no code golf), not long winded (so, not Fortran, Java or C and it's modern ilk), that has good out-of-the-box support for string manipulation, and can occasionally do MD5 hashes and bit-twiddling. It doesn't have to be fast. This [this blog entry](https://www.benkraft.org/2017/12/26/advent-of-code/) for a discussion of several languages (TLDR: he didn't like Perl).

## Notable repos
- [Go](https://github.com/alexchao26/advent-of-code-go)
- [Nim](https://github.com/narimiran/advent_of_code_2015)
- [Racket](https://github.com/goderich/aoc2020/blob/master/day07.rkt)
- [Python](https://sharick.xyz/projects/advent-of-code)
- [Clojure](https://github.com/tschady/advent-of-code/tree/main)
- [F#](https://github.com/CameronAavik/AdventOfCode)
- [AWK](https://github.com/phillbush/aoc)

## Notable puzzles

| Year | Day  | Notes                                                        |
| ---- | ---- | ------------------------------------------------------------ |
| 2018 | 15   | Works on all the test cases but not the actual input. Have respected all the advice from others regarding sort order. The answer depends on the order of direction looks in the the BFS function, which is a big clue to where I'm going wrong, but I can't decipher it. |
| 2019 | 7    | Part 2 - instructions unclear - turned out, instruction pointer needed to be saved between invocations of each amplifier, not just the []int instructions/memory. |
| 2019 | 10   | Proud of part 1 solution; did it all in integers without calculating any angles or creating an actual grid. Solution uses a map to hold the asteroid positions, and calculates if any asteroids lie between two other asteroids using an algorithm copied from stack overflow. It's a bit O(n3), and takes nearly a second on my machine, so no prizes for efficiency.  Part 2 does calculate angles, but only when sorting a slice of asteroids visible from the laser point. |

## Comments from u/vipul0092

Some highlights of the journey:

- 2019 is by far my favourite year, because of IntCode
- 2018 is definitely the most difficult one
- 2017 is definitely the easiest one

And finally shoutout to the nastiest puzzles that bit me for a long time (But every one of them taught me something):

2015 - Day 19 (CFG and CYK... what?)

2016 - Day 11 (Solving by hand is easier...DFS + State pruning)

2018 - Day 15 (Still don't know if it'll pass all the inputs)

2018 - Day 17 (Took me an year to solve, in multiple attempts )

2018 - Day 23 (Octrees or something, I dunno)

2019 - Day 18 (Dynamic programming at its very best, seriously tho, if you think you know DP, try this problem)

2019 - Day 22 (Become a mathematician)

2020 - Day 20 (Low level OpenCV implementation??)

2021 - Day 22 (Nothing to say about this one, was totally blank as to how to approach here)

2021 - Day 23 (Solving by hand is easier, for a coded solution tho...)

2022 - Day 19 (DFS + State pruning...ugh)

2023 - Day 21 (Keep counting until the world ends?)
