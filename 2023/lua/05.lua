-- https://adventofcode.com/2023/day/5

local log = require 'log'

-- seeds <list>
-- src-to-dst
-- dst-start src-start range-length
-- numbers are LARGE, so we won't be storing individual items, just ranges
-- maps: seed soil fertilizer water light temperature humidity location
-- map seed number (thru maps) to location number

local seeds -- list of seed numbers
-- list of maps 1 .. 7 don't care what they are called
-- eg maps[1] = {{range=2, soil=50, seed=98}, {range=48, soil=52, seed=50}}
-- eg maps[1] = {{50, 98, 2}, {52, 50, 48}}
local maps = {}

local function loadSeedsAndMaps(filename)
	local mapno = 0
	seeds = {}
	local currdst, currsrc
	for line in io.lines(filename) do
		local seedlist, dst, src, range
		if #line == 0 then
			-- empty line, start new map
			currdst = nil
			currsrc = nil
			goto nextline
		end
		if #seeds == 0 then
			seedlist = line:match'^seeds: ([%d ]+)$'
			if seedlist ~= nil then
				for seed in seedlist:gmatch'%d+' do
					table.insert(seeds, tonumber(seed))
				end
				-- log.trace('seedlist=%s\n', seedlist)
				goto nextline
			end
		end
		if currsrc == nil and currdst == nil then
			currsrc, currdst = line:match'^(%l+)%-to%-(%l+) map:$'
			if currsrc ~= nil and currdst ~= nil then
				-- start new map
				mapno = mapno + 1
				maps[mapno] = {}
				goto nextline
			end
		else
			dst, src, range = line:match'^(%d+) (%d+) (%d+)$'
			if dst ~= nil and src ~= nil and range ~= nil then
				table.insert(maps[mapno], {tonumber(dst), tonumber(src), tonumber(range)})
				goto nextline
			end
		end
		log.error('unhandled line %s\n', line)
::nextline::
	end
end

local function runMap(map, num)
	for _, m in ipairs(map) do
		local dest, src, range = m[1], m[2], m[3]
		if num >= src and num <= src + range then
			num = num + (dest - src)
			break
		end
	end
	return num
end

local function runAllMaps(num)
	for i = 1, #maps do
		num = runMap(maps[i], num)
	end
	return num
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)

	loadSeedsAndMaps(filename)
	-- log.list(seeds)

	local result = 1/0 -- inf

	for _, seed in ipairs(seeds) do
		local num = runAllMaps(seed)
		result = math.min(result, num)
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

	loadSeedsAndMaps(filename)

	-- input has 10 seed ranges, none of which overlap (shame)
	-- 22,746,097,099 total seeds to check

	-- A quicker brute force is to loop from 0 to infinity and apply the mapping in reverse,
	-- iterating until the mapped value falls in a seed range.
	-- The unmapped value is your answer.
	-- That runs in about 15 seconds for me.

	-- collapse maps into a single seed-to-location map?

	---@type number
	local result = 1/0	-- inf

	-- seeds = {seeds[9], seeds[10]}	-- knobble for incremental testing

	for i = 1, #seeds, 2 do
		local seed1 = seeds[i]
		local last = seed1 + seeds[i+1] - 1
		for seed = seed1, last do
			local num = runAllMaps(seed)
			result = math.min(result, num)
		end
	end

	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('05-test.txt', 35))
log.report('part one      %d\n', partOne('05-input.txt', 289863851))
log.report('part two test %d\n', partTwo('05-test.txt', 46))
-- log.report('part two      %d\n', partTwo('05-input.txt', 60568880))
-- part 2, 1 2 num range =  554772016 (but at least Mark II finished)
-- part 2, 3 4 num range =  289863851
-- part 2, 5 6 num range = 2036266413
-- part 2, 7 8 num range =   60568880 (it's a bingo!)
-- part 2, 9 10 num range =  90229603