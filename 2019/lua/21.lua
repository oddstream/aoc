-- https://adventofcode.com/2019/day/21

-- ProggyVector

local log = require 'log'

---@param input string?
---@return integer[]
local function load(input)

	if input == nil then
		local f = assert(io.open('21-input.txt', 'r'))
		input = f:read('*all')
		f:close()
	end

	local prog = {}
	-- some numbers are -ve
	for num in input:gmatch('[0-9-]+') do
		prog[#prog+1] = tonumber(num)
	end
	for _=1,50 do
		prog[#prog+1] = 0
	end

	return prog
end

local function shallowcopy(orig)
	local orig_type = type(orig)
	local copy
	if orig_type == 'table' then
		copy = {}
		for orig_key, orig_value in pairs(orig) do
			copy[orig_key] = orig_value
		end
	else -- number, string, boolean, etc
		copy = orig
	end
	return copy
end

local function intcode(data, inputs)
	local POSITION, IMMEDIATE, RELATIVE = '0', '1', '2'
	local ip, relbase, ninput = 1, 0, 1
	---return the next data item, advance instruction pointer
	---@return integer
	local function read()
		local v = data[ip]	-- the instruction pointer is already 1-based
		ip = ip + 1
		return v
	end

	local function store(loc, val)
		data[loc+1] = val	-- +1 because Lua arrays are 1-based
	end

	---@comment tied in knots over the 203 problem, and this makes no sense
	---@param mode string
	---@param parameter integer
	---@return integer
	local function getAddress(mode, parameter)
		if mode == POSITION then
			return parameter
		elseif mode == RELATIVE then
			return relbase + parameter
		else
			log.error('unexpected mode in getAddress %s\n', mode)
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
		return data[address+1]
	end

	while true do
		local header = string.format('%05d', read())

		local opcode = header:sub(4,5)

		if opcode == '99' then
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
			-- print('INPUT')
			io.write(inputs:sub(ninput,ninput))
			store(getAddress(modeA, parameterA), string.byte(inputs:sub(ninput,ninput)))
			ninput = ninput + 1
			goto continue
		end

		if opcode == '04' then
			-- print('OUTPUT', valueA, string.char(valueA))
			if valueA > 256 then
				-- log.report('damage %d\n', valueA)
				return valueA
			else
				io.write(string.char(valueA))
			end
			-- return data
			goto continue
		end

		if opcode == '09' then
			relbase = relbase + valueA
			goto continue
		end

		-- opcodes with 2 parameters

		local parameterB = read()
		local valueB = getValue(modeB, parameterB)

		if opcode == '05' then
			if valueA ~= 0 then
				ip = valueB + 1
			end
			goto continue
		end
		if opcode == '06' then
			if valueA == 0 then
				ip = valueB + 1
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
	return -1
end

local data = load()

local function partOne()
	local springscript = [[
NOT C J
AND D J
NOT A T
OR T J
WALK

]]
	return intcode(shallowcopy(data), springscript)
end

log.report('Part One %d\n', partOne())	-- 19352493

local function partTwo()
	local springscript = [[
NOT C J
AND D J
AND H J
NOT B T
AND D T
OR T J
NOT A T
OR T J
RUN

]]
	return intcode(shallowcopy(data), springscript)
end

log.report('Part Two %d\n', partTwo())	-- 1141896219

--[[
$ time luajit 21.lua
Input instructions:
NOT C J
AND D J
NOT A T
OR T J
WALK

Walking...

Part One 19352493
Input instructions:
NOT C J
AND D J
AND H J
NOT B T
AND D T
OR T J
NOT A T
OR T J
RUN

Running...

Part Two 1141896219

real	0m0.034s
user	0m0.034s
sys	0m0.000s
]]