-- https://adventofcode.com/2022/day/14 Regolith Reservoir

local log = require 'log'

-- max number in test = 503, input = 571

---use string key for debugging, numer for performance
---@param x integer
---@param y integer
---@return integer
local function key(x,y)
	return x * 1000 + y
	-- return tostring(x) .. ',' .. tostring(y)
end

-- local function unkey(n)
-- 	local y = n % 1000
-- 	local x = (n - y) / 1000
-- 	return x, y
-- end

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
---@return table
---@return integer
local function loadMap(filename)
	local map = {}
	local maxy = 0
	for line in io.lines(filename) do
		local pos
		for x,y in line:gmatch'(%d+),(%d+)' do
			x = tonumber(x)
			y = tonumber(y)
			if y > maxy then maxy = y end

			if not pos then
				pos = {x=x, y=y}
			else
				local d = {x=sign(x - pos.x), y=sign(y - pos.y)}
				repeat
					map[key(pos.x, pos.y)] = '#'
					pos.x = pos.x + d.x
					pos.y = pos.y + d.y
				until (pos.x == x) and (pos.y == y)
				-- pos = {x=x, y=y}
			end
			map[key(pos.x, pos.y)] = '#'
		end
	end
	return map, maxy
end

local dirs = {
	{x=0, y=1},	-- down
	{x=-1, y=1}, -- down left
	{x=1, y=1},	-- down right
}

local function dropsand1(map, maxy, x, y)
	repeat
		local fallen = false
		for _, d in ipairs(dirs) do
			local nx = x + d.x
			local ny = y + d.y
			if not map[key(nx,ny)] then
				x, y = nx, ny
				if ny > maxy then
					return false
				end
				fallen = true
				break
			end
		end
	until not fallen

	map[key(x,y)] = 'o'
	return true
end

local function dropsand2(map, maxy, x, y)
	if map[key(x,y)] then
		return false
	end
	repeat
		local fallen = false
		for _, d in ipairs(dirs) do
			local nx = x + d.x
			local ny = y + d.y
			if not map[key(nx,ny)] then
				x, y = nx, ny
				if y == maxy + 1 then
					break
				end
				fallen = true
				break
			end
		end
	until not fallen
	map[key(x,y)] = 'o'
	return true
end

local function showmap1(map)
	-- visual check of test input
	for y = 0, 9 do
		for x = 494, 503 do
			if map[key(x,y)] then
				io.write(map[key(x,y)])
			else
				io.write('.')
			end
		end
		io.write'\n'
	end
	io.write'\n'
end

local function showmap2(map)
	-- visual check of test input
	for y = 0, 12 do
		for x = 480, 520 do
			if map[key(x,y)] then
				io.write(map[key(x,y)])
			else
				io.write('.')
			end
		end
		io.write'\n'
	end
	io.write'\n'
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	local map, maxy = loadMap(filename)

	print('void after', maxy)

	while dropsand1(map, maxy, 500,0) do
		-- showmap1(map)
		result = result + 1
	end

	if expected and result ~= expected then
		log.error('expecting %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result = 0

	local map, maxy = loadMap(filename)

	while dropsand2(map, maxy, 500,0) do
		-- showmap2(map)
		result = result + 1
	end

	if expected and result ~= expected then
		log.error('expecting %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
-- log.report('part one test %d\n', partOne('14-input-test.txt', 24))
-- log.report('part one      %d\n', partOne('14-input.txt', 1001))
log.report('part two test %d\n', partTwo('14-input-test.txt', 93))
log.report('part two      %d\n', partTwo('14-input.txt', 27976))

--[[
$ time luajit 14.lua
Lua 5.1
part two test 93
part two      27976

real	0m0.216s
user	0m0.202s
sys	0m0.008s
]]