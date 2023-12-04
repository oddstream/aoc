-- https://adventofcode.com/2023/day/4

local log = require 'log'

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0
	-- split each line into three parts
	for line in io.lines(filename) do
		-- don't care about the card number in part one
		local _, winning, numbers = line:match'Card%s+(%d+): ([%d ]+)%|([%d ]+)'
		-- create a map of winning numbers
		local winMap = {}
		for num in winning:gmatch'(%d+)' do
			winMap[num] = true
		end
		-- see how many points we've won
		local points = 0
		for num in numbers:gmatch'(%d+)' do
			if winMap[num] == true then
				if points == 0 then
					points = 1
				else
					points = points * 2
				end
			end
		end
		result = result + points
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
	local copies = {}
	--@param n integer
	local function bumpCopies(n)
		copies[n] = copies[n] or 0
		copies[n] = copies[n] + 1
	end
	local result = 0
	-- split each line into three parts
	for line in io.lines(filename) do
		local cardno, winning, numbers = line:match'Card%s+(%d+): ([%d ]+)%|([%d ]+)'
		cardno = tonumber(cardno)
		bumpCopies(cardno)

		-- create a map of winning numbers
		local winMap = {}
		for num in winning:gmatch'(%d+)' do
			winMap[num] = true
		end
		-- see how many matches this card has
		local matches = 0
		for num in numbers:gmatch'(%d+)' do
			if winMap[num] == true then
				matches = matches + 1
			end
		end
		-- if card 10 were to have 5 matching numbers,
		-- you would win one copy each of cards 11, 12, 13, 14, and 15
		-- we have won an extra copy of the next `wins` cards
		for _ = 1, copies[cardno] do
			for n = 1, matches do
				bumpCopies(cardno+n)
			end
		end
	end
	for _, v in ipairs(copies) do
		result = result + v
	end
	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('04-test.txt', 13))
log.report('part one      %d\n', partOne('04-input.txt', 19135))
log.report('part two test %d\n', partTwo('04-test.txt', 30))
log.report('part two      %d\n', partTwo('04-input.txt', 5704953))

--[[
$ time luajit 04.lua
Lua 5.1
part one test 13
part one      19135
part two test 30
part two      5704953

real	0m0.101s
user	0m0.094s
sys	0m0.004s
]]