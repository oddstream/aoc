-- https://adventofcode.com/2023/day/1

local log = require 'log'
dofile 'strings.lua'

local function partOne(filename)
	-- remove all non-numeric characters
	-- make the first and last character into a number
	local result = 0
	for line in io.lines(filename) do
		local nums = {}
		for num in line:gmatch'%d' do
			table.insert(nums, num)
		end
		result = result + tonumber(nums[1] .. nums[#nums])
	end
	return result
end

local function partTwo(filename)
	local result = 0
	local numbers = {'one', 'two', 'three', 'four', 'five', 'six', 'seven', 'eight', 'nine'}
	for line in io.lines(filename) do
		local num = ""
		local tmp = line -- make copy because we will destroy original
		while #tmp > 0 do
			local c = tmp:sub(1,1) -- first character
			if c >= '0' and c <= '9' then
				num = num .. c
				break
			end
			for k, v in ipairs(numbers) do
				if tmp:startswith(v) then
					num = num .. k
					break
				end
			end
			if #num > 0 then
				break
			end
			tmp = tmp:sub(2, #tmp) -- remove first character
		end
		while #line > 0 do
			local c = line:sub(#line)	-- last character
			if c >= '0' and c <= '9' then
				num = num .. c
				break
			end
			for k, v in ipairs(numbers) do
				if line:endswith(v) then
					num = num .. k
					break
				end
			end
			if #num > 1 then
				break	-- we found second (last) number
			end
			line = line:sub(1, -2) -- remove last character
		end
		result = result + tonumber(num)
	end
	return result
end

log.report('%s\n', _VERSION)
log.info('part one test %d\n', partOne('01-test1.txt'))
log.info('part one      %d\n', partOne('01-input.txt'))
log.info('part two test %d\n', partTwo('01-test2.txt'))
log.info('part two      %d\n', partTwo('01-input.txt'))
