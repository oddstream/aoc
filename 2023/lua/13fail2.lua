-- https://adventofcode.com/2023/day/13 Point of Incidence

local log = require 'log'

local function rotate_grid(grid)
    local rotated_grid = {}
    for i = 1, #grid[1] do
        local row = {}
        for j = #grid, 1, -1 do
            row[#row + 1] = grid[j]:sub(i, i)
        end
        rotated_grid[#rotated_grid + 1] = table.concat(row)
    end
    return rotated_grid
end

local function rotate_grid2(grid)
    local rotated_grid = {}
    for i = 1, #grid[1] do
        local row = {}
        for j = #grid, 1, -1 do
            row[#row + 1] = grid[j][i]
        end
        rotated_grid[#rotated_grid + 1]	= row -- = table.concat(row)
    end
    return rotated_grid
end

--------------------------------------------------------------------------------
-- table.zip: zips given tables
-- table.zip({'a', 'b'}, {3, 4, 5})
-- result {{'a', 3}, {'b', 4}}
--
-- arguments: table(s)
-- return table with tables of which the size is of the smalles given table
--------------------------------------------------------------------------------
function _G.table.zip(...)
	local idx, ret, args = 1, {}, {...}
	while true do -- loop smallest table-times
		local sub_table = {}
		local value
		for _, table_ in ipairs(args) do
			value = table_[idx] -- becomes nil if index is out of range
			if value == nil then break end -- break for-loop
			table.insert(sub_table, value)
		end
		if value == nil then break end -- break while-loop
		table.insert(ret, sub_table) -- insert the sub result
		idx = idx + 1
	end
	return ret
end

local grids = {{}}
for line in io.lines('13-test.txt') do
	if #line == 0 then
		table.insert(grids, {})
	else
		local grid = grids[#grids]
		local arr = {}
		for ch in line:gmatch'.' do
			table.insert(arr, ch)
		end
		table.insert(grid, arr)
	end
end

local s

local function f(grid)
	for i = 1, #grid do
		local sum = 0
		-- compare each pair of adjacent chars
		-- If the number of positions where the characters in the two strings differ
		-- is equal to the variable s, the index of the current string is returned.
		-- If no such index is found, the function returns 0.
	end
	return 0
end

--[[
local function f(p)
    for i = 1, #p do
        local l = p[i-1] or ""
        local m = p[i]
        local count = 0
        for c, d in _G.table.zip(l:reverse(), m) do
            if c ~= d then
                count = count + 1
            end
        end
        if count == s then
            return i
        end
    end
    return 0
end
]]

s = 0
-- for _, grid in ipairs(grids) do
-- 	local sum = 100 * f(grid) + f(rotate_grid(grid))
-- 	print(sum)
-- end

s = 1

-- local rot = rotate_grid2(grids[1])
-- log.list(grids[1])
-- log.list(rot)

-- local tz = _G.table.zip({'a', 'b'}, {3, 4, 5})
-- print(#tz)
-- result {{'a', 3}, {'b', 4}}