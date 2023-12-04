-- https://adventofcode.com/2015/day/1

local log = require 'log'

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local f = assert(io.open(filename, "r"))
	local content = f:read("*all")
	f:close()

	local result
	local level = 0
	for i = 1, #content do
		local c = content:sub(i,i)
		if c == '(' then
			level = level + 1
		elseif c == ')' then
			level = level - 1
		else
			log.error('unexpected character \'%s\' at position %d\n', c, i)
		end
	end
	result = level
	if expected ~= nil and result ~= expected then
		log.error('part one should be %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local f = assert(io.open(filename, "r"))
	local content = f:read("*all")
	f:close()

	local result = 0
	local level = 0
	for i = 1, #content do
		local c = content:sub(i,i)
		if c == '(' then
			level = level + 1
		elseif c == ')' then
			level = level - 1
		else
			log.error('unexpected character \'%s\' at position %d\n', c, i)
		end
		if level < 0 then
			result = i
			break
		end
	end
	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('part one %d\n', partOne('01-input.txt', 138))
log.report('part two %d\n', partTwo('01-input.txt', 1771))
