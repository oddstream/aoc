local inputFilename = "02-input.txt"

local ROCK, PAPER, SCISSORS = 1, 2, 3
local LOSE, DRAW, WIN = 0, 3, 6

local outcome1 = {
	["A X"] = DRAW + ROCK,
	["A Y"] = WIN + PAPER,
	["A Z"] = LOSE + SCISSORS,
	["B X"] = LOSE + ROCK,
	["B Y"] = DRAW + PAPER,
	["B Z"] = WIN + SCISSORS,
	["C X"] = WIN + ROCK,
	["C Y"] = LOSE + PAPER,
	["C Z"] = DRAW + SCISSORS,
}

local outcome2 = {
	["A X"] = LOSE + SCISSORS,
	["A Y"] = DRAW + ROCK,
	["A Z"] = WIN + PAPER,
	["B X"] = LOSE + ROCK,
	["B Y"] = DRAW + PAPER,
	["B Z"] = WIN + SCISSORS,
	["C X"] = LOSE + PAPER,
	["C Y"] = DRAW + SCISSORS,
	["C Z"] = WIN + ROCK,
}

local score1, score2 = 0, 0

for line in io.lines(inputFilename) do
	score1 = score1 + outcome1[line]
	score2 = score2 + outcome2[line]
end

io.write(string.format('part 1 score=%d\n', score1))
io.write(string.format('part 2 score=%d\n', score2))

--[[
part 1 op=13359, me=12794
part 2 op=4938, me=14979
]]