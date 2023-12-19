-- https://adventofcode.com/2023/day/17 Clumsy Crucible

local pq = require 'pq'
local log = require 'log'

-- test input is 13x13
-- input is 141x141

local function loadGrid(filename)
	local grid = {}
	for line in io.lines(filename) do
		local arr = {}
		for ch in line:gmatch'.' do
			arr[#arr+1] = tonumber(ch)
		end
		grid[#grid+1] = arr
	end
	return grid
end

local function displayGrid(grid)
	local H = #grid
	local W = #grid[1]
	for i = 1, H do
		for j = 1, W do
			if grid[i][j] ~= math.huge then
				io.write(('%3d '):format(grid[i][j]))
			else
				io.write('  - ')
			end
		end
		io.write('\n')
	end
end

--[[
	This implementation uses a heap to keep track of the minimum distance to each cell.
	The heap is initialized with the starting cell (1, 1) and its distance.
	The algorithm then iteratively extracts the cell with the minimum distance from the heap
		and updates the distances of its neighbors.
	The algorithm terminates when the bottom-right cell is extracted from the heap.
]]
---Sounds like dyke-stra
---@param grid table[]
---@params mindist integer
---@params maxdist integer
---@return integer
local function dijkstra(grid, mincount, maxcount)
	-- each block in the grid contains a
	-- "single digit that represents the amount of heat loss if the crucible enters that block"
	local H = #grid
	local W = #grid[1]

	local function ingrid(y,x)
		return x >= 1 and y >= 1 and x <= W and y <= H
	end

	local function manhatten(y, x)
		return H-y + W-x
	end

	-- marks the distance (from the starting point) to every other intersection on the map with infinity.
	-- this is done not to imply that there is an infinite distance,
	-- but to note that those intersections have not been visited yet.
	local dist = {}
	for i = 1, H do
		dist[i] = {}
		for j = 1, W do
			dist[i][j] = math.huge
		end
	end

	-- "Because you already start in the top-left block,
	-- you don't incur that block's heat loss unless you leave that block and then return to it."
	-- dist[1][1] = grid[1][1]
	-- dist[1][1] = 0

	-- local heap = {{1, 1, dist[1][1]}}
	-- local heap = {{1, 1, 0, 0, 0}} -- y, x, weight, dir, count
	local heap = {{1, 1, 0, 0, 1}}
	local loops = 0
	while #heap > 0 do
		loops = loops + 1
		local y, x, w, d, c = table.unpack(table.remove(heap, 1))
		if x == H and y == W then
			-- displayGrid(dist)
			log.info('%d loops (%d nodes)\n', loops, W*H)
			return dist[H][W]
		end
		if w > dist[y][x] then
			goto continue
		end
		local prohibited = {}
		for dirn, dir in ipairs({{x=0,y=-1},{x=1,y=0},{x=-1,y=0},{x=0,y=1}}) do	-- news
			if #prohibited == 2 and (dirn == prohibited[1] or dirn == prohibited[2]) then
				goto nextdir
			end
			-- The crucible also can't reverse direction;
			-- after entering each city block, it may only turn left, continue straight, or turn right.
			if dirn ~= d then
				if (dirn == 1 and d == 4) or (dirn == 4 and d == 1) then
					goto nextdir
				end
				if (dirn == 2 and d == 3) or (dirn == 3 and d == 2) then
					goto nextdir
				end
			end
			-- move at most three blocks in a single direction before it must turn 90 degrees left or right
			-- ie not the direction we were already going, or it's opposite
			prohibited = {}
			local newc
			if d ~= dirn then
				newc = 1
			else
				newc = c + 1
				if newc > maxcount then
					if d == 1 or d == 4 then	-- n or s
						prohibited = {1, 4}		-- must turn 90 deg, go e or w
					else						-- e or w
						prohibited = {2, 3}		-- must turn 90 deg, go n or s
					end
					goto nextdir
				end
			end

			local yy = y + dir.y
			local xx = x + dir.x
			if ingrid(yy,xx) then
				local alt = w + grid[yy][xx]
				if alt < dist[yy][xx] then
					dist[yy][xx] = alt
					table.insert(heap, {yy, xx, alt, dirn, newc})
					-- table.sort(heap, function(a,b) return a[3] < b[3] end)
					-- table.sort(heap, function(a,b) return manhatten(a[1], a[2]) < manhatten(b[1],b[2]) end)
					table.sort(heap, function(a,b)
						if a[3] == b[3] then
							return manhatten(a[1], a[2]) < manhatten(b[1], b[2])
						else
							return a[3] < b[3]
						end
					end)
				end
			end
			::nextdir::
		end
		::continue::
	end
	return 0
end

local function run(grid, mindist, maxdist)

	local H = #grid
	local W = #grid[1]
	local DIRS = { {0, 1}, {1, 0}, {0, -1}, {-1, 0} } -- eswn

	local function ingrid(y,x)
		return x >= 1 and y >= 1 and x <= W and y <= H
	end

	local function key(a, b, c)
		return ('%d,%d,%d'):format(a, b, c)
	end

	local function oppdir(d)
		if d == 1 then
			return 3
		elseif d == 2 then
			return 4
		elseif d == 3 then
			return 1
		elseif d == 4 then
			return 2
		end
		assert(false, 'oppdir bad input')
	end

	local q = {{0, 1, 1, 0}} -- cost, y 1..H, x 1..W, direction/dd 1..4
	local seen = {}		-- set
	local costs = {}	-- set
	while #q > 0 do
		local cost, y, x, dd = table.unpack(table.remove(q, 1))
		if y == H and x == W then
			return cost
		end
		if seen[key(y, x, dd)] == true then
			goto continue
		end
		seen[key(y, x, dd)] = true
		for direction = 1, 4 do
			local costincrease = 0
			if direction == dd or oppdir(direction) == dd then -- 94,1392
				goto nextdirection
			end
			for distance = 1, maxdist do
				local yy = y + DIRS[direction][1] * distance
				local xx = x + DIRS[direction][2] * distance
				if ingrid(yy, xx) then
					costincrease = costincrease + grid[yy][xx]
					if distance < mindist then
						goto nextdistance
					end
					local nc = cost + costincrease
					local co = costs[key(yy,xx,direction)]
					if co == nil then
						co = 1/0
					end
					if co <= nc then
						goto nextdistance
					end
					costs[key(yy,xx,direction)] = nc
					table.insert(q, {nc, yy, xx, direction})
					table.sort(q, function (a,b) return a[1] < b[1] end)
				end
::nextdistance::
			end
::nextdirection::
		end

::continue::
	end

	return 0
end

--[[
	The modified function adds two new parameters to the heap: c and d.
	c keeps track of the number of consecutive moves in the same direction,
	while d stores the direction of the previous move.
	The algorithm only considers neighbors that are not in the same direction as the previous move.
	If a new direction is taken, c is reset to 0, otherwise it is incremented by 1.
	The algorithm terminates when the bottom-right cell is extracted from the heap.
]]
---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result

	local grid = loadGrid(filename)
	-- result = dijkstra(grid, 1, 3)
	result = run(grid, 1, 3)

	-- local q = pq.new()
	-- q:add(4, 'four')
	-- q:add(-2, 'minus')
	-- q:add(3, 'three')
	-- q:add(0, 'zero')
	-- while q:len() > 0 do
	-- 	local p, v = q:pop()
	-- 	print(p, v)
	-- end
	-- result = 0

	if expected and result ~= expected then
		log.error('expecting %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result

	local grid = loadGrid(filename)
	-- result = dijkstra(grid, 4, 10)
	result = run(grid, 4, 10)

	if expected and result ~= expected then
		log.error('expecting %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
-- log.report('part one test %d\n', partOne('17-test.txt', 102))
-- log.report('part one      %d\n', partOne('17-input.txt', 1244))
log.report('part two test %d\n', partTwo('17-test.txt', 94))
log.report('part two      %d\n', partTwo('17-input.txt', 1367))
