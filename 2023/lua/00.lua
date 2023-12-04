-- https://adventofcode.com/2023/day/

local log = require 'log'


---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	-- local f = assert(io.open(filename, "r"))
	-- local content = f:read("*all")
	-- f:close()
	for line in io.lines(filename) do
		log.trace(line)
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

	-- local f = assert(io.open(filename, "r"))
	-- local content = f:read("*all")
	-- f:close()
	for line in io.lines(filename) do
		log.trace(line)
	end

	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('00-test.txt', 0))
log.report('part one      %d\n', partOne('00-input.txt', 0))
log.report('part two test %d\n', partTwo('00-test.txt', 0))
log.report('part two      %d\n', partTwo('00-input.txt', 0))

