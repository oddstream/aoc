-- https://adventofcode.com/2023/day/ Camel Cards

local log = require 'log'

local cardranks -- forward declaration

---@param hand string
---@return number
local function handScorer(hand)

	local cardmap = {}

	local function isOfAKind(n)
		-- local nj
		-- if jokers then
		-- 	nj = cardmap.J
		-- else
		-- 	nj = 0
		-- end
		for _, v in pairs(cardmap) do
			if v == n then
				return true
			end
		end
		return false
	end

	local function isFullHouse()
		for _, v in pairs(cardmap) do
			if v == 3 then
				for _, v2 in pairs(cardmap) do
					if v2 == 2 then
						return true
					end
				end
			end
		end
		return false
	end

	local function countPairs()
		local twos = 0
		for _, v in pairs(cardmap) do
			if v == 2 then
				twos = twos + 1
			end
		end
		return twos
	end

	local function highCard()
		for _, v in pairs(cardmap) do
			if v > 1 then
				return false
			end
		end
		return true
	end

	for i = 1, #cardranks do
		cardmap[cardranks:sub(i,i)] = 0
	end
	for i = 1, #hand do
		local c = hand:sub(i,i)
		cardmap[c] = cardmap[c] + 1
	end

	local score = 0

	if isOfAKind(5) then
		score = 7
	elseif isOfAKind(4) then
		score = 6
	elseif isFullHouse() then
		score = 5
	elseif isOfAKind(3) then
		score = 4
	elseif countPairs() == 2 then
		score = 3
	elseif countPairs() == 1 then
		score = 2
	elseif highCard() then
		score = 1
	else
		log.error('unknown hand %s\n', hand)
	end

	return score
end

---@param hand string
---@return number
local function jokerHandScorer(hand)
	-- brute force, try every subsitution of Joker and return the highest ranked
	local max = 0
	for i = 1, #cardranks do
		local c = cardranks:sub(i,i)
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
			local ia = cardranks:find(ca)
			local cb = b.hand:sub(i,i)
			local ib = cardranks:find(cb)
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

	cardranks = 'AKQJT98765432' -- conventional

	local hands = {}
	for line in io.lines(filename) do
		local hand, bid
		hand, bid = line:match'(%S+) (%d+)'
		table.insert(hands, {hand=hand, bid = bid, score=handScorer(hand)})
	end

	table.sort(hands, handSorter)

	-- for rank, h in pairs(hands) do
	-- 	log.trace('%d. %s %d, %d %d\n', rank, h.hand, h.score, h.bid, rank * h.bid)
	-- end

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

	cardranks = 'AKQT98765432J'	-- J is the lowest ranked card

	local hands = {}
	for line in io.lines(filename) do
		local hand, bid
		hand, bid = line:match'(%S+) (%d+)'
		table.insert(hands, {hand=hand, bid = bid, score=jokerHandScorer(hand)})
	end

	table.sort(hands, handSorter)

	-- for rank, h in pairs(hands) do
	-- 	log.trace('%d. %s %d, %d %d\n', rank, h.hand, h.score, h.bid, rank * h.bid)
	-- end

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