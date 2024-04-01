-- https://adventofcode.com/2019/day/19

local log = require 'log'

---@param input string?
---@return integer[]
local function load(input)

	if input == nil then
		local f = assert(io.open('19-input.txt', 'r'))
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
	local ip, relbase = 1, 0
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
			store(getAddress(modeA, parameterA), table.remove(inputs, 1))
			goto continue
		end

		if opcode == '04' then
			-- print('OUTPUT', valueA)
			return valueA
			-- io.write(string.char(valueA))
			-- return data
			-- goto continue
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
	local SIZE = 50

	local result = 0

	for y=1,SIZE do
		for x=1,SIZE do
			local code = intcode(shallowcopy(data), {x-1, y-1})
			if code == 1 then
				result = result + 1
			end
		end
	end

	return result
end

log.report('Part One %d\n', partOne())	-- 169

local function partTwo()
	local BOXSIZE = 100
	local SIZEX = BOXSIZE*10
	local SIZEY = BOXSIZE*15
	local grid = {}

	-- this takes virtually all the time ...
	for y=1,SIZEY do
		local row = {}
		for x=1,SIZEX do
			local code = intcode(shallowcopy(data), {x-1, y-1})
			if code == 0 then
				row[#row+1] = '.'
			elseif code == 1 then
				row[#row+1] = '#'
			else
				row[#row+1] = '?'
			end
		end
		grid[#grid+1] = row
	end

--[[
	for y=1,SIZEY do
		for x=1,SIZEX do
			local ch = grid[y][x]
			io.write(ch)
		end
		io.write('\n')
	end
]]
	-- ... whereas this takes almost no time
	for x=BOXSIZE,SIZEX-BOXSIZE do
		for y=BOXSIZE,SIZEY-BOXSIZE do
			local ch = grid[y][x]	-- top left
			if ch == '#' then
				ch = grid[y][x+BOXSIZE-1]	-- top right
				if ch == '#' then
					ch = grid[y+BOXSIZE-1][x] -- bottom left
					if ch == '#' then
						ch = grid[y+BOXSIZE-1][x+BOXSIZE-1] -- bottom right
						if ch == '#' then
							log.report('x=%d, y=%d\n', x-1, y-1)
							return ((x-1) * 10000) + (y-1)
						end
					end
				end
			end
		end
	end

	return -1
end

log.report('Part Two %d\n', partTwo()) -- 700,1134

--[[
$ time luajit 19.lua
Part One 169
x=700, y=1134
Part Two 7001134

real	0m36.373s
user	0m36.082s
sys	0m0.288s
]]