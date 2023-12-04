-- https://adventofcode.com/2023/day/2

local log = require 'log'

local bagContents = {
	red = 12,
	green = 13,
	blue = 14,
}

local function partOne(filename)
	local result = 0
	for line in io.lines(filename) do
		local badGame = false
		-- extract the game number from the start of the line
		-- and everything to the right hand side
		-- append a ; to make parsing easier
		local game, rhs = string.match(line .. ';', 'Game (%d+): (.+)')
		game = tonumber(game)
		-- log.trace('GAME: %d RHS: %s\n', game, rhs)
		-- for each of the ; delimited sets of cubes ...
		for set in string.gmatch(rhs, '([%l%d ,]+);') do
			-- log.trace('set %s\n', set)
			-- for each of the list of <number> <color> in the set of cubes ...
			for sub in string.gmatch(set, '(%d+ %l+)') do
				-- extract the number of cubes and the cube color
				local num, col = string.match(sub, '(%d+) (%l+)')
				-- log.trace('sub: %s, num: %d, col: %s\n', sub, num, col)
				-- if the number of this color exceeds the target number ...
				if tonumber(num) > bagContents[col] then
					-- log.error('nope\n')
					badGame = true
					break
				end
			end
		end
		if not badGame then
			result = result + game
		end
	end
	return result
end

local function partTwo(filename)
	local result = 0
	for line in io.lines(filename) do
		local maxColors = {
			red = 0,
			green = 0,
			blue = 0,
		}
		-- we don't care about the game number
		local _, rhs = string.match(line ..';', 'Game (%d+): (.+)')
		for set in string.gmatch(rhs, '([%l%d ,]+);') do
			for sub in string.gmatch(set, '(%d+ %l+)') do
				local num, col = string.match(sub, '(%d+) (%l+)')
				num = tonumber(num)
				if num > maxColors[col] then
					maxColors[col] = num
				end
			end
		end
		local gamePower = 1
		for _, v in pairs(maxColors) do
			gamePower = gamePower * v
		end
		result = result + gamePower
	end
	return result
end

print(_VERSION)
log.report('part one test %d\n', partOne('02-test.txt'))	-- 8
log.report('part one      %d\n', partOne('02-input.txt'))	-- 1853
log.report('part two test %d\n', partTwo('02-test.txt'))	-- 2286
log.report('part two      %d\n', partTwo('02-input.txt'))	-- 72706