-- https://adventofcode.com/2019/day/2

local log = require 'log'

local function load()
	local f = assert(io.open('02-input.txt', "r"))
	local input = f:read("*all")
	f:close()

	local program = {}
	for num in input:gmatch('(%d+)') do
		program[#program+1] = tonumber(num)
	end

	return program
end

local function run(prog, noun, verb)
	prog[2] = noun
	prog[3] = verb
	local pc = 1
	while true do
		local opcode = prog[pc]
		if opcode == 1 then
			local a, b, c = prog[pc+1], prog[pc+2], prog[pc+3]
			prog[c+1] = prog[a+1] + prog[b+1]
		elseif opcode == 2 then
			local a, b, c = prog[pc+1], prog[pc+2], prog[pc+3]
			prog[c+1] = prog[a+1] * prog[b+1]
		elseif opcode == 99 then
			break
		else
			log.error('unexpected opcode %d at position %d\n', prog[pc], pc)
			break
		end
		pc = pc + 4
	end
	return prog[1]
end

local function part1()
	return run(load(), 12, 2)
end

local function part2()
	for noun = 0, 99 do
		for verb = 0, 99 do
			if run(load(), noun, verb) == 19690720 then
				return 100 * noun + verb
			end
		end
	end
	return -1
end

log.report('part one %d\n', part1())
log.report('part two %d\n', part2())

--[[
$ time luajit 02.lua
part one 3101878
part two 8444
real	0m0.153s
user	0m0.124s
sys	0m0.028s

$ time lua54 02.lua
part one 3101878
part two 8444
real	0m0.250s
user	0m0.177s
sys	0m0.064s

$ time lua51 02.lua
part one 3101878
part two 8444
real	0m0.401s
user	0m0.377s
sys	0m0.024s
]]