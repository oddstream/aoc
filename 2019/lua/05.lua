-- https://adventofcode.com/2019/day/5

local log = require 'log'

local function load()
	local f = assert(io.open('05-input.txt', 'r'))
	local input = f:read('*all')
	f:close()

	local prog = {}
	-- input contains -ve numbers
	for num in input:gmatch('[0-9-]+') do
		prog[#prog+1] = tonumber(num)
	end

	return prog
end

---incode interpreter with input and output via functions
---@param program integer[]
---@param ip integer instruction pointer
---@param inputfn function
---@param outputfn function
---@return integer ip instruction pointer
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
			break
		else
			log.error('unknown opcode %d\n', opcode)
			break
		end
	end

	return ip
end

local function part1output(n)
	log.report('part 1 %d\n', n)
end

local function part2output(n)
	log.report('part 2 %d\n', n)
end

intcode(load(), 1, function() return 1 end, part1output)
intcode(load(), 1, function() return 5 end, part2output)

--[[
$ time luajit 05.lua
part 1 3
part 1 0
part 1 0
part 1 0
part 1 0
part 1 0
part 1 0
part 1 0
part 1 0
part 1 13294380
part 2 11460760

real	0m0.002s
user	0m0.002s
sys	0m0.000s
]]