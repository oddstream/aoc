local log = require 'log'

-- the input is 140x140 characters
-- the input DOES contain line duplicates eg line 20 has two 20s

local function loadEngine(filename)
	local engine = {}
	for line in io.lines(filename) do
		table.insert(engine, line)
	end
	return engine
end

local function mapSymbols(engine)
	local symbols = {}
	local function isSymbol(ch)
		return not(ch == '.' or (ch >= '0' and ch <= '9'))
	end
	for y, line in ipairs(engine) do
		for x=1, #line do
			local ch = line:sub(x,x)
			if isSymbol(ch) then
				symbols[x .. ',' .. y] = ch
			end
		end
	end
	return symbols
end

local function isPosAdjacentToSymbol(symbols, x, y)
	local deltas = {-1, 0, 1}
	for _, col in ipairs(deltas) do
		for _, row in ipairs(deltas) do
			local key = x+col .. ',' .. y+row
			if symbols[key] ~= nil then
				return true
			end
		end
	end
	return false
end

local function isNumAdjacentToSymbol(symbols, num, x, y)
	for i = 0, #num-1 do
		if isPosAdjacentToSymbol(symbols, x + i, y) then
			return true
		end
	end
	return false
end

local function getSurroundingNumbers(numbers, x, y)
	local numMap = {}
	local deltas = {-1, 0, 1}
	for _, col in ipairs(deltas) do
		for _, row in ipairs(deltas) do
			local key = x+col .. ',' .. y+row
			if numbers[key] ~= nil then
				numMap[numbers[key]] = true
			end
		end
	end
	local numList = {}
	for num, _ in pairs(numMap) do
		table.insert(numList, num)
	end
	return numList
end

local function partOne(filename, expected)
	local result = 0
	local engine = loadEngine(filename)
	local symbols = mapSymbols(engine)
	for y, line in ipairs(engine) do
		-- () empty capture == current string position
		for x, num in string.gmatch(line, "()(%d+)") do
			if isNumAdjacentToSymbol(symbols, num, x, y) then
				result = result + tonumber(num)
			end
		end
	end
	if expected ~= nil and result ~= expected then
		log.error('part one should be %d\n', expected)
	end
	return result
end

local function partTwo(filename, expected)
	local result = 0
	local engine = loadEngine(filename)
	local numberMap = {}
	-- build a map of all the digit positions
	for y, line in ipairs(engine) do
		-- () empty capture == current string position
		for x, num in string.gmatch(line, "()(%d+)") do
			for i = 0, #num - 1 do
				local key = x + i ..',' .. y
				numberMap[key] = num
			end
		end
	end
	-- for each of the gear '*' positions ...
	for y, line in ipairs(engine) do
		-- () empty capture == current string position
		for x, _ in string.gmatch(line, "()(%*)") do
			local nums = getSurroundingNumbers(numberMap, x, y)
			if #nums == 2 then
				local ratio = tonumber(nums[1]) * tonumber(nums[2])
				result = result + ratio
			end
		end
	end
	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('part one test %d\n', partOne('03-test.txt', 4361))
log.report('part one      %d\n', partOne('03-input.txt', 535078))
log.report('part two test %d\n', partTwo('03-test.txt', 467835))
log.report('part two      %d\n', partTwo('03-input.txt', 75312571))
-- wasted 3 hours because engine and symbols were not being reset
-- between running test and input
-- should have spotted earlier that result was 4361 too high

--[[
$ time ~/lua-5.1.5/src/lua 03.lua
part one test 4361
part one      535078
part two test 467835
part two      75312571

real    0m0.101s
user    0m0.090s
sys     0m0.012s

$ time ~/lua-5.4.3/src/lua 03.lua
part one test 4361
part one      535078
part two test 467835
part two      75312571

real    0m0.068s
user    0m0.064s
sys     0m0.004s
]]