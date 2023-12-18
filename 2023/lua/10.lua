-- https://adventofcode.com/2023/day/10 Pipe Maze

local log = require 'log'

-- https://en.wikipedia.org/wiki/Nonzero-rule
-- https://en.wikipedia.org/wiki/Point_in_polygon
-- https://en.wikipedia.org/wiki/Even%E2%80%93odd_rule
-- https://www.reddit.com/media?url=https%3A%2F%2Fi.redd.it%2Finqgpahhuj5c1.png

-- input is 140x140 chars
-- there is one S at line 51 column 40
-- there is no branching, just (two reciprocal) paths from start

-- throughout, use a table like {x=3, y=4} as a pos
-- surround grid with a border to avoid tedious bounds checking
-- Lua 5.4 only, as we are using integer division

local dirs = {
	['n'] = {x=0, y=-1},
	['e'] = {x=1, y=0},
	['s'] = {x=0, y=1},
	['w'] = {x=-1, y=0},
}

local okDirs = {
	['S'] = {'n','e','w','s'},
	['|'] = {'n','s'},
	['-'] = {'e','w'},
	['F'] = {'e','s'},
	['J'] = {'n','w'},
	['L'] = {'n','e'},
	['7'] = {'w','s'},
}

---Return next pos from input pos, or nil
---@param grid table[]
---@param state table
---@return table?
local function move(grid, state)
	local ch = grid[state.y][state.x]
	local okdirs = okDirs[ch]
	if not okdirs then return end
	for _, dir in ipairs(okdirs) do
		local newx = state.x + dirs[dir].x
		local newy = state.y + dirs[dir].y
		local newch =  grid[newy][newx]
		if not (newch == '.' or newch == 'X') then
			return {x=newx, y=newy}
		end
	end
end

---Returns list of possible pos-es from input pos
---@param grid table[]
---@param state table
---@return table[]
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
---@return table
local function loadGrid(filename)
	local grid = {}
	local lineno = 1
	for line in io.lines(filename) do
		grid[lineno] = {}
		line = '.' .. line .. '.'
		for i = 1, #line do
			local ch = line:sub(i,i)
			table.insert(grid[lineno], ch)
		end
		lineno = lineno + 1
	end

	local linelen = #grid[1]
	table.insert(grid, 1, {})
	table.insert(grid, {})
	for _ = 1, linelen do
		table.insert(grid[1], '.')
		table.insert(grid[#grid], '.')
	end

	return grid
end

---Returns pos of 'S' in grid. could be done when loading grid, I dunno.
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

---@params filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0
	local grid = loadGrid(filename)
	local start = findStart(grid)
	if not start then return -1 end
	print('start', grid[start.y][start.x], 'at', start.y, start.x)
	local starts = moves(grid, start)

	for i, pos in ipairs(starts) do
		local path_length = 0
		grid = loadGrid(filename)
		repeat
			path_length = path_length + 1
			local newpos = move(grid, pos)
			if newpos ~= nil then
				grid[pos.y][pos.x] = 'X'
			end
			pos = newpos
		until pos == nil
		log.trace('%d path length %d\n', i, path_length)
		-- the longest path is the correct one,
		-- and it will occur twice, and be an even number
		if path_length > result then result = path_length end
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

	-- pass #1: load the grid, walk path, build visited map
	local grid = loadGrid(filename)
	local visited = {}	-- map of y,x = true

	local function key(y, x)
		return tostring(y) .. ',' .. tostring(x)
	end

	local start = findStart(grid)
	if start == nil then return -1 end
	local starts = moves(grid, start)
	local pos = starts[1]	-- we get the 1 from output of part 1
	visited[key(pos.y, pos.x)] = true
	-- grid[pos.y][pos.x] = '|'
	-- local path_length = 0
	repeat
		-- path_length = path_length + 1
		local newpos = move(grid, pos)
		if newpos ~= nil then
			grid[pos.y][pos.x] = 'X'
			visited[key(newpos.y, newpos.x)] = true
		end
		pos = newpos
	until pos == nil
	-- print(1, path_length // 2)

	-- pass #2: use in/out algo to count tiles inside path

	grid = loadGrid(filename)

	-- for y = 1, #grid do
	-- 	for x = 1, #grid[y] do
	-- 		if visited[key(y,x)] then
	-- 			io.write(grid[y][x])
	-- 		else
	-- 			io.write(' ')
	-- 		end
	-- 	end
	-- 	io.write('\n')
	-- end

	-- local n = 0	for _, _ in pairs(visited) do n = n + 1 end	print(n, 'visited', n//2)

	for y = 1, #grid do
		local inside = false
		local count = 0
		for x = 1, #grid[y] do
			local ch = grid[y][x]
			-- if ch == 'S' then ch = 'J' end	-- KLUDGE! KLUDGE! KLUDGE! Can be S7F- but not |LJ
			if visited[key(y,x)] then
				if ch == '|' or ch == 'L' or ch == 'J' then
					inside = not inside
				end
			else
			-- "Any tile that isn't part of the main loop can count as being enclosed by the loop."
				if inside then
					count = count + 1
				end
			end
		end
		result = result + count
	end
--[[
	print('grid is', #grid, #grid[1])	-- 140 140 (plus border)
	local dots = 0
	for y = 1, #grid do
		for x = 1, #grid[y] do
			local ch = grid[y][x]
			if ch == '.' then
				dots = dots + 1
			end
		end
	end
	print(dots, 'dots')	-- 507
]]
	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
-- log.report('part one test1 %d\n', partOne('10-test1.txt', 4))
-- log.report('part one test2 %d\n', partOne('10-test2.txt', 8))
-- log.report('part one test3 %d\n', partOne('10-test3.txt', 4))
-- log.report('part one test4 %d\n', partOne('10-test4.txt', 8))
log.report('part one       %d\n', partOne('10-input.txt', 7145))

-- log.report('part two test1 %d\n', partTwo('10-test1.txt', 1)) -- S is an F
-- log.report('part two test3 %d\n', partTwo('10-test3.txt', 4)) -- S is an F
-- log.report('part two test4 %d\n', partTwo('10-test4.txt', 10)) -- S is ...
log.report('part two       %d\n', partTwo('10-input.txt', 445)) -- S is an F?

--[[
$ time lua54 10.lua
Lua 5.4
start   S       at      52      41
1 path length 14290
2 path length 14247
3 path length 14290
4 path length 14246
part one       7145
S       at      52      41
1       7145
14290   visited 7145
part two       445

real    0m0.183s
user    0m0.166s
sys     0m0.008s
]]