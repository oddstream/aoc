-- https://adventofcode.com/2023/day/10 Pipe Maze

local log = require 'log'

-- input is 140x140 chars
-- there is one S at line 51 column 40
-- there is no branching, just (two reciprocal) paths from start

local dirs = {
	['n'] = {x=0, y=-1},
	['e'] = {x=1, y=0},
	['s'] = {x=0, y=1},
	['w'] = {x=-1, y=0},
}

-- state: x, y, dir, steps

local okDirs = {
	['S'] = {'n','e','w','s'},
	['|'] = {'n','s'},
	['-'] = {'e','w'},
	['F'] = {'e','s'},
	['J'] = {'n','w'},
	['L'] = {'n','e'},
	['7'] = {'w','s'},
}

local function move(grid, state)
	local ch = grid[state.y][state.x]
	for _, dir in ipairs(okDirs[ch]) do
		local newx = state.x + dirs[dir].x
		local newy = state.y + dirs[dir].y
		local newch =  grid[newy][newx]
		if not (newch == '.' or newch == 'X') then
			return {x=newx, y=newy}
		end
	end
end

local function moves(grid, state)
	local out = {}
	local ch = grid[state.y][state.x]
	for _, dir in ipairs(okDirs[ch]) do
		local newx = state.x + dirs[dir].x
		local newy = state.y + dirs[dir].y
		local newch =  grid[newy][newx]
		if not (newch == '.' or newch == 'X') then
			table.insert(out, {x=newx, y=newy})
		end
	end
	return out
end

---@param filename string
---@return table, table
local function loadGrid(filename)
	local grid = {}
	local shadow = {}
	local lineno = 1
	for line in io.lines(filename) do
		grid[lineno] = {}
		shadow[lineno] = {}
		line = '.' .. line .. '.'
		for i = 1, #line do
			table.insert(grid[lineno], line:sub(i,i))
			table.insert(shadow[lineno], '.')
		end
		lineno = lineno + 1
	end

	local linelen = #grid[1]
	table.insert(grid, 1, {})
	table.insert(shadow, 1, {})
	table.insert(grid, {})
	table.insert(shadow, {})
	for _ = 1, linelen do
		table.insert(grid[1], '.')
		table.insert(grid[#grid], '.')
		table.insert(shadow[1], '.')
		table.insert(shadow[#shadow], '.')
	end

	return grid, shadow
end

---@param grid table
---@return table?
local function findStart(grid)
	for y, row in ipairs(grid) do
		for x, ch in ipairs(row) do
			if ch == 'S' then
				return {x=x, y=y}
			end
		end
	end
end

local function crossings(shadow, x, y)
	local n = 0
	while x > 0 and y > 0 do
		local ch = shadow[y][x]
		if ch ~= '.' and (not (ch == '7' or ch == 'L')) then
			n = n + 1
		end
		x = x - 1
		y = y - 1
	end
	return n
end

local function crossings2(shadow, x, y)
	local n = 0
	local row = shadow[y]
	while x < #row do
		local ch = row[x]
		if ch ~= '.' then --and (not (ch == '7' or ch == 'L')) then
			n = n + 1
		end
		x = x + 1
	end
	return n
end

---@params filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	local grid, shadow = loadGrid(filename)
	local start = findStart(grid)
	print(grid[start.y][start.x], shadow[start.y][start.x])
	shadow[start.y][start.x] = 'S'
	local starts = moves(grid, start)

	for i, pos in ipairs(starts) do
		local path_length = 0
		grid = loadGrid(filename)
		repeat
			path_length = path_length + 1
			local newpos = move(grid, pos)
			if newpos ~= nil then
				shadow[pos.y][pos.x] = grid[pos.y][pos.x]
				grid[pos.y][pos.x] = 'X'
			end
			pos = newpos
		until pos == nil
		log.trace('%d path length %d\n', i, path_length)
		if path_length > result then result = path_length end

		if i == 1 then
			for _, row in ipairs(shadow) do
				for _, x in ipairs(row) do
					io.write(x)
				end
				io.write('\n')
			end
--[[
Imagine a person standing on one of the squares,
and they shoot a laser in any direction
except for the four cardinal directions (to prevent the collinearity problem).
Diagonally, say.
Now trace that laser starting at the square shooting it until it exits the field.
If it it's inside the shape, it will cross the boundary an odd number of times,
if it's outside it will cross an even number of times.
Be careful if it hits a corner from the outside: it either crosses it twice, or zero times.

The simplest example is if the boundary is just a circle.
If you're inside the circle and shoot your laser, it will hit it once, on its way out.
If you're outside the circle it will hit it either zero times (misses the circle entirely)
or two times (hit once going in, hit once going out).
The corner case is if your laser is EXACTLY tangent to the circle,
then it hits it "once" (but really twice, from a mathematical perspective the tangent point is two hits).

It's intuitive for circles, but the concept generalizes to ANY enclosed shape:
if you hit the shape twice, you've "gone in" and then "gone out".
But if you were inside from the beginning, there's an extra "gone out",
so the number of crossings is odd. You just have to be careful about tangents and collinearity.
]]
			local inside = 0
			for y, row in ipairs(shadow) do
				for x, ch in ipairs(row) do
					if ch == '.' then
						local n = crossings(shadow, x, y)
						if (n > 0) and ((n % 2) == 1) then
							inside = inside + 1
						end
					end
				end
			end
			print('inside', inside)
		end
	end
	result = result // 2

	if expected ~= nil and result ~= expected then
		log.error('part one should be %d\n', expected)
	end
	return result
end

---@params filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result = 0

	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
-- log.report('part one test1 %d\n', partOne('10-test1.txt', 4))
-- log.report('part one test2 %d\n', partOne('10-test2.txt', 8))
log.report('part one       %d\n', partOne('10-input.txt', 7145))
-- part 2 490 too high
-- part 2 445
-- log.report('part one test3 %d\n', partOne('10-test3.txt', 4))
-- log.report('part one test3 %d\n', partOne('10-test4.txt', 8))

-- log.report('part two test %d\n', partTwo(0))
-- log.report('part two      %d\n', partTwo(0))

