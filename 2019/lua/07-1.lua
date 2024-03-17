-- https://adventofcode.com/2019/day/7

local log = require 'log'

local function load()
	local f = assert(io.open('07-input.txt', 'r'))
	local input = f:read('*all')
	f:close()

	local prog = {}
	for num in input:gmatch('[0-9]+') do
		prog[#prog+1] = tonumber(num)
	end

	return prog
end

local function clone(master)
	local program = {}
	for i, n in ipairs(master) do
		program[i] = n
	end
	return program
end

---incode interpreter with input and output via functions
---@param program integer[]
---@param ip integer instruction pointer
---@param inputfn function
---@param outputfn function
---@return integer, integer ip instruction pointer, output
local function intcode(program, ip, inputfn, outputfn)

	local function readOpcode(val)
		local opcode, modes = 0, {0, 0, 0}
		local digit = 1
		for v in string.gmatch(tostring(val):reverse(), '.') do
			if digit == 1 then
				opcode = tonumber(v)
			elseif digit == 2 then
				opcode = opcode + (tonumber(v) * 10)
			elseif digit == 3 then
				modes[1] = tonumber(v)
			elseif digit == 4 then
				modes[2] = tonumber(v)
			elseif digit == 5 then
				modes[3] = tonumber(v)
			end
			digit = digit + 1
		end
		return opcode, modes
	end

	local function readValue(pos, mode)
		if mode == 0 then
			-- position mode
			return program[pos + 1]
		elseif mode == 1 then
			-- immediate mode
			return pos
		else
			log.error('unknown mode %d\n', mode)
		end
	end

	while true do
		local opcode, modes = readOpcode(program[ip])
		if opcode == 1 then
			-- add
			local x = readValue(program[ip + 1], modes[1])
			local y = readValue(program[ip + 2], modes[2])
			local z = program[ip + 3] + 1
			program[z] = x + y
			ip = ip + 4
		elseif opcode == 2 then
			-- multiply
			local x = readValue(program[ip + 1], modes[1])
			local y = readValue(program[ip + 2], modes[2])
			local z = program[ip + 3] + 1
			program[z] = x * y
			ip = ip + 4
		elseif opcode == 3 then
			-- input
			program[program[ip + 1] + 1] = inputfn()
			ip = ip + 2
		elseif opcode == 4 then
			-- output
			-- return ip, program[program[ip + 1] + 1]
			outputfn(program[program[ip + 1] + 1])
			ip = ip + 2
		elseif opcode == 5 then
			-- jump-if-true
			local state = readValue(program[ip + 1], modes[1])
			if state > 0 then
				ip = readValue(program[ip + 2], modes[2]) + 1
			else
				ip = ip + 3
			end
		elseif opcode == 6 then
			-- jump-if-false
			local state = readValue(program[ip + 1], modes[1])
			if state == 0 then
				ip = readValue(program[ip + 2], modes[2]) + 1
			else
				ip = ip + 3
			end
		elseif opcode == 7 then
			-- less-than
			local var_one = readValue(program[ip + 1], modes[1])
			local var_two = readValue(program[ip + 2], modes[2])
			if var_one < var_two then
				program[program[ip + 3] + 1] = 1
			else
				program[program[ip + 3] + 1] = 0
			end
			ip = ip + 4
		elseif opcode == 8 then
			-- equals
			local var_one = readValue(program[ip + 1], modes[1])
			local var_two = readValue(program[ip + 2], modes[2])
			if var_one == var_two then
				program[program[ip + 3] + 1] = 1
			else
				program[program[ip + 3] + 1] = 0
			end
			ip = ip + 4
		elseif opcode == 99 then
			-- break
			return -1, -1
		else
			log.error('unknown opcode %d\n', opcode)
			break
		end
	end

	return ip, -1	-- we never come here
end

---incode interpreter with input and output via functions
---@param program integer[]
---@param ip integer instruction pointer
---@param inputs integer[]
---@return integer, integer ip instruction pointer, output
local function intcode2(program, ip, inputs)

	local function readOpcode(val)
		local opcode, modes = 0, {0, 0, 0}
		local digit = 1
		for v in string.gmatch(tostring(val):reverse(), '.') do
			if digit == 1 then
				opcode = tonumber(v)
			elseif digit == 2 then
				opcode = opcode + (tonumber(v) * 10)
			elseif digit == 3 then
				modes[1] = tonumber(v)
			elseif digit == 4 then
				modes[2] = tonumber(v)
			elseif digit == 5 then
				modes[3] = tonumber(v)
			end
			digit = digit + 1
		end
		return opcode, modes
	end

	local function readValue(pos, mode)
		if mode == 0 then
			-- position mode
			return program[pos + 1]
		elseif mode == 1 then
			-- immediate mode
			return pos
		else
			log.error('unknown mode %d\n', mode)
		end
	end

	local ninput = 1
	while true do
		local opcode, modes = readOpcode(program[ip])
		if opcode == 1 then
			-- add
			local x = readValue(program[ip + 1], modes[1])
			local y = readValue(program[ip + 2], modes[2])
			local z = program[ip + 3] + 1
			program[z] = x + y
			ip = ip + 4
		elseif opcode == 2 then
			-- multiply
			local x = readValue(program[ip + 1], modes[1])
			local y = readValue(program[ip + 2], modes[2])
			local z = program[ip + 3] + 1
			program[z] = x * y
			ip = ip + 4
		elseif opcode == 3 then
			-- input
			program[program[ip + 1] + 1] = inputs[ninput]
			ninput = ninput + 1
			ip = ip + 2
		elseif opcode == 4 then
			-- output
			return ip + 2, program[program[ip + 1] + 1]
			-- ip = ip + 2
		elseif opcode == 5 then
			-- jump-if-true
			local state = readValue(program[ip + 1], modes[1])
			if state > 0 then
				ip = readValue(program[ip + 2], modes[2]) + 1
			else
				ip = ip + 3
			end
		elseif opcode == 6 then
			-- jump-if-false
			local state = readValue(program[ip + 1], modes[1])
			if state == 0 then
				ip = readValue(program[ip + 2], modes[2]) + 1
			else
				ip = ip + 3
			end
		elseif opcode == 7 then
			-- less-than
			local var_one = readValue(program[ip + 1], modes[1])
			local var_two = readValue(program[ip + 2], modes[2])
			if var_one < var_two then
				program[program[ip + 3] + 1] = 1
			else
				program[program[ip + 3] + 1] = 0
			end
			ip = ip + 4
		elseif opcode == 8 then
			-- equals
			local var_one = readValue(program[ip + 1], modes[1])
			local var_two = readValue(program[ip + 2], modes[2])
			if var_one == var_two then
				program[program[ip + 3] + 1] = 1
			else
				program[program[ip + 3] + 1] = 0
			end
			ip = ip + 4
		elseif opcode == 99 then
			-- break
			return -1, -1
		else
			log.error('unknown opcode %d\n', opcode)
			break
		end
	end

	return ip, -1	-- we never come here
end

---generate all permutations using Heap's algorithm
---@param arr any[]
---@return function iterator
local function permutations(arr)
	local function swap(a, b)
		arr[a], arr[b] = arr[b], arr[a]
	end

	local function heap_permute(n)
		if n == 1 then
			coroutine.yield(arr)
		else
			for i = 1, n do
				heap_permute(n - 1)
				if n % 2 == 0 then
					swap(i, n)
				else
					swap(1, n)
				end
			end
		end
	end

	return coroutine.wrap(function()
		heap_permute(#arr)
	end)
end

---@return integer
local function part1()
	local master = load()
	local result = 0
	for perm in permutations({0,1,2,3,4}) do
		local output = 0
		for _, phase in ipairs(perm) do
			_, output = intcode2(clone(master), 1, {phase, output})
			if output == -1 then
				log.error('unexpected intcode termination')
			end
		end
		if output > result then
			-- log.trace('%s = %d\n', table.concat(perm, ' '), output)
			-- 0 1 2 3 4 = 40708
			-- 1 0 2 3 4 = 46108
			-- 1 0 2 4 3 = 46248
			result = output
		end
	end
	return result
end

---@return integer
local function part2()
	local master = load()
	local result = 0
	for phases in permutations({5,6,7,8,9}) do
		local amplifiers = {}
		for i = 1, 5 do
			amplifiers[i] = clone(master)
		end
		local pcs = {1,1,1,1,1}
		local output = 0
		for loop = 0, 99 do
			for i, phase in ipairs(phases) do
				if loop == 0 then
					pcs[i], output = intcode2(amplifiers[i], pcs[i], {phase, output})
				else
					pcs[i], output = intcode2(amplifiers[i], pcs[i], {output})
				end
				if pcs[i] == -1 then
					-- log.report('%d HALT\n', i)
					goto exitloop
				end
				if output == -1 then
					log.report('-ve output\n')
				end
			end
		end
::exitloop::
		if output > result then
			result = output
		end
	end
	return result
end

log.report('part one %d\n', part1())
-- log.report('part two %d\n', part2())
