import math, algorithm

type Package = tuple
  amount: int
  quantum: int

var instructions: seq[int] = @[1, 3, 5, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61,
67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113]

var presents = sorted(instructions, cmp, SortOrder.Descending)

func calc(remaining: int, presents: seq[int], combinations: var seq[Package], used=0, qe=1) =
  if remaining == 0:
    combinations.add((used, qe))
  elif remaining > 0 and used < 6 and len(presents) > 0:
    let
      first = presents[0]
      rest = presents[1 .. presents.high]
    calc(remaining-first, rest, combinations, used+1, qe*first)
    calc(remaining, rest, combinations, used, qe)

proc findSolution(goal: int): int =
  var combinations = newSeq[Package]()
  calc(goal, presents, combinations)
  return min(combinations)[1]

echo findSolution(instructions.sum() div 3)
echo findSolution(instructions.sum() div 4)