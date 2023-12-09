-- https://adventofcode.com/2023/day/9 Mirage Maintenance

local log = require 'log'

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
	return out
end

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
		until #nums == 0 or all0(nums)

		table.insert(history[#history], 0)

		for i = #history, 2, -1 do
			local numsAbove = history[i-1]
			local numsHere = history[i]
			local newValue = numsHere[#numsHere] + numsAbove[#numsAbove]
			table.insert(history[i-1], newValue)
		end

		local firstNums = history[1]
		result = result + firstNums[#firstNums]
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
		until #nums == 0 or all0(nums)

		for i = 1, #history do
			table.insert(history[i], 1, 0)
		end

		for i = #history, 2, -1 do
			local numsAbove = history[i-1]
			local numsHere = history[i]
			local newValue = numsAbove[2] - numsHere[1]
			history[i-1][1] = newValue
		end

		local firstNums = history[1]
		result = result + firstNums[1]
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

real    0m0.040s
user    0m0.036s
sys     0m0.002s
]]