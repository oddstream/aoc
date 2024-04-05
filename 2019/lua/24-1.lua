-- https://adventofcode.com/2019/day/24

-- ProggyVector

-- false and nil are false, anything else is true
-- 0 and "" are both true

local log = require 'log'

local function key(y, x)
	-- return string.format('%d,%d', y, x)
	return y * 1000 + x
end

local function unkey(k)
	-- local y, x = k:match('([0-9-]+),([0-9-]+)')
	-- return tonumber(y), tonumber(x)
	return math.floor(k / 1000), k % 1000
end

local function bigkey(grid)
	-- clunky and slow, but quick to write
	local result = ''
	for k in pairs(grid) do
		local y, x = unkey(k)
		result = result .. tonumber(y) .. ',' .. tonumber(x) .. '/'
	end
	return result
end

local directions = {
	{y=0, x=1},
	{y=0, x=-1},
	{y=-1, x=0},
	{y=1, x=0},
}

local function neighbours(grid, y, x)
	local count = 0
	for _, dir in ipairs(directions) do
		local ny, nx = y + dir.y, x + dir.x
		local k = key(ny,nx)
		if grid[k] then
			count = count + 1

		end
	end
	return count
end

local function minute(oldgrid, h, w)
	local newgrid = {}
	for y = 1, h do
		for x = 1, w do
			local n = neighbours(oldgrid, y, x)
			local k = key(y,x)
			-- "A bug dies (becoming an empty space)
			-- unless there is exactly one bug adjacent to it."
			if oldgrid[k] and n == 1 then
				newgrid[k] = true
			end
			-- "An empty space becomes infested with a bug
			-- if exactly one or two bugs are adjacent to it."
			if not oldgrid[k] and (n == 1 or n == 2) then
				newgrid[k] = true
			end
		end
	end

	return newgrid
end

local function display(grid, h, w)
	for y = 1, h do
		for x = 1, w do
			local k = key(y,x)
			if grid[k] then
				io.write('#')
			else
				io.write('.')
			end
		end
		io.write('\n')
	end
end

local function biodiversity(grid, h, w)
	local count, i = 0, 1
	for y=1,h do
		for x=1,w do
			local k = key(y,x)
			if grid[k] then
				count = count + i
			end
			i = i * 2
		end
	end
	return count
end

local function partOne()
	local grid = {}
	local y = 1
	for line in io.lines('24-input.txt') do
		local x = 1
		for ch in line:gmatch'.' do
			local k = key(y,x)
			if ch == '#' then
				grid[k] = true
			end
			x = x + 1
		end
		y = y + 1
	end

	local seen = {
		[bigkey(grid)] = true
	}
	for i=1,20 do
		grid = minute(grid, 5, 5)
		local nbk = bigkey(grid)
		if seen[nbk] then
			return biodiversity(grid, 5, 5)	-- on iteration 13
		end
		seen[nbk] = true
	end
	return -1
end

log.report('Part One %d\n', partOne())	-- 18842609
