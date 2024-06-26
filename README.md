# [Advent of Code](https://en.wikipedia.org/wiki/Advent_of_Code)

Chapeau to Eric Wastl (http://was.tl/) (also for [vanilla js](http://vanilla-js.com/))

The aim is to - eventually - do all puzzles from all years. Some of the code is quick-and-dirty and I'm sometimes surprised to get the right answer. Occasionally I ask GPT-4 for help with a low-level function or algorithm in case I've missed something simple; I've given up on Google's Bard, which has yet to yield any correct answers. When I get hopelessly stuck or a brute force solution is taking too long, I have a look at the solutions mega thread on reddit to get ideas. 

"every problem has a solution that completes in at most 15 seconds on ten-year-old hardware"

## Coding style

The first line of each solution should have a comment linking to the problem description page.

Solutions should have separate partOne() and partTwo() functions/procedures - no command line flags switching.

Favour readability over being clever, for example in Go, use `i += 1` rather than `i++`. Code golfing is an ingenious art form, but not for me.

Discourage screen clutter, for example in Go use a `switch` rather than a nested cascade of `if else if else`. (Except that, the Go debugger gets a bit jumpy around switch statements.)

Keep all the code self-contained; no helper functions or libraries unless there is a very good reason (like MD5 for Lua).

Use one logging/trace system, for example in Go use `log.Println`, in Lua use a simple `log` library.

## Lessons

1. Remove all irrelevant characters from input before processing it. (2023 day 1 part 1). Be mindful of any trailing newlines. Ignore all the fluff and just get the values (2023 day 5).
2. Use maps (2023 day 3) with numeric keys (Lua), especially when key is multi-part.
3. Brute force can be done in reverse (2023 day 5).
4. Avoid regular expressions, especially in Go. Like someone said: "if you solve a problem with regular expressions, you now have two problems". `fmt.Sscanf()` , `strings.Split()` and `strings.Fields()` can go a long way.
5. Seeing the word 'infinite' in the description means you should be using maps, not a fixed size grid.
6. A lot of stumbles and bad starts happen because I don't fully understand the puzzle descriptions, or think I can see ambiguities. When this happens, it's often helpful to work out a solution by hand in a text editor. At first, I laughed when I saw someone doing AoC in vim, but now I understand that completely.
7. Visualisation can be key; in a lot of puzzles being able to see the data/solution as it progresses can be very helpful when tracking down unexpected behaviour. 2018 day 17 (the one where water cascades down from a spring) is a key example.

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
- Getting hung up on the use of colons in Go type definitions: you use them when filling in a struct `{x: 32, y:74}` and in other languages (eg Pascal) when declaring variables and function parameters, but just really noticed that Go quietly drops the latter, shouldn't it be `func dostuff(thing: int)` rather than just `func dostuff(thing int)`? I think the colon makes it clearer.
- The most popular choice of language for AoC seems to be [Python](https://www.python.org/). However, I don't know Python and don't really want to. The nearest I'd be willing to go in that direction would be [Nim](https://nim-lang.org/). Other well-suited languages include [Perl](https://www.perl.org/), [Ruby](https://www.ruby-lang.org/en/), Dart, Tcl, ...
- AoC likes a language that is compact without being unreadable (so, SED, and certainly no code golf), not long winded (so, not Fortran, Java or C and it's modern ilk), that has good out-of-the-box support for string manipulation, and can occasionally do MD5 hashes and bit-twiddling. It doesn't have to be fast. This [this blog entry](https://www.benkraft.org/2017/12/26/advent-of-code/) for a discussion of several languages (TLDR: he didn't like Perl).
- You can write very readable vanilla JavaScript/ECMAScript/TypeScript and use the Bun or [Deno runtime](https://deno.com/).

## Notable resources

[The Competitive Programming Handbook](https://cses.fi/book/book.pdf) [with roadmap](https://www.geeksforgeeks.org/competitive-programming-cp-handbook-with-complete-roadmap/)

## Notable repos/commentaries/blogs
- [Michael Fogleman](https://www.michaelfogleman.com/aoc18/#17)
- [Go](https://github.com/alexchao26/advent-of-code-go)
- [Go 2023, 2022, 2021, 2020, 2019](https://github.com/lynerist?tab=repositories)
- [Go](https://github.com/xorkevin?tab=repositories)
- [Nim](https://github.com/narimiran/advent_of_code_2015)
- [Racket](https://github.com/goderich/aoc2020/blob/master/day07.rkt)
- [Python](https://sharick.xyz/projects/advent-of-code)
- [Clojure](https://github.com/tschady/advent-of-code/tree/main)
- [F#](https://github.com/CameronAavik/AdventOfCode)
- [AWK](https://github.com/phillbush/aoc)
- [Lua 2015, 2019, 2023](https://github.com/DeybisMelendez/AdventOfCode/tree/master)
- [Clear and *BLAZINGLY FAST* solutions in JS](https://github.com/JoanaBLate/advent-of-code-js/tree/main)
- https://www.ericburden.work/blog/
- [Go commentary 2019](https://dhconnelly.com/advent-of-code-2019-commentary.html)
- [Awesome bit of Go for 2019 day 15](https://github.com/mnml/aoc/blob/main/2019/15/2.go), some very succinct and clever code in this [https://github.com/mnml/aoc](repo) by mnml (who is this masked person?)
- [Nice bit of Go for 2019 day 19](https://github.com/stevotvr/adventofcode2019/blob/master/day19/day19.go)


## Notable puzzles

| Year | Day  | Notes                                                        |
| ---- | ---- | ------------------------------------------------------------ |
| 2018 | 15   | Part one works on all the test cases but not the actual input. Have respected all the advice from others regarding sort order. The computed result depends on the order of directions in the the BFS function, which is a big clue to where I'm going wrong, but I can't decipher it. Finally got correct answers after the third rewrite, total billable hours approx. 40. |
| 2018 | 16   | Great puzzle, whittling down occurrence lists.               |
| 2018 | 17   | Tricky. Got it on the third attempt after two days, after two false starts (1) by running a stepped simulation on the water blocks (nearly had it, but they spilled in unexpected places), then (2) by trying a DFS queue-style solution, and finally (3) a doubtful-looking but finally successful and quick recursive approach. Much relief that part two was already solved by part one. Also, tried several solutions from other participants from the solutions subreddit: only one worked on my input, the rest got stuck in some nested loop (I can barely read Python, let alone debug it). |
| 2018 | 19   | Part one just borrows the guts of the assembly language interpreter from day 16. Part two was one of those "run forever" problems, so study the registers at intervals; reg[4] contains a large number that does not change, and reg[5] contains an increasing count of something. From previous experience with AoC you can guess that it's the sum of all factors (integer divisors) of reg[4], so compute that directly. Feels like cheating. |
| 2018 | 20   | Seemed obvious from the start that you should use a stack, push when finding a ( [remember this place], peek when finding a \| [what if] and pop when finding a ) [okay, back to where we were]. Only afterwards I [read](https://www.reddit.com/r/adventofcode/comments/a7w4dj/2018_day_20_why_does_this_work/) that this somehow may not being doing it as per the instructions (aka not 'properly'). Eh? When I got over the idea suggested by the description that you should be making two steps for every move, instead of simply 'passing through a door', then it was straightforward. |
| 2018 | 22   | Part two - Dijkstra to find the shortest path through a three-dimensional graph - I know what my 'official' shortest path result should be, but I can find a valid shorter path with my input. This has happened to other people, according to the subreddits and internet chatter. I've spent best part of a day trying to create a sub-optimal version of my solution that creates the 'official' result, without success - why should I have to do this? Have submitted the 'official' result to get the star, and wrestling with the moral dilemma. |
| 2019 | 7    | Part two - instructions unclear - turned out, instruction pointer needed to be saved between invocations of each amplifier, not just the []int instructions/memory. |
| 2019 | 9    | Rewrote intcode interpreter; original Go version produced correct output, but troubled that the puzzle instructions and the code don't seem to match. Second, Lua, version fell into the ['203' trap](https://www.reddit.com/r/adventofcode/comments/e8aw9j/2019_day_9_part_1_how_to_fix_203_error/), which took a while to resolve. I now have two versions of the intcode interpreter, neither of which should work according to my reading of the puzzle instructions.  This is discouraging; other folks seem to like the intcode puzzles, however I greatly fear them. |
| 2019 | 10   | Proud of part one solution; did it all in integers without calculating any angles or creating an actual grid. Solution uses a map to hold the asteroid positions, and calculates if any asteroids lie between two other asteroids using an algorithm copied from stack overflow. It's a bit O(n3) as it does three nested loops over the asteroid map, and takes nearly a second on my machine, so no prizes for efficiency.  Part two does calculate angles, but only when sorting a slice of asteroids visible from the laser point. |
| 2019 | 14   | Could see the data structures and general approach as soon as I saw the input, but there was something in the middle that just wouldn't gel. Eventually borrowed heavily from a couple of other solutions. |
| 2019 | 15   | Got tied in knots by the philosophical problem - does the intcode call the repair droid code, or the other way around? Got it running after the 3-4th attempt (the droid calls the intcode every step). |
| 2019 | 16   | Instructions not fully grokked, eventually muddled my way to the solutions. I may have reached my personal brain-limit. |
| 2019 | 18   | Part one - tried several obvious and simple BFS/permutation approaches to this, but none of them would yield a decent run time (26! is a very big number). Eventually stumbled on a simple approach that works like magic; I understand how it find all the keys, I'm just not sure why it finds the shortest path. Part two takes an hour to run, and crashes luajit, so there's some work to do here still. |
| 2019 | 19   | Simpler than other recent puzzles, and back to that pesky intcode. I really should make a more elegant incode interpreter. Eschewing any tricky math, I treated myself to some simple brute-force solutions. |
| 2019 | 21   | Part one was easy once I stopped thinking I should actually model a deck of cards, and just keep track of card 2019 using mod arithmetic. Read part two and thought "nope", I just don't have the math skills for that, and moved on. |
| 2020 | 7    | Only half-grokked this one; loaded the input into a map of maps, which was fine, but maybe should have used a map of trees. |
| 2020 | 16   | Like 2018/6; great puzzle, whittling down occurrence lists.  |
| 2020 | 18   | For part one, used a simple infix-to-postfix converter (that I first borrowed from Donald Alcock's "Illustrating Pascal" back in 1989) and a postfix evaluator. Delighted when I found that I only had to change one character in the precedence map to solve part two. |
| 2020 | 19   | Part one tricky enough. Saw part two, thought for a moment, then said "nope". |
| 2020 | 20   | Stonking big puzzle. Stumbled upon a cheat way of solving part 1, can't yet wrap my brain around a way to do the first half of part 2 (assemble the tiles) in an efficient way (the second part, find the pattern in the picture, seems straight forward, if I could just get to it). |
| 2020 | 21   | Trouble grokking the instructions for part one, still not convinced that the algorithm implied by the instructions is accurate. Part two is a clusterfuck - there are two possible solutions for the example input, not one, so I don't see how the ingredient occurrence lists can be reduced. |
| 2021 | 8    | Part two looked tricky, and most of the solutions in the subreddit looked, well, messy. Eventually found an obvious and simple way of doing it. Moral: stare at the input data until your eyes go fuzzy, then stare some more. |
| 2021 | 9    | Classic AoC puzzle. After becoming dispirited with 2020 19-21, confidence was restored when I found I could type out a BFS from memory and have it work first time. |
| 2021 | 15   | Dijkstra is just BFS with a priority queue, yes? TIL implementing the priority queue with container/heap is 50x faster than doing it with a sort. |
| 2021 | 18   | Snailfish numbers? Can't decide between a regexp based or binary tree based solution. Suspect there's a third, simpler, solution but couldn't see it. Eventually the indecision got me beat and I put this aside for later. |
| 2022 | 16   | "This looks alright" I remember thinking, "just an adapted A-star". Start tapping, but soon get tied up in knots because of the two constraints (max flow and 30 minutes) and the way the flow rate accumulates as the paths are followed. A quick scan through the solutions subreddit keeps hinting "Floyd–Warshall algorithm", which is new to me, and looks like another of those black magic algorithms. So, the solution is to use this to make another input from the original, and run a depth first search on that. Part two is just to run another path finder after the first, with tighter constraints (less time, and no duplicated valves). |
| 2022 | 19   | Threw together a BFS, adding all possible futures to the queue. Seemed logically sound, but took too long (minutes) for part one and ran out of memory for part two. Consulted the solutions subreddit, disregarded several heuristic prunes and shortcuts, in favour of a subtle and remarkably effective optimization by Jonathan Paulson (blessed be his name). Run times now 1 and 5 seconds. |
| 2023 | 19   | Just not grokking part two at the moment ...                 |
| 2023 | 21   | Part one coded up in 15 minutes, got right answer first time. Second part? Well, a brute force extension of part one would probably never finish before the electricity ran out (6 hours to do 5000 steps, gets exponentially slower), so spending time exploring patterns in the input and partial outputs. Generating lots of numbers that don't seem to fit together. |
