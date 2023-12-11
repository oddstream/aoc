-- https://adventofcode.com/2023/day/11 Cosmic Expansion

local log = require 'log'


---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	local galaxy = {}

	local function emptyColumn(col)
		for row = 1, #galaxy do
			if galaxy[row][col].ch == '#' then
				return false
			end
		end
		return true
	end

	local function insertEmptyColumn(col)
		for row = 1, #galaxy do
			local new = {}
			for i = 1, col do
				table.insert(new, galaxy[row][i])
			end
			table.insert(new, {ch='.'})
			for i = col+1, #galaxy[row] do
				table.insert(new, galaxy[row][i])
			end
			table.remove(galaxy, row)
			table.insert(galaxy, row, new)
		end
	end

	local function allStars()
		local arr = {}
		for row = 1, #galaxy do
			for col = 1, #galaxy[1] do
				if galaxy[row][col].ch == '#' then
					table.insert(arr, galaxy[row][col])
				end
			end
		end
		return arr
	end

	local function distance (x1, y1, x2, y2)
		return math.abs(x1 - x2 + y1 - y2)
	end

	local function displayGalaxy()
		for _, row in ipairs(galaxy) do
			for col = 1, #row do
				io.write(row[col].ch)
			end
			io.write('\n')
		end
	end

	--

	local lineno = 1
	for line in io.lines(filename) do
		local isEmpty = true
		galaxy[lineno] = {}
		for ch in line:gmatch'.' do
			if ch == '#' then isEmpty = false end
			table.insert(galaxy[lineno], {ch=ch})
		end
		if isEmpty then
			galaxy[lineno+1] = {}
			for _ = 1, #galaxy[lineno] do
				table.insert(galaxy[lineno+1], {ch='.'})
			end
			lineno = lineno + 1
		end
		lineno = lineno + 1
	end

	for i = #galaxy[1], 1, -1 do
		if emptyColumn(i) then
			insertEmptyColumn(i)
		end
	end

	-- print('width ', #galaxy[1], 'height', #galaxy)

	for row, _ in ipairs(galaxy) do
		for col, _ in ipairs(galaxy[row]) do
			galaxy[row][col].row = row
			galaxy[row][col].col = col
		end
	end

--[[
	local start = galaxy[1][5]
	local goal = galaxy[3][1]
	bfs(start, goal)
	print(bfscount(start, goal))
]]

	local stars = allStars()
	-- print('#all stars', #all)
	-- for _, v in ipairs(all) do
	-- 	log.trace('%d,%d\n', v.col, v.row)
	-- end

	assert(distance(2, 7, 6, 12)==9)
	assert(distance(7, 2, 6, 12)==9)
	assert(distance(7, 2, 12, 6)==9)
	assert(distance(2, 7, 12, 6)==9)
	-- assert(distance(12, 2, 7, 6)==9)

	local npairs = 0
	for i = 1, #stars do
		-- assert(all[i].ch=='#')
		for j = i+1, #stars do
			-- assert(all[j].ch=='#')
			local dist = distance(stars[i].col, stars[i].row, stars[j].col, stars[j].row)
			-- print(all[i].row, all[i].col, all[j].row, all[j].row, ':=', dist)
			result = result + dist
			npairs = npairs + 1
		end
	end
	log.trace('stars = %d, npairs = %d\n', #stars, npairs)

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

	-- local f = assert(io.open(filename, "r"))
	-- local content = f:read("*all")
	-- f:close()
	for line in io.lines(filename) do
		log.trace(line)
	end

	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('11-test.txt', 374))
log.report('part one      %d\n', partOne('11-input.txt'))	-- 6931238 too low
-- log.report('part two test %d\n', partTwo('11-test.txt', 0))
-- log.report('part two      %d\n', partTwo('11-input.txt', 0))

