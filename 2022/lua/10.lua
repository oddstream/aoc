-- https://adventofcode.com/2022/day/10 Cathode-Ray Tube

local log = require 'log'


---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0
	local nums = {}
	for line in io.lines(filename) do
		local num = line:match'addx ([%-%d]+)'
		if num then
			nums[#nums+1] = 0
			nums[#nums+1] = tonumber(num)
		else
			nums[#nums+1] = 0
		end
	end

	local register = 1
	local n = 20
	for cycle, num in ipairs(nums) do
		if cycle == n then
			local strength = (cycle * register)
			-- print(cycle, x, strength)
			result = result + strength
			n = n + 40
		end
		register = register + num
	end

	if expected and result ~= expected then
		log.error('expected %d\n', expected)
	end
	return result
end

---@param filename string
local function partTwo(filename)

	local function key(y,x)
		return tostring(y) .. ',' .. tostring(x)
	end

	local nums = {}
	for line in io.lines(filename) do
		local num = line:match'addx ([%-%d]+)'
		if num then
			nums[#nums+1] = 0
			nums[#nums+1] = tonumber(num)
		else
			nums[#nums+1] = 0
		end
	end

	local register = 1
	local onpix = {} -- record the locations of the on pixels in a map
	for cycle, num in ipairs(nums) do
		if math.abs(((cycle-1)%40)-register) < 2 then
			onpix[key(cycle//40,cycle%40)] = true
		end
		register = register + num
	end

	for y = 0, 5 do
		for x = 0, 40 do
			if onpix[key(y,x)] then
				io.write('#')
			else
				io.write('.')
			end
		end
		io.write('\n')
	end

end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('10-input-test.txt', 13140))
log.report('part one      %d\n', partOne('10-input.txt', 13220))
partTwo('10-input.txt') -- RUAKHBEK

--[[
$ time lua54 10.lua
Lua 5.4
part one test 13140
part one      13220
.###..#..#..##..#..#.#..#.###..####.#..#.
.#..#.#..#.#..#.#.#..#..#.#..#.#....#.#..
.#..#.#..#.#..#.##...####.###..###..##...
.###..#..#.####.#.#..#..#.#..#.#....#.#..
.#.#..#..#.#..#.#.#..#..#.#..#.#....#.#..
.#..#..##..#..#.#..#.#..#.###..####.#..#.

real	0m0.005s
user	0m0.001s
sys	0m0.005s
]]