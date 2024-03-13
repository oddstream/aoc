-- https://adventofcode.com/2015/day/4

local log = require 'log'
local md5 = require 'md5'

local input = 'iwrupvqb'

local function run(start, n)
	local zeros = string.rep('0', n)
	for i = start, 10000000 do
		local inputd = input .. tostring(i)
		local hash = md5.sumhexa(inputd)
		if hash:sub(1,n) == zeros then
			return i
		end
	end
	return -1
end

log.report('%s\n', _VERSION)

local result = run(0, 5)
log.report('part one %d\n', result)
result = run(result + 1, 6)
log.report('part two %d\n', result)

--[[
$ time luajit 04.lua
Lua 5.1
part one 346386
part two 9958218

real	1m2.711s
user	1m2.710s
sys	0m0.000s
]]