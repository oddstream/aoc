-- https://adventofcode.com/2023/day/12 Hot Springs

-- largest number in input's damaged list = 16
-- longest run of consecutive ? is ????????????????? = 17

-- the numbers could be expanded into regex (e.g. 1,1,3 == .*#.+#.+###.*) b

-- see https://old.reddit.com/r/adventofcode/comments/18hbbxe/2023_day_12python_stepbystep_tutorial_with_bonus/
-- https://old.reddit.com/r/adventofcode/comments/18hbjdi/2023_day_12_part_2_this_image_helped_a_few_people/?utm_source=reddit&utm_medium=usertext&utm_name=adventofcode&utm_content=t1_kd5xzfs

local log = require 'log'

---@param i integer the number to turn into a 'binary' string
---@param n integer the maximum + 1 of i
---@return string
local function num2Bits(i, n)
	local str = ''
	local b = 1
	while b < n do
		if i & b == b then
			str = '#' .. str
		else
			str = '.' .. str
		end
		b = b << 1
	end
	return str -- nb caching this INCREASED the run time
end

---@param nqueries integer number of ?
---@return function (an iterator)
local function permutations(nqueries)

	local function permute()
		for i = 0, (2 ^ nqueries) - 1 do
			coroutine.yield(num2Bits(i, 2 ^ nqueries))
		end
	end

	local co = coroutine.create(function() permute() end)
	return function()
		local code, res = coroutine.resume(co)
		if code then return res end
	end
end

---@param str string eg '..###.#.##.'
---@return string eg '3,1,2'
local function hashruns2nums(str)
	local out = ''
	local count = 0
	for c in str:gmatch'.' do
		if c == '#' then
			count = count + 1
		else
			if count > 0 then
				if #out > 0 then out = out .. ',' end
				out = out .. tostring(count)
				count = 0
			end
		end
	end
	if count > 0 then
		if #out > 0 then out = out .. ',' end
		out = out .. tostring(count)
	end
	return out
end

--[[
local permCache = {}
local function createPermCache()
	for i = 1, 17 do
		local arr = {}
		for perm in permutations(i) do
			local test = ('%s%s%s'):format(
				perm:sub(1, i-1),
				perm,
				perm:sub(i+1, #springs))
			testperms(test, damaged)

			table.insert(arr, perm)
		end
		permCache[i] = arr
	end
end
]]

---@param filename string
---@param reps integer
---@return integer
local function solve(filename, reps)
	local result = 0

	local function testperms(springs, damaged)
		local start, stop = string.find(springs, '%?+')
		if start ~= nil then
			-- local queries = springs:sub(start, stop)
			-- turn a sequence like ???? into all permutations
			-- eg ?? := '..', '##', '#.', '.#'
			-- the generated permutations could be precomputed
			-- or cached as they apply to all input lines
			for perm in permutations(stop - start + 1) do
				-- local test = ('%s%s%s'):format(
				-- 	springs:sub(1, start-1),
				-- 	perm,
				-- 	springs:sub(stop+1, #springs))
				local test = springs:sub(1, start-1) .. perm .. springs:sub(stop+1, #springs)
				testperms(test, damaged)
			end
		else
			local nums = hashruns2nums(springs)
			if nums == damaged then
				result = result + 1
				-- print('    match', springs, nums, damaged)
			-- else
				-- print('not match', springs, nums, damaged)
			end
		end
	end

	-- local lineno = 1
	for line in io.lines(filename) do
		local springs, damaged = line:match'([%?%.%#]+) ([%d,]+)'
		-- log.trace('springs=%s damaged=%s\n', springs, damaged)
		springs = string.rep(springs, reps, '?')
		damaged = string.rep(damaged, reps, ',')
		-- lineno = lineno + 1
		-- log.trace('line %d\n', lineno)
		testperms(springs, damaged)
	end

	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = solve(filename, 1)
	if expected ~= nil and result ~= expected then
		log.error('part one should be %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result = solve(filename, 2)
	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
-- createPermCache()
log.report('part one test  %d\n', partOne('12-test.txt', 21))
log.report('part one      %d\n', partOne('12-input.txt', 7236))
-- log.report('part two test %d\n', partTwo('12-test.txt', 525152))
-- log.report('part two      %d\n', partTwo('12-input.txt', 11607695322318))