-- https://adventofcode.com/2020/day/22

-- ProggyVector

---@return integer[][]
local function load()
	local decks = {}
	for line in io.lines('22-input.txt') do
		local n = line:match'Player (%d+):'
		if n then
			table.insert(decks, {})
		else
			n = line:match'%d+'
			if n then
				table.insert(decks[#decks], tonumber(n))
			end
		end
	end
	return decks
end

---@type integer[][]
local decks = load()

local function combat1()
	-- the rules do not state what happens when both cards have the same value
	-- but that doesn't happen with the supplied input data
	local c1, c2 = table.remove(decks[1], 1), table.remove(decks[2], 1)
	if c1 > c2 then
		table.insert(decks[1], c1)
		table.insert(decks[1], c2)
	elseif c2 > c1 then
		table.insert(decks[2], c2)
		table.insert(decks[2], c1)
	else
		print('cards are the same')
	end
end

---@param deck integer[]
---@return integer
local function score(deck)
	local n = 0
	local f = #deck
	for i=1,#deck do
		n = n + (deck[i] * f)
		f = f - 1
	end
	return n
end

--[[
for i=1, #decks[1] do
	io.write(tostring(decks[1][i]), ' ')
end
io.write('\n')
for i=1, #decks[2] do
	io.write(tostring(decks[2][i]), ' ')
end
io.write('\n')
]]

while #decks[1] > 0 and #decks[2] > 0 do
	combat1()
end

--[[
for i=1, #decks[1] do
	io.write(tostring(decks[1][i]), ' ')
end
io.write('\n')
for i=1, #decks[2] do
	io.write(tostring(decks[2][i]), ' ')
end
io.write('\n')
]]

if #decks[1] ~= 0 then
	print('Part One' , score(decks[1]))	-- 33421
end
if #decks[2] ~= 0 then
	print('Part One ', score(decks[2]))	-- 0
end

decks = load()

---@param d1 integer[]
---@param d2 integer[]
---@return string
local function seenkey(d1, d2)
	return table.concat(d1, ',') .. '|' .. table.concat(d2, ',')
end

---@param orig integer[]
---@param n integer
---@return integer[]
local function copyn(orig, n)
	local new = {}
	for i=1,n do
		table.insert(new, orig[i])
	end
	return new
end

---@param d1 integer[]
---@param d2 integer[]
---@return integer, integer[]
local function combat2(d1, d2)
	--[[
		I'd thought that the seen map should be outside this function,
		but that produces the wrong answer (too low)
	]]
	local seen = {}
	while #d1 > 0 and #d2 > 0 do
		local sk = seenkey(d1, d2)
		--[[
			Before either player deals a card,
			if there was a previous round in this game that had exactly the same cards
			in the same order in the same players' decks,
			the game instantly ends in a win for player 1.
			Previous rounds from other games are not considered.
		]]
		if seen[sk] then
			return 1, d1
		end
		seen[sk] = true

		--[[
			Otherwise, this round's cards must be in a new configuration;
			the players begin the round by each drawing the top card of their deck as normal.
		]]
		local c1, c2 = table.remove(d1, 1), table.remove(d2, 1)

		local winner
		--[[
			If both players have at least as many cards remaining in their deck as the value of the card they just drew,
			the winner of the round is determined by playing a new game of Recursive Combat
		]]
		if #d1 >= c1 and #d2 >= c2 then
			--[[
				To play a sub-game of Recursive Combat,
				each player creates a new deck by making a copy of the next cards in their deck
				(the quantity of cards copied is equal to the number on the card they drew to trigger the sub-game).
			]]
			winner, _ = combat2(copyn(d1, c1), copyn(d2, c2))
		else
			--[[
				Otherwise, at least one player must not have enough cards left in their deck to recurse;
				the winner of the round is the player with the higher-value card.
			]]
			if c1 > c2 then
				winner = 1
			else
				winner = 2
			end
		end
		if winner == 1 then
			table.insert(d1, c1)
			table.insert(d1, c2)
		else
			table.insert(d2, c2)
			table.insert(d2, c1)
		end
	end
	if #d1 > 0 then
		return 1, d1
	else
		return 2, d2
	end
end

local _, winningd = combat2(decks[1], decks[2])
print('Part Two' , score(winningd))	-- 33651

--[[
$ time luajit 22.lua
Part One	33421
Part Two	33651

real	0m1.520s
user	0m1.515s
sys	0m0.004s
]]