-- https://adventofcode.com/2019/day/18

-- https://github.com/mnml/aoc/blob/main/2019/18/1.go

local log = require 'log'

dofile 'permutations.lua'

local function key(y,x) return y * 100 + x end
local function unkey(n)	return n // 100, n % 100 end

local function load(fname)
	local grid, keys, doors = {}, {}, {}
	local start = {}
	for line in io.lines(fname) do
		local row = {}
		for x = 1, #line do
			local ch = line:sub(x,x)
			if not (ch == '.' or ch == '#') then
				if ch == '@' then
					start = {y=#grid+1, x=x}
				elseif string.match(ch, '%l') ~= nil then
					keys[key(#grid+1,x)] = ch
				elseif string.match(ch, '%u') ~= nil then
					doors[key(#grid+1,x)] = ch
				end
				ch = '.'
			end
			row[#row+1] = ch
		end
		grid[#grid+1] = row
	end
	return grid, keys, doors, start
end

local function display(grid)
	for y = 1, #grid do
		for x = 1, #grid[y] do
			io.write(grid[y][x])
		end
		io.write('\n')
	end
end

local directions = {
	{x=1, y=0},
	{x=0, y=1},
	{x=0, y=-1},
	{x=-1, y=0},
}

---@return table?
local function bfs(grid, keys, doors, start)
	local visibles = {}
	start.steps = 0
	local q = {start}
	local seen = {[key(start.y, start.x)]=true}
	while #q > 0 do
		local p = table.remove(q, 1)
		local k = key(p.y, p.x)
		if keys[k] ~= nil then
			-- log.report('found key %s at %d,%d, %d steps\n', keys[k], p.y, p.x, p.steps)
			visibles[#visibles+1] = {y=p.y, x=p.x, key=keys[k], steps=p.steps}
			-- return {y=p.y, x=p.x, steps=p.steps, key=keys[k]}
		end

		for _, dir in ipairs(directions) do
			local ny, nx = p.y + dir.y, p.x + dir.x
			local nk = key(ny,nx)
			if not seen[nk] then
				seen[nk] = true
				local nch = grid[ny][nx]
				if nch == '.' and doors[nk] == nil then
					q[#q+1] = {y=ny, x=nx, steps=p.steps + 1}
				end
			end
		end
	end
	return visibles
end

local function calcDistances(grid, keys)

	local function dist(grid, start, stop)
		start.steps = 0
		local q = {start}
		local seen = {[key(start.y, start.x)]=true}
		while #q >0 do
			local p = table.remove(q, 1)
			if p.y == stop.y and p.x == stop.x then
				return p.steps
			end
			for _, dir in ipairs(directions) do
				local ny, nx = p.y + dir.y, p.x + dir.x
				local nk = key(ny,nx)
				if not seen[nk] then
					seen[nk] = true
					local nch = grid[ny][nx]
					if nch == '.' then
						q[#q+1] = {y=ny, x=nx, steps=p.steps + 1}
					end
				end
			end
		end
		return -1
	end

	local distances = {}	-- distances['az'] = 32

	for ki, i in pairs(keys) do
		for kj, j in pairs(keys) do
			if i ~= j then
				local x, y
				x, y = unkey(ki)
				local start = {x=x, y=y}
				x, y = unkey(kj)
				local stop = {x=x, y=y}
				local d = dist(grid, start, stop)
				if d == -1 then print(i, j, 'not found') break end
				distances[i .. j] = d
				-- distances[j .. i] = d
			end
		end
	end

	return distances
end

--[[
-- xkvncotfh are initially visible
-- 26! permutations is not brute-forceable!
]]
-- 5034 too high - key order xkvncotfhaegqyzbusdiwpjlrm
-- 4526 too high
local grid, gkeys, gdoors, gstart = load('18-input.txt')
local distances = calcDistances(grid, gkeys)
print('a', 'z', distances['az'])
print('z', 'a', distances['za'])

--[[
local result = 0
while true do
	local visibles = bfs(grid, gkeys, gdoors, gstart)
	if visibles == nil or #visibles == 0 then break end
	-- assume visibles is sorted with smallest steps first
	local v = visibles[1]
	io.write(v.key)
	-- remove key and door
	gkeys[key(v.y,v.x)] = nil
	for dk, dv in pairs(gdoors) do
		if dv == string.upper(v.key) then
			gdoors[dk] = nil
		end
	end
	-- go find next
	gstart = {y=v.y, x=v.x}
	result = result + v.steps
end
log.report('Part One %d\n', result)
]]


local function distanceToCollectKeys(currentKey, keys, cache)
	if #keys == 0 then
		return 0
	end
	if cache[currentKey .. ',' .. keys] ~= nil then
		return cache[currentKey .. ',' .. keys]
	end
	local result = 32767
	for _, vkey in ipairs(bfs(grid, gkeys, gdoors, gstart)) do
		local subkeys = ''
		for i=1, #keys do
			if keys:sub(i,i) ~= vkey.key then
				subkeys = subkeys .. vkey.key
			end
			local d = distances[currentKey .. vkey.key] + distanceToCollectKeys(vkey.key, subkeys, cache)
			if d < result then result = d end
		end
	end
	cache[currentKey .. ',' .. keys] = result
	return result
end

local allkeys = ''
for _, kv in pairs(gkeys) do
	allkeys = allkeys .. kv
end

print(distanceToCollectKeys('x', allkeys, {}))

--[[

distanceToCollectKeys(currentKey, keys):

    if keys is empty:
        return 0

    result := infinity
    foreach key in reachable(keys):
       d := distance(currentKey, key) + distanceToCollectKeys(key, keys - key)
       result := min(result, d)

    return result;

distanceToCollectKeys(currentKey, keys, cache):

    if keys is empty:
        return 0

    cacheKey := (currentKey, keys)
    if cache contains cacheKey:
        return cache[cacheKey]

    result := infinity
    foreach key in reachable(keys):
       d := distance(currentKey, key) + distanceToCollectKeys(key, keys - key, cache)
       result := min(result, d)

    cache[cacheKey] := result
    return result
]]

local k = key(32, 42)
print(k)
print(unkey(k))
