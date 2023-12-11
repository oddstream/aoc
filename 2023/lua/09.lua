-- https://adventofcode.com/2023/day/9 Mirage Maintenance

-- nb solution could be recursive
-- and part 2 could use same function as part 1, but with a reversed array

local log = require 'log'

--[[
---@param x table list of things
local function reverse(x)
	local n, m = #x, #x/2
	for i=1, m do
		x[i], x[n-i+1] = x[n-i+1], x[i]
	end
	-- in-place reversal, nothing to return
end
]]

local function all0(nums)
	for _, v in ipairs(nums) do
		if v ~= 0 then
			return false
		end
	end
	return true
end

local function diffs(nums)
	local out = {}
	for i = 2, #nums do
		table.insert(out, nums[i] - nums[i-1])
	end
	-- out is never empty
	return out
end

--[[
local function process(history, esrever)
	if esrever == nil then esrever = false end
	-- add a 'placeholder' 0 to the end of all history rows
	for _, row in ipairs(history) do
		table.insert(row, 0)
	end
	if esrever then reverse(history) end
	for i = #history, 2, -1 do
		local numsAbove = history[i-1]
		local numsHere = history[i]
		local newValue = numsHere[#numsHere] + numsAbove[#numsAbove - 1]
		assert(numsAbove[#numsAbove] == 0)
		numsAbove[#numsAbove] = newValue
	end

	local firstNums = history[1]
	return firstNums[#firstNums]
end
]]

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	for line in io.lines(filename) do
		local nums = {}
		for num in line:gmatch'([%-%d]+)' do
			table.insert(nums, tonumber(num))
		end

		local history = {}
		table.insert(history, nums)
		repeat
			nums = diffs(nums)
			table.insert(history, nums)
		until all0(nums)

		-- just append a 'placeholder' to the last row of nums
		table.insert(history[#history], 0)

		for i = #history, 2, -1 do
			local numsAbove = history[i-1]
			local numsHere = history[i]
			local newValue = numsHere[#numsHere] + numsAbove[#numsAbove]
			table.insert(numsAbove, newValue)
		end

		result = result + history[1][#history[1]]	-- last number in first row
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
	local result = 0

	for line in io.lines(filename) do
		local nums = {}
		for num in line:gmatch'([%-%d]+)' do
			table.insert(nums, tonumber(num))
		end

		local history = {}
		table.insert(history, nums)
		repeat
			nums = diffs(nums)
			table.insert(history, nums)
		until all0(nums)

		table.insert(history[#history], 1, 0)

		for i = #history, 2, -1 do
			local numsAbove = history[i-1]
			local numsHere = history[i]
			local newValue = numsAbove[1] - numsHere[1]
			table.insert(numsAbove, 1, newValue)
		end

		result = result + history[1][1] -- first number in first row
	end

	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('09-test.txt', 114))
log.report('part one      %d\n', partOne('09-input.txt', 2043677056))
log.report('part two test %d\n', partTwo('09-test.txt', 2))
log.report('part two      %d\n', partTwo('09-input.txt', 1062))

--[[
$ time luajit 09.lua
Lua 5.1
part one test 114
part one      2043677056
part two test 2
part two      1062

real    0m0.008s
user    0m0.008s
sys     0m0.000s
]]