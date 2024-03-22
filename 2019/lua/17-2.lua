-- https://adventofcode.com/2019/day/17

local log = require 'log'

local POSITION, IMMEDIATE, RELATIVE = '0', '1', '2'

---@param input string?
---@return integer[]
local function load(input)

	if input == nil then
		local f = assert(io.open('17-input.txt', 'r'))
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

local function robot(data, program)
	local ip, relbase, ii = 1, 0, 1
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
			-- io.write('INPUT: ')
			-- local str = io.read('*l')
			-- if str == '' then str = '\n' else str = string.upper(str) end
			-- store(getAddress(modeA, parameterA), string.byte(str:sub(1,1)))
			store(getAddress(modeA, parameterA), string.byte(program:sub(ii,ii)))
			ii = ii + 1
			goto continue
		end

		if opcode == '04' then
			-- print('OUTPUT')
			if valueA ~= nil then
				if valueA > 255 then
					log.report('Part Two %d\n', valueA)	-- 785733
				elseif valueA ~= 0 then
					io.write(string.char(valueA))
				end
			end
			-- if valueA == -1 then print('OUTPUT', valueA) end
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
end

--[[
path plotted by hand by referring to 17-grid.txt

R,4,L,12,				Ln 15 Col 45
L,8,R,4,				Ln 11 Col 37
L,8,					Ln 11 Col 29
R,10,R,10,R,6,			Ln  7 Col 39
R,4,L,12				Ln 19 Col 35
L,8,R,4,R,4,			Ln 23 Col 39
R,10,L,12				Ln 13 Col 27
R,4,L,12,				Ln  9 Col 15
L,8,R,4,L,8,			Ln 25 Col 11
R,10,R,10,				Ln 15 Col  1
R,6,R,4,L,12			Ln 19 Col 19
L,8,R,4,R,4,R,10		Ln 15 Col 13
L,12,L,8,R,10			Ln 37 Col 21
R,10,R,6,R,4			Ln 31 Col 15
R,10,L,12				Ln 41 Col 27

R,4,L,12,L,8,R,4,L,8,R,10,R,10,R,6,R,4,L,12,L,8,R,4,R,4,R,10,L,12,R,4,L,12,L,8,R,4,L,8,R,10,R,10,R,6,R,4,L,12,L,8,R,4,R,4,R,10,L,12,L,8,R,10,R,10,R,6,R,4,R,10,L,12

A,B,A,C,A,B,A,C,B,C

A = R,4,L,12,L,8,R,4
B = L,8,R,10,R,10,R,6
C = R,4,R,10,L,12
]]

local mem = load()
mem[1] = 2
local prog = {
	'A,B,A,C,A,B,A,C,B,C',
	'R,4,L,12,L,8,R,4',
	'L,8,R,10,R,10,R,6',
	'R,4,R,10,L,12',
	'y',
	'',
}
robot(mem, table.concat(prog, '\n'))
