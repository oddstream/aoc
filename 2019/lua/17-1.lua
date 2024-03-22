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

local function robot(data)
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
			-- print('HALT')
			break
		end

		local modeA = header:sub(3,3)
		local modeB = header:sub(2,2)
		local modeC = header:sub(1,1)

		-- opcodes with 1 parameter

		local parameterA = read()
		local valueA = getValue(modeA, parameterA)

		if opcode == '03' then
			store(getAddress(modeA, parameterA), 0)
			print('INPUT')
			goto continue
		end

		if opcode == '04' then
			-- io.write(string.char(valueA))
			coroutine.yield(valueA)
			-- if valueA == -1 then print('OUTPUT', valueA) end
			-- return bot
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

local function boterator(mem)
	local co = coroutine.create(function() robot(mem) end)
	return function()
		local success, i = coroutine.resume(co)
		if success then return i end
	end
end

local grid = {}
local row = {}
for ch in boterator(load()) do
	if ch == 10 then
		grid[#grid+1] = row
		row = {}
	else
		row[#row+1] = string.char(ch)
	end
	-- io.write(string.char(ch))
end

-- for y=1, #grid do
-- 	for x=1,#grid[y] do
-- 		io.write(grid[y][x])
-- 	end
-- 	io.write('\n')
-- end

local result = 0
for y=2, #grid-1 do
	for x = 2, #grid[y]-1 do
		if grid[y][x] == '#' then
			local neighbours = 0
			if grid[y-1][x] == '#' then neighbours = neighbours + 1 end
			if grid[y+1][x] == '#' then neighbours = neighbours + 1 end
			if grid[y][x-1] == '#' then neighbours = neighbours + 1 end
			if grid[y][x+2] == '#' then neighbours = neighbours + 1 end
			if neighbours == 4 then
				-- log.trace('intersection at %d,%d\n', y, x)
				result = result + (y-1) * (x-1)
			end
		end
	end
end

log.report('Part One %d\n', result)	-- 3936

-- you can visit the entire scaffold using 'go forward until you can't then make the only turn'

-- robot(load())
