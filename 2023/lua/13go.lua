-- https://adventofcode.com/2023/day/13 Point of Incidence
-- https://github.com/Benjababe/Advent-of-Code/blob/main/2023/Day%2013/pt1_2.go

local log = require 'log'

local function checkLeftRight(grid)
	local p1, p2 = 0, 0

	for mirror = 1, #grid[1]-1 do
		local sRange = math.min(mirror, #grid[1]-mirror)
		local l, r = mirror-1, mirror

		local count = 0

		for dx = 0, sRange-1 do
			for _, line in ipairs(grid) do
				local nl, nr = l-dx+1, r+dx+1
				if line:sub(nl,nl) == line:sub(nr,nr) then
					count = count + 1
				end
			end
		end
		local maxCount = sRange * #grid
		if count == maxCount then
			p1 = mirror
		elseif count == maxCount - 1 then
			p2 = mirror
		end
	end

	return p1, p2
end

local function checkUpDown(grid)
	local p1, p2 = 0, 0

	for mirror = 1, #grid-1 do
		local sRange = math.min(mirror, #grid-mirror)
		local up, dn = mirror-1, mirror

		local count = 0

		for dy = 0, sRange-1 do
			local nUp, nDn = up-dy+1, dn+dy+1
			local lineUp, lineDown = grid[nUp], grid[nDn]

			for i = 1, #lineUp do
				local c1, c2 = lineUp:sub(i,i), lineDown:sub(i,i)
				if c1 == c2 then
					count = count + 1
				end
			end
		end

		local maxCount = sRange * #grid[1]
		if count == maxCount then
			p1 = mirror
		elseif count == maxCount - 1 then
			p2 = mirror
		end
	end
	return p1, p2
end

local function checkGrid(grid)
	local lr1, lr2 = checkLeftRight(grid)
	local ud1, ud2 = checkUpDown(grid)
	return lr1 + ud1 * 100, lr2 + ud2 * 100
end

local function solve(grids)
	local p1, p2 = 0, 0
	for _, grid in ipairs(grids) do
		local res1, res2 = checkGrid(grid)
		p1 = p1 + res1
		p2 = p2 + res2
	end
	return p1, p2
end

local grids = {{}}
for line in io.lines('13-test.txt') do
	if #line == 0 then
		table.insert(grids, {})
	else
		local grid = grids[#grids]
		-- do this if you want an array of chars,
		-- rather than a string for each row:
		-- local arr = {}
		-- for ch in line:gmatch'.' do
			-- table.insert(arr, ch)
		-- end
		-- table.insert(grid, arr)
		table.insert(grid, line)
	end
end


log.report('%d %d\n', solve(grids))	-- 405, 400

grids = {{}}
for line in io.lines('13-input.txt') do
	if #line == 0 then
		table.insert(grids, {})
	else
		local grid = grids[#grids]
		-- do this if you want an array of chars,
		-- rather than a string for each row:
		-- local arr = {}
		-- for ch in line:gmatch'.' do
			-- table.insert(arr, ch)
		-- end
		-- table.insert(grid, arr)
		table.insert(grid, line)
	end
end

log.report('%d %d\n', solve(grids)) -- 34911, 33183
