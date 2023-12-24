-- https://adventofcode.com/2022/day/9 Rope Bridge

local log = require 'log'

local directions = {
	L = {x = -1, y = 0},
	R = {x = 1, y = 0},
	U = {x = 0, y = -1},
	D = {x = 0, y = 1},
}

---@param x integer
---@param y integer
---@return string
local function key(x,y)
	return tostring(x) .. ',' .. tostring(y)
end

---@param n integer
---@return integer
local function sign(n)
	if n < 0 then
		return -1
	elseif n > 0 then
		return 1
	end
	return 0
end

---@param filename string
---@param n integer number of knots (2 or 10)
---@return integer
local function solve(filename, n)
	local knots = {}
	for _ = 1, n do
		knots[#knots+1] = {x=30, y=30}
	end

	local tvisited = {} -- map of positions visited by last knot
	tvisited[key(knots[1].x,knots[1].y)] = true

	for line in io.lines(filename) do
		local dir, steps = line:match'(%u) (%d+)'
		for _ = 1, tonumber(steps) do
			-- move the head (knots[1])
			knots[1].x = knots[1].x + directions[dir].x
			knots[1].y = knots[1].y + directions[dir].y

			-- move each subsequent knot (knots[2 .. n])
			for k = 2, #knots do
				local diff = {x=knots[k-1].x - knots[k].x, y=knots[k-1].y - knots[k].y}
				if math.abs(diff.x) > 1 or math.abs(diff.y) > 1 then
					knots[k].x = knots[k].x + sign(diff.x)
					knots[k].y = knots[k].y + sign(diff.y)
				end
			end

			-- record position of last knot
			tvisited[key(knots[#knots].x, knots[#knots].y)] = true
		end
	end

	local tcount = 0; for _, _ in pairs(tvisited) do tcount = tcount + 1 end; return tcount
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = solve(filename, 2)
	if expected and result ~= expected then
		log.error('expecting %d:\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result = solve(filename, 10)
	if expected and result ~= expected then
		log.error('expecting %d:\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('09-input-test.txt', 13))
log.report('part one      %d\n', partOne('09-input.txt', 5619))
log.report('part two test %d\n', partTwo('09-input-test.txt', 1))
log.report('part two test %d\n', partTwo('09-input-test2.txt', 36))
log.report('part two      %d\n', partTwo('09-input.txt', 2376))

--[[
$ time luajit 09.lua
Lua 5.1
part one test 13
part one      5619
part two test 1
part two test 36
part two      2376

real    0m0.068s
user    0m0.051s
sys     0m0.016s
]]