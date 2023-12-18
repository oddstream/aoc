-- https://adventofcode.com/2023/day/16

-- very workman-like and slow implementation (part 2 is around 15 minutes)
-- uses a horribly inefficient whole-grid checksum map to check for loops

-- for a clean and straightforward Go implementation, see:
-- https://github.com/akash-panda-dev/advent-of-code/blob/main/2023/day16/main.go

local log = require 'log'

local dirs = {
	n = {x=0, y=-1},
	e = {x=1, y=0},
	s = {x=0, y=1},
	w = {x=-1, y=0},
}

---comment
---@param grid table[]
---@return string
local function gridChecksum(grid)
	local str = ''
	for _, row in ipairs(grid) do
		for _, obj in ipairs(row) do
			for b, _ in pairs(obj.beams) do
				str = str .. b
			end
		end
	end
	return str
end

---comment
---@param grid table[]
---@return table
local function activeBeams(grid)
	local arr = {}
	for y, row in ipairs(grid) do
		for x, obj in ipairs(row) do
			for _, v in pairs(obj.beams) do
				if v == true then
					arr[#arr+1] = {x=x, y=y}
				end
			end
		end
	end
	return arr
end

---comment
---@param grid table[]
---@param y integer
---@param x integer
local function stepBeams(grid, y, x)

	local function inGrid(y0, x0)
		if y0 < 1 or x0 < 1 then
			return false
		elseif y0 > #grid then
			return false
		elseif x0 > #grid[1] then
			return false
		end
		return true
	end

	local function moveBeam(newd)
		local ny, nx =  y + dirs[newd].y, x + dirs[newd].x
		if inGrid(ny, nx) then
			grid[ny][nx].energized = true
			grid[ny][nx].beams[newd] = true -- add dir to set of beams
		end
	end

	local obj = grid[y][x]
	for d, _ in pairs(obj.beams) do	-- d will be one of news

		grid[y][x].beams[d] = false

		if obj.ch == '.' then
			moveBeam(d)
		elseif obj.ch == '/' then
			if d == 'e' then
				moveBeam('n')
			elseif d == 'w' then
				moveBeam('s')
			elseif d == 'n' then
				moveBeam('e')
			elseif d == 's' then
				moveBeam('w')
			end
		elseif obj.ch == '\\' then
			if d == 'e' then
				moveBeam('s')
			elseif d == 'w' then
				moveBeam('n')
			elseif d == 'n' then
				moveBeam('w')
			elseif d == 's' then
				moveBeam('e')
			end
		elseif obj.ch == '-' then
			if d == 'e' or d == 'w' then
				moveBeam(d)
			else
				moveBeam('e')
				moveBeam('w')
			end
		elseif obj.ch == '|' then
			if d == 'n' or d == 's' then
				moveBeam(d)
			else
				moveBeam('n')
				moveBeam('s')
			end
		else
			print('unknown ch', obj.ch, 'at', y, x)
		end
	end
end

---returns the number of energized cells in the gid
---@param grid table[]
---@return integer
local function countEnergized(grid)
	local count = 0
	for _, row in ipairs(grid) do
		for _, obj in ipairs(row) do
			if obj.energized == true then
				count = count + 1
			end
		end
	end
	return count
end

---comment
---@param grid table[]
---@return integer
local function run(grid)
	local states = {}
	states[gridChecksum(grid)] = true

	local active = activeBeams(grid)
	while #active > 0 do
		for _, pos in ipairs(active) do
			stepBeams(grid, pos.y, pos.x)
		end
		local chk = gridChecksum(grid)
		if states[chk] == true then
			break
		else
			states[chk] = true
		end
		active = activeBeams(grid)
	end
	return countEnergized(grid)
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result

	local grid = {}
	for line in io.lines(filename) do
		local arr = {}
		for ch in line:gmatch'.' do
			arr[#arr+1] = {ch=ch, energized=false, beams={}}
		end
		grid[#grid+1] = arr
	end

	grid[1][1] = {ch=grid[1][1].ch, energized=true, beams={['e']=true}}
	result = run(grid)

	if expected and result ~= expected then
		log.error('part one should be %d\n', expected)
	end
	return result
end

---comment
---@param grid table[]
local function resetGrid(grid)
	for _, row in ipairs(grid) do
		for _, obj in ipairs(row) do
			obj.energized = false
			obj.beams = {}
		end
	end
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result = 0

	local grid = {}
	for line in io.lines(filename) do
		local arr = {}
		for ch in line:gmatch'.' do
			arr[#arr+1] = {ch=ch, energized=false, beams={}}
		end
		grid[#grid+1] = arr
	end

	local res
	for y = 1, #grid do
		resetGrid(grid)
		grid[y][1] = {ch=grid[y][1].ch, energized=true, beams={['e']=true}}
		res = run(grid)
		if res > result then result = res end

		resetGrid(grid)
		grid[y][#grid[1]] = {ch=grid[y][#grid[1]].ch, energized=true, beams={['w']=true}}
		res = run(grid)
		if res > result then result = res end
		print('ew', result)
	end
	for x = 1, #grid[1] do
		resetGrid(grid)
		grid[1][x] = {ch=grid[1][x].ch, energized=true, beams={['s']=true}}
		res = run(grid)
		if res > result then result = res end

		resetGrid(grid)
		grid[#grid][x] = {ch=grid[#grid][x].ch, energized=true, beams={['n']=true}}
		res = run(grid)
		if res > result then result = res end
		print('ns', result)
	end

	if expected and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('16-test.txt', 46))
log.report('part one      %d\n', partOne('16-input.txt', 7242))
log.report('part two test %d\n', partTwo('16-test.txt', 51))
log.report('part two      %d\n', partTwo('16-input.txt', 7572))