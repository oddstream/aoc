-- https://adventofcode.com/2023/day/11 Cosmic Expansion

-- took a complete wrong turn with an actual grid (not needed!) and a BFS (not needed!)
-- nb LuaJIT is 10x faster than Lua 5.4
-- could speed it up by using sets for empty rows and cols, instead of a list

local log = require 'log'

local function distance(x1, y1, x2, y2)
	return math.abs(x1 - x2 + y1 - y2)
end

local function intersection(set, a, b)
	local function contains(n)
		for _, v in ipairs(set) do
			if v == n then return true end
		end
		return false
	end

	local count = 0
	for i = a, b do
		if contains(i) then count = count + 1 end
	end
	return count
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	local stars = {}	-- list of {row=, col=}
	local empty_rows = {} -- list of empty rows
	local empty_cols = {}	-- list of empty cols

	local row = 1
	local colstrs = {}
	for line in io.lines(filename) do
		if line:find'#' == nil then
			table.insert(empty_rows, row)
		else
			for col in line:gmatch'()#' do
				table.insert(stars, {row=row, col=tonumber(col)})
			end
		end
		row = row + 1

		for i = 1, #line do
			if colstrs[i] == nil then colstrs[i] = '' end
			colstrs[i] = colstrs[i] .. line:sub(i,i)
		end
	end

	for i = 1, #colstrs do
		if colstrs[i]:find'#' == nil then
			table.insert(empty_cols, i)
		end
	end

	-- for i = 1, #stars do
	-- 	print('star at', stars[i].row, stars[i].col)
	-- end

	local npairs = 0
	for i = 1, #stars do
		for j = i+1, #stars do
			local ax, ay = stars[i].col, stars[i].row
			local bx, by = stars[j].col, stars[j].row
            ax, bx = math.min(ax, bx), math.max(ax, bx)
            ay, by = math.min(ay, by), math.max(ay, by)
			local count = (bx - ax) + intersection(empty_cols, ax, bx)
			count = count + (by - ay) + intersection(empty_rows, ay, by)
			result = result + count
			npairs = npairs + 1
		end
	end
	log.trace('stars = %d, npairs = %d\n', #stars, npairs)


	if expected ~= nil and result ~= expected then
		log.error('part one should be %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result = 0

	local stars = {}	-- list of {row=, col=}
	local empty_rows = {} -- list of empty rows
	local empty_cols = {}	-- list of empty cols

	local row = 1
	local colstrs = {}
	for line in io.lines(filename) do
		if line:find'#' == nil then
			table.insert(empty_rows, row)
		else
			for col in line:gmatch'()#' do
				table.insert(stars, {row=row, col=tonumber(col)})
			end
		end
		row = row + 1

		for i = 1, #line do
			if colstrs[i] == nil then colstrs[i] = '' end
			colstrs[i] = colstrs[i] .. line:sub(i,i)
		end
	end

	for i = 1, #colstrs do
		if colstrs[i]:find'#' == nil then
			table.insert(empty_cols, i)
		end
	end

	-- for i = 1, #stars do
	-- 	print('star at', stars[i].row, stars[i].col)
	-- end

	local npairs = 0
	for i = 1, #stars do
		for j = i+1, #stars do
			local ax, ay = stars[i].col, stars[i].row
			local bx, by = stars[j].col, stars[j].row
            ax, bx = math.min(ax, bx), math.max(ax, bx)
            ay, by = math.min(ay, by), math.max(ay, by)
			local count = (bx - ax) + intersection(empty_cols, ax, bx) * (1000000-1)
			count = count + (by - ay) + intersection(empty_rows, ay, by) * (1000000-1)
			result = result + count
			npairs = npairs + 1
		end
	end
	log.trace('stars = %d, npairs = %d\n', #stars, npairs)


	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('11-test.txt', 374))
log.report('part one      %d\n', partOne('11-input.txt', 9918828))
-- log.report('part two test %d\n', partTwo('11-test.txt', 0))
log.report('part two      %d\n', partTwo('11-input.txt', 692506533832))

--[[
$ time lua54 11.lua
Lua 5.4
stars = 440, npairs = 96580
part two      692506533832

real    0m2.468s
user    0m2.439s
sys     0m0.006s

$ time luajit 11.lua
Lua 5.1
stars = 440, npairs = 96580
part two      692506533832

real    0m0.285s
user    0m0.263s
sys     0m0.005s
]]