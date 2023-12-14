-- https://adventofcode.com/2023/day/13 Point of Incidence

local log = require 'log'

local function strsame(a,b)
	local n = math.min(#a, #b)
	for i = 1, n do
		if a:sub(i,i) ~= b:sub(i,i) then
			return false
		end
	end
	return true
end

local function strdiffs(a,b)
	assert(#a)
	assert(#b)
	-- assert(#a==#b)
	local diffs = 0
	local i = 1
	while i <= #a and i <= #b do
		if a:sub(i,i) ~= b:sub(i,i) then
			diffs = diffs + 1
		end
		i = i + 1
	end
	return diffs
end

local function mirrordiffs(a,b)
	assert(#a)
	assert(#b)
	local diffs = 0
	local ia = 1
	local ib = #b
	while ia <= #a and ib > 1 do
		if a:sub(ia,ia) ~= b:sub(ib,ib) then
			diffs = diffs + 1
		end
		ia = ia + 1
		ib = ib - 1
	end
	return diffs
end

local difference = 0

local function testLeftRight(grid, x)
	for y = 1, #grid do
		local row = grid[y]
		-- left and right should be same length, yes?
		local left = row:sub(1, x)
		local right = row:sub(x+1, x+#row)
		if strdiffs(left:reverse(), right) ~= difference then
			return false
		end
	end
	return true
end

local function testTopBottom(grid, y)
	for x = 1, #grid[1] do
		local col = '' -- column that is x across
		for i = 1, #grid do
			col = col .. grid[i]:sub(x,x)
		end
		local top = col:sub(1, y)
		local bottom = col:sub(y+1, #col)
		if strdiffs(top:reverse(), bottom) ~= difference then
			return false
		end
	end
	return true
end

local function calc(grid)
	local res = -1
	-- print('grid:')
	-- for _, y in ipairs(grid) do
	-- 	print(y)
	-- end

	for x = 1, #grid[1] - 1 do
		if testLeftRight(grid, x) then
			-- print('found x', x)
			res = x
			break
		end
	end
	if res == -1 then
		-- print('x not found looking for y')
		for y = 1, #grid - 1 do
			if testTopBottom(grid, y) then
				-- print('found y', y)
				res = y * 100
				break
			end
		end
	end
	return res
end

--[=[
local function flip0(str, x)
	local ch = str:sub(x,x)
	if ch == '.' then ch = '#' else ch = '.' end
	return str:sub(1, x-1) .. ch .. str:sub(x+1, #str)
end

local function flip(grid, x, y)
	grid[y] = flip0(grid[y], x)
end

local function calc2(grid)
	local old_result = calc(grid)
	for y = 1, #grid do
		for x = 1, #grid[1] do
			flip(grid, x,y)
			-- the smudge does not necessarily make the original result invalid
			-- so may need to keep looking
			local new_result = calc(grid)
			flip(grid, x,y)
			if new_result ~= -1 and new_result ~= old_result then
				return new_result
			end
		end
	end
	return 0
end
]=]

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	difference = 0
	local result = 0

	local grids = {{}}
	for line in io.lines(filename) do
		if #line == 0 then
			table.insert(grids, {})
		else
			local grid = grids[#grids]
			-- local arr = {}
			-- for ch in line:gmatch'.' do
				-- table.insert(arr, ch)
			-- end
			-- table.insert(grid, arr)
			table.insert(grid, line)
		end
	end
	for _, grid in ipairs(grids) do
		result = result + calc(grid)
	end

	if expected ~= nil and result ~= expected then
		log.error('part one should be %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	difference = 1
	local result = 0

	local grids = {{}}
	for line in io.lines(filename) do
		if #line == 0 then
			table.insert(grids, {})
		else
			local grid = grids[#grids]
			-- local arr = {}
			-- for ch in line:gmatch'.' do
				-- table.insert(arr, ch)
			-- end
			-- table.insert(grid, arr)
			table.insert(grid, line)
		end
	end
	for _, grid in ipairs(grids) do
		result = result + calc(grid)
	end

	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('13-test.txt', 405))
-- log.report('part one      %d\n', partOne('13-input.txt', 34911))
log.report('part two test %d\n', partTwo('13-test.txt', 400))
-- log.report('part two      %d\n', partTwo('13-input.txt', 0))

-- do
-- 	local str = '...###'
-- 	for i=1, #str do
-- 		print(flip0(str, i))
-- 	end
-- end