-- https://adventofcode.com/2019/day/22

-- ProggyVector

local log = require 'log'

local function dealnew(c, n)
	-- 0123456789
	-- 9876543210
	-- 4 moves from pos 4 to pos 5
	-- -4 - 1 == 5 % 10 = 5
	return (-c - 1) % n
end

local function cut(c, n, i)
	-- cut 3
	-- 0123456789
	-- 3456789012
	-- 4 moves from pos 4 to pos 1
	-- 4 - 3 == 1 % 10 = 1
	return (c - i) % n
end

local function dealinc(c, n, i)
	-- deal 3
	-- 0123456789
	-- 0741852963
	-- 4 moves from pos 4 to pos 2
	-- 4 * 3 == 12 % 10 == 2
	return (c * i) % n
end

local function partOne()
	-- don't make an actual deck,
	-- which would be slow and fiddly with Lua 1-based arrays,
	-- instead just track where the target card is
	local n = 10007	-- length of deck
	local c = 2019	-- initial position of card 2019
	for line in io.lines('22-input.txt') do
		local m = line:match'deal with increment (%d+)'
		if m then
			c = dealinc(c, n, tonumber(m))
		else
			m = line:match'cut ([0-9-]+)'
			if m then
				c = cut(c, n, tonumber(m))
			else
				m = line:match'deal into new stack'
				if m then
					c = dealnew(c, n)
				else
					log.error('cannot parse %s\n', line)
				end
			end
		end
	end
	return c
end

local function partTwo()
	-- the deck size is prime
	-- so mod and/or gcd will be used
	-- that's all I've got
end

log.report('Part One %d\n', partOne())	-- 1234