-- https://adventofcode.com/2019/day/24

-- ProggyVector

-- false and nil are false, anything else is true
-- 0 and "" are both true

local log = require 'log'

local function key(y, x, z)
	return string.format('%d,%d,%d', y, x, z)
end

local function unkey(k)
	local y, x, z = k:match('([0-9-]+),([0-9-]+),([0-9-]+)')
	return tonumber(y), tonumber(x), tonumber(z)
end

local directions = {
	{y=0, x=1},
	{y=0, x=-1},
	{y=-1, x=0},
	{y=1, x=0},
}

local function neighbours(grid, y, x, z)
	local count = 0

	if x == 1 then
		if grid[key(3,2,z-1)] then
			count = count + 1
		end
	elseif x == 5 then
		if grid[key(3,4,z-1)] then
			count = count + 1
		end
	end
	if y == 1 then
		if grid[key(2,3,z-1)] then
			count = count + 1
		end
	elseif y == 5 then
		if grid[key(4,3,z-1)] then
			count = count + 1
		end
	end

	if x == 2 and y == 3 then
		for yy=1,5 do
			if grid[key(yy,1,z+1)] then
				count = count + 1
			end
		end
	elseif x == 4 and y == 3 then
		for yy=1,5 do
			if grid[key(yy,5,z+1)] then
				count = count + 1
			end
		end
	elseif x == 3 and y == 2 then
		for xx=1,5 do
			if grid[key(1,xx,z+1)] then
				count = count + 1
			end
		end
	elseif x == 3 and y == 4 then
		for xx=1,5 do
			if grid[key(5,xx,z+1)] then
				count = count + 1
			end
		end
	end

	if not (x == 3 and y == 3) then
		for _, dir in ipairs(directions) do
			local ny, nx = y + dir.y, x + dir.x
			if not (nx < 1 or ny < 1 or nx > 5 or ny > 5) then
				if grid[key(ny,nx,z)] then
					count = count + 1
				end
			end
		end
	end
	return count
end

local function minute(oldgrid, h, w)
	local newgrid = {}
	local minz, maxz = 32767, -32767
	for bug in pairs(oldgrid) do
		local _, _, z = unkey(bug)
		minz = math.min(minz, z)
		maxz = math.max(maxz, z)
	end

	for z=minz-1, maxz+1 do
		for y = 1, h do
			for x = 1, w do
				if not (y == 3 and x == 3) then
					local n = neighbours(oldgrid, y, x, z)
					local k = key(y,x,z)
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
		end
	end

	return newgrid
end

local function partTwo()
	local grid = {}
	local y = 1
	for line in io.lines('24-input.txt') do
		local x = 1
		for ch in line:gmatch'.' do
			local k = key(y,x,0)
			if ch == '#' then
				grid[k] = true
			end
			x = x + 1
		end
		y = y + 1
	end

	for _=1,200 do
		grid = minute(grid, 5, 5)
	end

	local result = 0
	for _ in pairs(grid) do
		result = result + 1
	end

	return result
end

log.report('Part Two %d\n', partTwo())	-- 2059
