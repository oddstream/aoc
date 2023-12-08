-- https://adventofcode.com/2023/day/ Camel Cards

local log = require 'log'

local cardOrder -- forward declaration

---@param hand string
---@return number
local function handScorer(hand)

	local handMap = {}

	local function isOfAKind(n)
		for _, v in pairs(handMap) do
			if v == n then
				return true
			end
		end
		return false
	end

	local function countPairs()
		local twos = 0
		for _, v in pairs(handMap) do
			if v == 2 then
				twos = twos + 1
			end
		end
		return twos
	end

	for c in cardOrder:gmatch'.' do
		handMap[c] = 0
	end
	for c in hand:gmatch'.' do
		handMap[c] = handMap[c] + 1
	end

	local score

	if isOfAKind(5) then						-- 5
		score = 7
	elseif isOfAKind(4) then					-- 41
		score = 6
	elseif isOfAKind(3) and isOfAKind(2) then	-- 321
		score = 5	-- full house
	elseif isOfAKind(3) then					-- 311
		score = 4
	elseif countPairs() == 2 then				-- 221
		score = 3
	elseif countPairs() == 1 then				-- 2111
		score = 2
	else
		score = 1 -- high card
	end

	-- could extract the above numbers from handMap
	-- and map that value to 7 .. 1

	-- or a generic 'get largest' function that returns the two largest numbers

	-- but then it just gets less readable

	return score
end

---@param hand string
---@return number
local function jokerHandScorer(hand)
	-- brute force, try every substitution of Joker and return the highest ranked
	local max = 0
	for c in cardOrder:gmatch'.' do
		if c ~= 'J' then
			local h = hand:gsub('J', c)
			local s = handScorer(h)
			if s > max then max = s end
		end
	end
	return max
end

local function handSorter(a, b)
	-- return true if first element must come before the second in the final order
	-- we need to rank in order of increasing strength
	if a.score == b.score then
		for i = 1, 5 do
			local ca = a.hand:sub(i,i)
			local ia = cardOrder:find(ca)
			local cb = b.hand:sub(i,i)
			local ib = cardOrder:find(cb)
			if ia > ib then
				return true
			elseif ia < ib then
				return false
			end
			-- else ia == ib, go round the loop again
		end
	else
		return a.score < b.score
	end
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	cardOrder = 'AKQJT98765432' -- conventional

	local hands = {}
	for line in io.lines(filename) do
		local hand, bid = line:match'(%S+) (%d+)'
		table.insert(hands, {hand=hand, bid=bid, score=handScorer(hand)})
	end

	table.sort(hands, handSorter)

	for i, h in pairs(hands) do
		local winnings = i * h.bid
		result = result + winnings
	end

	if expected ~= nil and result ~= expected then
		log.error('part one should be %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result = 0

	cardOrder = 'AKQT98765432J'	-- J is the lowest ranked card

	local hands = {}
	for line in io.lines(filename) do
		local hand, bid = line:match'(%S+) (%d+)'
		table.insert(hands, {hand=hand, bid=bid, score=jokerHandScorer(hand)})
	end

	table.sort(hands, handSorter)

	for i, h in pairs(hands) do
		local winnings = i * h.bid
		result = result + winnings
	end

	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('07-test.txt', 6440))
log.report('part one      %d\n', partOne('07-input.txt', 251545216))
log.report('part two test %d\n', partTwo('07-test.txt', 5905))
log.report('part two      %d\n', partTwo('07-input.txt', 250384185))

-- 251911491 too high
-- 251681499 too high
-- 251339180 too high
-- 250599963 wrong, wait 5 minutes
-- 250384185 that's a bingo!

--[[
$ time luajit 07.lua
Lua 5.1
part one test 6440
part one      251545216
part two test 5905
part two      250384185

real    0m0.168s
user    0m0.143s
sys     0m0.016s
]]