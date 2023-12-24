-- https://adventofcode.com/2022/day/12 Hill Climbing Algorithm

local log = require 'log'

local directions = {
	{x=1, y=0},
	{x=0, y=1},
	{x=0, y=-1},
	{x=-1, y=0},
}

---@param grid table[]
---@param y integer
---@param x integer
---@return boolean
local function ingrid(grid, y,x)
	return y >= 1 and x >= 1 and y <= #grid and x <= #grid[1]
end

---@param filename string
---@return table[], integer, integer, integer, integer
local function loadGrid(filename)
	local grid = {}
	local y = 1
	local starty, startx, goaly, goalx
	for line in io.lines(filename) do
		grid[y] = {}
		local x = 1
		for ch in line:gmatch'(%a)' do
			local height
			if ch == 'E' then
				goaly, goalx = y, x
				ch = 'z'
			elseif ch == 'S' then
				starty, startx = y, x
				ch = 'a'
			end
			height = string.byte(ch) - string.byte('a')
			assert(height>=0 and height <=26)
			table.insert(grid[y], {y=y, x=x, height=height})
			x = x + 1
		end
		y = y + 1
	end

	return grid, starty, startx, goaly, goalx
end

---@param grid table[]
---@param start table
---@param goal table
---@return boolean
local function bfs(grid, start, goal)
	local q = {start}
	-- start.parent = start
	while #q > 0 do
		local t = table.remove(q, 1)
		if t.y == goal.y and t.x == goal.x then
			return true
		end
		for _, dir in ipairs(directions) do
			local ny, nx = t.y + dir.y, t.x + dir.x
			if ingrid(grid, ny, nx) then
				local tn = grid[ny][nx]
				if tn.parent == nil then
					if tn.height <= t.height + 1 then
						tn.parent = t
						q[#q+1] = tn
						-- print(tn.y, tn.x, tn.height)
					end
				end
			end
		end
	end
	return false
end

---@param start table
---@param goal table
---@return integer
local function bfscount(start, goal)
	local count = 0
	while goal ~= nil do
		if goal.y == start.y and goal.x == start.x then
			break
		end
		goal = goal.parent
		count = count + 1
	end
	return count
end

---@param grid table[]
local function reset(grid)
	for _, row in ipairs(grid) do
		for _, t in ipairs(row) do
			t.parent = nil
		end
	end
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	local grid, starty, startx, goaly, goalx = loadGrid(filename)

	if bfs(grid, grid[starty][startx], grid[goaly][goalx]) then
		result = bfscount(grid[starty][startx], grid[goaly][goalx])
	end

	if expected and result ~= expected then
		log.error('expected %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)

	-- run from all 'a' positions and find shortest path
	-- For part 2, we can perform a neat trick by starting BFS at the endpoint,
	-- and find the shortest starting point available (stop at an a/height=0)

	local grid, _, _, goaly, goalx = loadGrid(filename)
	local goal = grid[goaly][goalx]
	local result = 1/0
	for _, row in ipairs(grid) do
		for _, start in ipairs(row) do
			if start.height == 0 then	-- 518 of these
				reset(grid)
				if bfs(grid, start, goal) then
					local c = bfscount(start, goal)
					if c < result then result = c end
				end
			end
		end
	end

	if expected and result ~= expected then
		log.error('expected %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('12-input-test.txt', 31))
log.report('part one      %d\n', partOne('12-input.txt', 408))
log.report('part two test %d\n', partTwo('12-input-test.txt', 29))
log.report('part two      %d\n', partTwo('12-input.txt', 399))

--[[
$ time luajit 12.lua
Lua 5.1
part one test 31
part one      408
part two test 29
part two      399

real	0m0.236s
user	0m0.236s
sys	0m0.000s
]]