-- https://adventofcode.com/2019/day/25

-- ProggyVector

local log = require 'log'

---@param input string?
---@return integer[]
local function load(input)

	if input == nil then
		local f = assert(io.open('25-input.txt', 'r'))
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

local shortcuts = {
	['n'] = 'north',
	['e'] = 'east',
	['s'] = 'south',
	['w'] = 'west',
	['i'] = 'inv',
	['tca'] = 'take cake',
	['dca'] = 'drop cake',
	['tco'] = 'take coin',
	['dco'] = 'drop coin',
	['tdw'] = 'take dehydrated water',
	['ddw'] = 'drop dehydrated water',
	['tfc'] = 'take fuel cell',
	['dfc'] = 'drop fuel cell',
	['tma'] = 'take manifold',
	['dma'] = 'drop manifold',
	['tmu'] = 'take mutex',
	['dmu'] = 'drop mutex',
	['tpn'] = 'take prime number',
	['dpn'] = 'drop prime number',
}

local function intcode(data, preamble)
	local POSITION, IMMEDIATE, RELATIVE = '0', '1', '2'
	local ip, relbase = 1, 0
	local inputq = {}

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
			if #preamble > 0 then
				store(getAddress(modeA, parameterA), table.remove(preamble, 1))
				goto continue
			end

			-- io.write('INPUT')
			if #inputq == 0 then
				log.info('n e s w inv take/drop <item> ')
				local str = io.read()
				for k, v in pairs(shortcuts) do
					if str == k then
						str = v
						break
					end
				end
				for ch in str:gmatch'.' do
					table.insert(inputq, string.byte(ch))
				end
				table.insert(inputq, 10)
			end
			store(getAddress(modeA, parameterA), table.remove(inputq, 1))
			goto continue
		end

		if opcode == '04' then
			if valueA > 255 then
				log.report('Part One %d\n', valueA)
				return
			else
				io.write(string.char(valueA))
			end
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

local preambleString = [[
south
take fuel cell
south
take manifold
north
north
west
take mutex
south
south
take coin
west
take dehydrated water
south
take prime number
north
east
north
east
take cake
north
west
south
inv
]]
local preambleList = {}
for i=1,#preambleString do
	table.insert(preambleList, string.byte(string.sub(preambleString, i,i)))
end

intcode(shallowcopy(load()), preambleList)	-- 278664 cake coin mutex fuel cell
