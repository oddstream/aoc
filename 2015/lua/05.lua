-- https://adventofcode.com/2015/day/5

local log = require 'log'

local function vowels(s)
	local count = 0
	for i = 1, #s do
		local c = s:sub(i,i)
		if c == 'a' or c == 'e' or c == 'i' or c == 'o' or c == 'u' then
			count = count + 1
		end
	end
	return count
end

local function containsRepeatedChar(s)
	local prev
	for i = 1, #s do
		local c = s:sub(i,i)
		if c == prev then
			return true
		end
		prev = c
	end
	return false
end

local function containsNaughtyString(s)
	for i = 1, #s-1 do
		local cc = s:sub(i,i+1)
		if cc == 'ab' or cc == 'cd' or cc == 'pq' or cc == 'xy' then
			return true
		end
	end
	return false
end

local function test21(s)
	for i = 1, #s-1 do
		for j = i + 2, #s-1 do
			if s:sub(j,j) == s:sub(i,i) and s:sub(j+1,j+1) == s:sub(i+1,i+1) then
				return true
			end
		end
	end
	return false
end

local function test22(s)
	local prev1, prev2
	for i = 1, #s do
		local c = s:sub(i,i)
		if c == prev2 then
			return true
		end
		prev2 = prev1
		prev1 = c
	end
	return false
end

local result1, result2 = 0, 0
for line in io.lines('05-input.txt') do
	if not containsNaughtyString(line) then
		if vowels(line) >= 3 then
			if containsRepeatedChar(line) then
				result1 = result1 + 1
			end
		end
	end
	if test21(line) and test22(line) then
		result2 = result2 + 1
	end
end
log.report('part one %d\n', result1)
log.report('part two %d\n', result2)

--[[
$ time luajit 05.lua
part one 255
part two 55

real	0m0.003s
user	0m0.003s
sys	0m0.000s
]]