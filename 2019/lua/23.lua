-- https://adventofcode.com/2019/day/23

local log = require 'log'

---@param input string?
---@return integer[]
local function load(input)

	if input == nil then
		local f = assert(io.open('23-input.txt', 'r'))
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

local inputq = {}

local function intcode(data, id)
	local POSITION, IMMEDIATE, RELATIVE = '0', '1', '2'
	local ip, relbase = 1, 0
	local output3 = {}

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
			coroutine.yield()
			break
		end

		local modeA = header:sub(3,3)
		local modeB = header:sub(2,2)
		local modeC = header:sub(1,1)

		-- opcodes with 1 parameter

		local parameterA = read()
		local valueA = getValue(modeA, parameterA)

		if opcode == '03' then
			if #inputq[id+1] == 0 then
				store(getAddress(modeA, parameterA), -1)
			else
				-- print('INPUT', id, inputq[id+1][1])
				store(getAddress(modeA, parameterA), table.remove(inputq[id+1], 1))
			end
			coroutine.yield()
			goto continue
		end

		if opcode == '04' then
			table.insert(output3, valueA)
			if #output3 == 3 then
				-- print('OUTPUT', id, output3[1], output3[2], output3[3])
				coroutine.yield(output3)
				output3 = {}
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

local function partOne()

	inputq = {}
	for i = 0,49 do
		inputq[i+1] = {i}
	end

	local data = load()
	local interp = {}
	for i = 0,49 do
		interp[i+1] = coroutine.create(function() intcode(shallowcopy(data), i) end)
	end
	while true do
		for _, co in ipairs(interp) do
			local ok, val = coroutine.resume(co)
			if ok then
				if val ~= nil and #val == 3 then
					if val[1] == 255 then
						return val[3]
					else
						table.insert(inputq[val[1]+1], val[2])
						table.insert(inputq[val[1]+1], val[3])
					end
				end
			else
				print('resume error', val)
				return -1
			end
		end
	end
end

local function partTwo()

	local function idle()
		-- "If all computers have empty incoming packet queues
		-- and are continuously trying to receive packets without sending packets,
		-- the network is considered idle."
		for i=0,49 do
			if #inputq[i+1] > 0 then
				return false
			end
		end
		return true
	end

	inputq = {}
	for i = 0,49 do
		inputq[i+1] = {i}
	end
	local data = load()
	local interp = {}
	for i = 0,49 do
		interp[i+1] = coroutine.create(function() intcode(shallowcopy(data), i) end)
	end
	local iterations = 0
	local natx, naty, prevnaty
	while true do
		for _, co in ipairs(interp) do
			local ok, val = coroutine.resume(co)
			if ok then
				if val ~= nil and #val == 3 then
					if val[1] == 255 then
						-- "If a packet would be sent to address 255, the NAT receives it instead"
						-- print('send to NAT', val[1], val[2], val[3])
						natx, naty = val[2], val[3]
					else
						table.insert(inputq[val[1]+1], val[2])
						table.insert(inputq[val[1]+1], val[3])
					end
				end
			else
				print('resume error', val)
				return -1
			end
		end
		iterations = iterations + 1
		if iterations % 10 == 0 then
			if idle() and natx ~= nil and naty ~= nil then
				if naty == prevnaty then
					return naty
				end
				-- "Once the network is idle, the NAT sends only the last packet it received to address 0"
				table.insert(inputq[1], natx)
				table.insert(inputq[1], naty)
				prevnaty = naty
			end
		end
	end
end

log.report('Part One %d\n', partOne())	-- 24922
log.report('Part Two %d\n', partTwo())	-- 19478

--[[
$ which luajit
/usr/local/bin/luajit
$ time luajit 23.lua
Part One 24922
Part Two 19478

real	0m0.060s
user	0m0.060s
sys	0m0.000s
]]