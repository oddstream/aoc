-- https://adventofcode.com/2019/day/7

local log = require 'log'

---@return string[]
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

---@param master string[]
---@return string[]
local function clone(master)
	local program = {}
	for i, n in ipairs(master) do
		program[i] = n
	end
	return program
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

local function run(amp, input)
	while true do
		local header = string.format('%05d', amp.data[amp.ip])
		amp.ip = amp.ip + 1

		local opcode = header:sub(4,5)

		if opcode == '99' then
			amp.halted = true
			break
		end

		local modeA = header:sub(3,3)
		local modeB = header:sub(2,2)
		-- local modeC = header:sub(1,1)

		local parameterA = amp.data[amp.ip]
		amp.ip = amp.ip + 1

		local valueA
		if modeA == '0' then
			valueA = amp.data[parameterA + 1]
		elseif modeA == '1' then
			valueA = parameterA
		else
			log.error('unknown modeA %s\n', modeA)
		end

		if opcode == '03' then
			if amp.starting then
				amp.starting = false
				amp.data[parameterA + 1] = amp.phase
			else
				amp.data[parameterA + 1] = input
			end
			goto continue
		end

		if opcode == '04' then
			amp.output = valueA
			break
		end

		local parameterB = amp.data[amp.ip]
		amp.ip = amp.ip + 1

		local valueB
		if modeB == '0' then
			valueB = amp.data[parameterB + 1]
		elseif modeB == '1' then
			valueB = parameterB
		else
			log.error('unknown modeB %s\n', modeB)
		end

		if opcode == '05' then
			if valueA ~= 0 then
				amp.ip = valueB + 1
			end
			goto continue
		end
		if opcode == '06' then
			if valueA == 0 then
				amp.ip = valueB + 1
			end
			goto continue
		end

		local parameterC = amp.data[amp.ip]
		amp.ip = amp.ip + 1

		if opcode == '01' then
			amp.data[parameterC + 1] = valueA + valueB
		elseif opcode == '02' then
			amp.data[parameterC + 1] = valueA * valueB
		elseif opcode == '07' then
			if valueA < valueB then
				amp.data[parameterC + 1] = 1
			else
				amp.data[parameterC + 1] = 0
			end
		elseif opcode == '08' then
			if valueA == valueB then
				amp.data[parameterC + 1] = 1
			else
				amp.data[parameterC + 1] = 0
			end
		end

::continue::
	end
end

---@param program integer[]
---@param phases integer[]
---@return integer
local function tryPhases(program, phases)

	local amplifiers = {}
	for i = 1, 5 do
		amplifiers[i] = {
			starting = true,
			halted = false,
			phase = phases[i],
			output = 0,
			data = clone(program),
			ip = 1,
		}
	end

	while true do
		local prev = 5
		for i = 1, 5 do
			run(amplifiers[i], amplifiers[prev].output)
			prev = i
		end
		-- if amplifiers[5].output == 54163586 then
		-- 	print('bingo', table.concat(phases, ' '), amplifiers[5].halted)
		-- end
		if amplifiers[5].halted then
			return amplifiers[5].output
		end
	end
end

local function part1()
	local master = load()
	local biggest = 0
	for phases in permutations({0,1,2,3,4}) do
		local output = 0
		for i = 1, 5 do
			local amp = {
				starting = true,
				halted = false,
				phase = phases[i],
				output = output,
				data = clone(master),
				ip = 1,
			}
			run(amp, output)
			output =  amp.output
			if output > biggest then
				biggest = output
			end
		end
	end
	return biggest
end

local function part2()
	local master = load()
	local biggest = 0
	for phases in permutations({5,6,7,8,9}) do
		local result = tryPhases(master, phases)
		if result > biggest then
			biggest = result
		end
	end
	return biggest
end

log.trace('part one %d\n', part1())
log.trace('part two %d\n', part2())

--[[
$ time luajit 07-2.lua
part one 46248
part two 54163586

real	0m0.006s
user	0m0.005s
sys	0m0.000s
]]