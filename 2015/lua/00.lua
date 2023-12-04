-- https://adventofcode.com/2023/day/

local log = require 'log'

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0
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
	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('part one test %d\n', partOne('04-test.txt', 0))
log.report('part one      %d\n', partOne('04-input.txt', 0))
log.report('part two test %d\n', partTwo('04-test.txt', 0))
log.report('part two      %d\n', partTwo('04-input.txt', 0))

