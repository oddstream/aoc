-- https://adventofcode.com/2023/day/14 Parabolic Reflector Dish

-- input is 100x100 chars

local log = require 'log'

---@param platform table[]
local function show(platform)
	for y = 1, #platform do
		io.write(table.concat(platform[y]), '\n')
	end
end

---@param platform table[]
---@return boolean
local function rollNorth(platform)
	local moves = 0
	for y = 2, #platform do
		for x = 1, #platform[y] do
			if platform[y][x] == 'O' and platform[y-1][x] == '.' then
				platform[y][x], platform[y-1][x] = platform[y-1][x], platform[y][x]
				moves = moves + 1
			end
		end
	end
	return moves > 0
end

---@param platform table[]
---@return boolean
local function rollSouth(platform)
	local moves = 0
	for y = #platform-1, 1, -1 do
		for x = 1, #platform[y] do
			if platform[y][x] == 'O' and platform[y+1][x] == '.' then
				platform[y][x], platform[y+1][x] = platform[y+1][x], platform[y][x]
				moves = moves + 1
			end
		end
	end
	return moves > 0
end

---@param platform table[]
---@return boolean
local function rollWest(platform)
	local moves = 0
	for y = 1, #platform do
		for x = 2, #platform[y] do
			if platform[y][x] == 'O' and platform[y][x-1] == '.' then
				platform[y][x], platform[y][x-1] = platform[y][x-1], platform[y][x]
				moves = moves + 1
			end
		end
	end
	return moves > 0
end

---@param platform table[]
---@return boolean
local function rollEast(platform)
	local moves = 0
	for y = 1, #platform do
		for x = #platform[y]-1, 1, -1 do
			if platform[y][x] == 'O' and platform[y][x+1] == '.' then
				platform[y][x], platform[y][x+1] = platform[y][x+1], platform[y][x]
				moves = moves + 1
			end
		end
	end
	return moves > 0
end

---@param platform table[]
---@return integer
local function weight(platform)
	local result = 0
	local rowfactor = #platform
	for y = 1, #platform do
		local n = 0
		for x = 1, #platform[y] do
			if platform[y][x] == 'O' then
				n = n + 1
			end
		end
		result = result + (rowfactor * n)
		rowfactor = rowfactor - 1
	end
	return result
end

---@param platform table[]
---@return string
local function simplekey(platform)
	local k = ''
	for y = 1, #platform do
		k = k .. table.concat(platform[y])
	end
	return k
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result

	local platform = {}
	for line in io.lines(filename) do
		local arr = {}
		for ch in line:gmatch'.' do
			table.insert(arr, ch)
		end
		table.insert(platform, arr)
	end

	while rollNorth(platform) do
	end

	-- show(platform)

	result = weight(platform)

	if expected ~= nil and result ~= expected then
		log.error('part one should be %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result

	local platform = {}
	for line in io.lines(filename) do
		local arr = {}
		for ch in line:gmatch'.' do
			table.insert(arr, ch)
		end
		table.insert(platform, arr)
	end

	local seen = {}
	local cyclestart, cyclelength
	for i = 1, 1000000000 do
		while rollNorth(platform) do end
		while rollWest(platform) do end
		while rollSouth(platform) do end
		while rollEast(platform) do end
		local k = simplekey(platform)
		if seen[k] == nil then
			seen[k] = i
		else
			cyclestart = i
			cyclelength = i - seen[k]
			print('loop', i, 'seen already on loop', seen[k], 'cycle length', cyclelength)
			break
		end
	end

	for _ = 1, (1000000000 - cyclestart) % cyclelength do
		while rollNorth(platform) do end
		while rollWest(platform) do end
		while rollSouth(platform) do end
		while rollEast(platform) do end
	end

	result = weight(platform)

	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('14-test.txt', 136))
log.report('part one      %d\n', partOne('14-input.txt', 106997))
log.report('part two test %d\n', partTwo('14-test.txt', 64))
log.report('part two      %d\n', partTwo('14-input.txt', 99641))

--[[
$ time luajit 14.lua
Lua 5.1
part one test 136
part one      106997
loop    10      seen already on loop    3       cycle length    7
part two test 64
loop    185     seen already on loop    176     cycle length    9
part two      99641

real    0m1.023s
user    0m1.018s
sys     0m0.006s

$ # luajit is approx 5x faster than Lua 5.4, and 6x faster than Lua 5.1
]]