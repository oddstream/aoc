-- https://adventofcode.com/2023/day/8 Haunted Wasteland

local log = require 'log'

---@param filename string
---@return string, table[]table
local function loadInput(filename)
	local directions = ''
	local map = {}
	local file, err = io.open(filename, 'r')
	if file == nil then
		log.error('cannot open %s - %s\n', filename, err)
	else
		directions = file:read()	-- default is to read a line, skipping the EOL

		local line = file:read()
		while line == '' do
			line = file:read()
		end

		while line ~= nil do
			local start, left, right = line:match'(%w%w%w)[%W]+(%w%w%w)[%W]+(%w%w%w)'
			if start ~= nil and left ~= nil and right ~= nil then
				if map[start] ~= nil then
					log.error('repeat %s\n', start)
				end
				map[start] = {L=left, R=right}
			else
				log.error('unknown line %s\n', line)
			end

			line = file:read()
		end
		assert(io.close(file))
	end

	return directions, map
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	local dirs, mp = loadInput(filename)
	-- "You feel like AAA is where you are now,
	-- and you have to follow the left/right instructions until you reach ZZZ."
	-- "feel like"?
	-- brief time-wasting episode when I thought that alpha and omega might be
	-- the first and last entries in the input, viz DGK and TLL
	local pos = 'AAA'
	repeat
		for dir in dirs:gmatch'.' do
			-- if mp[pos] == nil or mp[pos][dir] == nil then
			-- 	log.error('no map entry for %s %s\n', pos, dir)
			-- 	break
			-- end
			pos = mp[pos][dir]
			result = result + 1
			if pos == 'ZZZ' then
				break
			end
		end
		-- if mp[pos] == nil or mp[pos][dir] == nil then
		-- 	log.error('no map entry for %s %s\n', pos, dir)
		-- 	break
		-- end
		until pos == 'ZZZ'

	if expected ~= nil and result ~= expected then
		log.error('part one should be %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result

	---@param a integer
	---@param b integer
	---@return integer
	local function gcd(a, b)
		-- Euclidean Algorithm
		return a == 0 and b or gcd(b % a, a)
	end

	---@param a integer
	---@param b integer
	---@return integer
	local function lcm(a, b)
		return a * b / gcd(a, b)
	end

	---@param arr integer[]
	---@return integer
	local function lcm_array(arr)
		if _VERSION == 'Lua 5.1' then
			return #arr == 1 and arr[1] or lcm(arr[1], lcm_array{unpack(arr, 2)})
		else
			return #arr == 1 and arr[1] or lcm(arr[1], lcm_array{table.unpack(arr, 2)})
		end
	end

	local dirs, mp = loadInput(filename)

	-- brute force wasn't going to work (result is 11,283,670,395,017)
	-- so found the LCM hint on the AOC solutions subreddit (thank you kind strangers)
	-- neat recursive LCM array function from Bing Chat
	local starts = {}
	for k, _ in pairs(mp) do
		if k:sub(3,3) == 'A' then
			table.insert(starts, k)
		end
	end
	-- test3 starts := 22A, 11A
	-- input starts := XSA, TTA, VVA, MHA, NBA, AAA
	local arr = {}
	for _, pos in ipairs(starts) do
		local n = 0
		repeat
			for dir in dirs:gmatch'.' do
				pos = mp[pos][dir]
				n = n + 1
			end
		until pos:sub(3,3) == 'Z'
		table.insert(arr, n)
	end
	-- test3 := 2 6
	-- input := 21251 12643 19637 15871 11567 16409
	result = lcm_array(arr)

	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('08-test1.txt', 2))
log.report('part one test %d\n', partOne('08-test2.txt', 6))
log.report('part one      %d\n', partOne('08-input.txt', 15871))
log.report('part two test %d\n', partTwo('08-test3.txt', 6))
log.report('part two      %d\n', partTwo('08-input.txt', 11283670395017))

--[[
$ time luajit 08.lua
Lua 5.1
part one test 2
part one test 6
part one      15871
part two test 6
part two      11283670395017

real    0m0.042s
user    0m0.038s
sys     0m0.004s
]]