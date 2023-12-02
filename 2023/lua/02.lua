local log = require 'log'

local input = '02-input.txt'

local bagContents = {
	red = 12,
	green = 13,
	blue = 14,
}

local function partOne()
	local result = 0
	for line in io.lines(input) do
		local badGame = false
		-- extract the game number from the start of the line
		local game = string.match(line, 'Game ([%d]+): ')
		-- extract the rest of the line
		-- append a ; to make parsing easier
		local rhs = string.sub(line, string.find(line, ': ') + 2, -1) .. ';'
		-- log.trace('GAME: %d RHS: %s\n', game, rhs)
		-- for each of the ; delimited sets of cubes
		for set in string.gmatch(rhs, '([%a%d ,]+);') do
			-- log.trace('hand %s\n', hand)
			-- for each of the list of <number> <color> in the set of cubes
			for sub in string.gmatch(set, '([%d]+ [%w]+)') do
				-- extract the number of cubes
				local num = tonumber(string.match(sub, '([%d]+)'))
				-- extract the cube color
				local col = string.match(sub, '([%l]+)')
				-- log.trace('sub: %s, num: %d, col: %s\n', sub, num, col)
				-- if the number of this color exceeds the target number ...
				if num > bagContents[col] then
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

local function partTwo()
	local result = 0
	for line in io.lines(input) do
		local maxColors = {
			red = 0,
			green = 0,
			blue = 0,
		}
		-- we don't care about the game number
		-- local game = string.match(line, 'Game ([%d]+): ')
		local rhs = string.sub(line, string.find(line, ': ') + 2, -1) .. ';'
		for set in string.gmatch(rhs, '([%a%d ,]+);') do
			for sub in string.gmatch(set, '([%d]+ [%w]+)') do
				local num = tonumber(string.match(sub, '([%d]+)'))
				local col = string.match(sub, '([%l]+)')
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

log.info('part one %d\n', partOne()) -- 1853
log.info('part two %d\n', partTwo()) -- 72706