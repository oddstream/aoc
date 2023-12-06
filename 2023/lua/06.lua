-- https://adventofcode.com/2023/day/6

local log = require 'log'

---@param filename string
---@return table, table
local function readInput1(filename)
	local times = {}
	local distances = {}
	for line in io.lines(filename) do
		if line:match'^Time: ' then
			for t in line:gmatch'(%d+)' do
				table.insert(times, tonumber(t))
			end
		elseif line:match'^Distance: ' then
			for d in line:gmatch'(%d+)' do
				table.insert(distances, tonumber(d))
			end
		end
	end
	return times, distances
end

---@param filename string
---@return number?, number?
local function readInput2(filename)
	local time = ''
	local distance = ''
	for line in io.lines(filename) do
		if line:match'^Time: ' then
			for t in line:gmatch'(%d+)' do
				time = time .. t
			end
		elseif line:match'^Distance: ' then
			for d in line:gmatch'(%d+)' do
				distance = distance .. d
			end
		end
	end
	return tonumber(time), tonumber(distance)
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)

	local times, distances = readInput1(filename)

	local result = 1
	for race = 1, #times do
		local wins = 0
		for hold = 1, times[race] - 1 do
			local speed = hold
			local time = times[race] - hold
			local distance = (speed * time)
			if distance > distances[race] then
				wins = wins + 1
			end
		end
		result = result * wins
		-- log.trace('race %d, ways %d\n', race, ways)
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

	local time, distance = readInput2(filename)

	local result = 0
	for hold = 1, time - 1 do
		local speed = hold
		local t = time - hold
		local d = (speed * t)
		if d > distance then
			result = result + 1
		end
	end

	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('06-test.txt', 288))
log.report('part one      %d\n', partOne('06-input.txt', 2269432))
log.report('part two test %d\n', partTwo('06-test.txt', 71503))
log.report('part two      %d\n', partTwo('06-input.txt', 35865985))

--[[
$ time luajit 06.lua
Lua 5.1
part one test 288
part one      2269432
part two test 71503
part two      35865985

real    0m0.250s
user    0m0.240s
sys     0m0.002s
]]