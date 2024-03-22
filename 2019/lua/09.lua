-- https://adventofcode.com/2019/day/9

local log = require 'log'

local POSITION, IMMEDIATE, RELATIVE = '0', '1', '2'

---@param input string?
---@return string[]
local function load(input)

	if input == nil then
		local f = assert(io.open('09-input.txt', 'r'))
		input = f:read('*all')
		f:close()
	end

	local prog = {}
	-- some numbers are -ve
	for num in input:gmatch('[0-9-]+') do
		prog[#prog+1] = tonumber(num)
	end

	return prog
end

---@param master string[]
---@return string[]
local function clone(master)
	local program = {}
	for i, n in ipairs(master) do
		program[i] = n
	end
	-- The computer's available memory should be much larger than the initial program.
	-- Memory beyond the initial program starts with the value 0
	-- and can be read or written like any other memory.
	-- BUT Lua doesn't seem to care about any extra memory
	return program
end

local function run(bot)

	---return the next data item, advance instruction pointer
	---@return integer
	local function read()
		local v = bot.data[bot.ip]	-- the instruction pointer is already 1-based
		bot.ip = bot.ip + 1
		return v
	end

	local function store(loc, val)
		bot.data[loc+1] = val	-- +1 because Lua arrays are 1-based
	end

	---@comment tied in knots over the 203 problem, and this makes no sense
	---@param mode string
	---@param parameter integer
	---@return integer
	local function getAddress(mode, parameter)
		if mode == POSITION then
			return parameter
		elseif mode == RELATIVE then
			return bot.relbase + parameter
		else
			log.error('unexpected mode in getAddress %d\n', mode)
			return 0
		end
	end

	---@comment tied in knots over the 203 problem, and this makes no sense
	---@param mode string
	---@param parameter integer
	---@return integer
	local function getValue(mode, parameter)
		if mode == IMMEDIATE then
			return parameter
		end
		local address = getAddress(mode, parameter)
		return bot.data[address+1]
	end

	while true do
		local header = string.format('%05d', read())

		local opcode = header:sub(4,5)

		if opcode == '99' then
			bot.halted = true
			print('HALT')
			break
		end

		local modeA = header:sub(3,3)
		local modeB = header:sub(2,2)
		local modeC = header:sub(1,1)

		-- opcodes with 1 parameter

		local parameterA = read()
		local valueA = getValue(modeA, parameterA)

		if opcode == '03' then
			store(getAddress(modeA, parameterA), bot.input)
			print('INPUT', bot.input)
			goto continue
		end

		if opcode == '04' then
			bot.output = valueA
			print('OUTPUT', valueA)
			-- break
			goto continue
		end

		if opcode == '09' then
			bot.relbase = bot.relbase + valueA
			goto continue
		end

		-- opcodes with 2 parameters

		local parameterB = read()
		local valueB = getValue(modeB, parameterB)

		if opcode == '05' then
			if valueA ~= 0 then
				bot.ip = valueB + 1
			end
			goto continue
		end
		if opcode == '06' then
			if valueA == 0 then
				bot.ip = valueB + 1
			end
			goto continue
		end

		-- opcodes with 3 parameters

		local parameterC = read()
		local addressC = getAddress(modeC, parameterC)

		if opcode == '01' then
			store(addressC, valueA + valueB)
		elseif opcode == '02' then
			store(addressC, valueA * valueB)
		elseif opcode == '07' then
			if valueA < valueB then
				store(addressC, 1)
			else
				store(addressC, 0)
			end
		elseif opcode == '08' then
			if valueA == valueB then
				store(addressC, 1)
			else
				store(addressC, 0)
			end
		end

::continue::
	end
end

local function test(input)
	local master = load(input)
	local bot = {
		data = clone(master, 100),
		ip = 1,
		relbase = 0,
		input = 0,
		output = 0,
	}
	run(bot)
	return bot.output
end

local function part1()
	local master = load()
	local bot = {
		data = clone(master),
		ip = 1,
		relbase = 0,
		input = 1,
		output = 0,
	}
	run(bot)
	return bot.output
end

local function part2()
	local master = load()
	local bot = {
		data = clone(master),
		ip = 1,
		relbase = 0,
		input = 2,
		output = 0,
	}
	run(bot)
	return bot.output
end

-- https://www.reddit.com/r/adventofcode/comments/e8aw9j/2019_day_9_part_1_how_to_fix_203_error/
-- with thanks to https://github.com/JoanaBLate/advent-of-code-js/blob/main/2019/day09/solve1.js#L58

-- log.trace('test 1 %d\n', test('109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99'))
-- log.trace('test 2 %d\n', test('1102,34915192,34915192,7,4,7,99,0'))
-- log.trace('test 3 %d\n', test('104,1125899906842624,99'))
log.report('part 1 %d\n', part1())
log.report('part 2 %d\n', part2())

--[[
$ time luajit 09.lua
INPUT	1
OUTPUT	3742852857
HALT
part 1 3742852857
INPUT	2
OUTPUT	73439
HALT
part 2 73439

real	0m0.022s
user	0m0.022s
sys	0m0.000s
]]