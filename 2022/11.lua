-- https://adventofcode.com/2022/day/11 Monkey in the Middle

local log = require 'log'

local function loadMonkeys(filename)
	local monkeys = {}
	for line in io.lines(filename) do
		if #line == 0 then
			-- ignore blank lines
		elseif line:match'Monkey %d+:' then
			table.insert(monkeys, {inspections=0})
		elseif line:match'  Starting items: .+' then
			local lst = line:match'  Starting items: ([%d ,]+)'
			local nums = {}
			for num in lst:gmatch'(%d+)' do
				nums[#nums+1] = tonumber(num)
			end
			monkeys[#monkeys].items = nums
		elseif line:match'  Operation: .+' then
			local op, rhs = line:match'  Operation: new = old ([%*%+]) (.+)'
			assert(op=='+' or op=='*')
			monkeys[#monkeys].op = op
			if rhs ~= 'old' then
				monkeys[#monkeys].op_num = tonumber(rhs) -- will be nil if 'old'
			end
		elseif line:match'  Test: .+' then
			local num = line:match'  Test: divisible by (%d+)'
			monkeys[#monkeys].test_num = tonumber(num)
		elseif line:match'    If true:' then
			local dst = line:match'    If true: throw to monkey (%d+)'
			monkeys[#monkeys].true_dst = tonumber(dst)
		elseif line:match'    If false:' then
			local dst = line:match'    If false: throw to monkey (%d+)'
			monkeys[#monkeys].false_dst = tonumber(dst)
		else
			log.error('cannot parse line %s\n', line)
		end
	end
	return monkeys
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result

	local monkeys = loadMonkeys(filename)

	for _ = 1, 20 do
		for _, monkey in ipairs(monkeys) do
			for _, item in ipairs(monkey.items) do
				local worry
				if monkey.op == '*' then
					worry = item * (monkey.op_num or item)
				elseif monkey.op == '+' then
					worry = item + (monkey.op_num or item)
				end
				worry = worry // 3
				-- if worry / monkey.test_num == worry // monkey.test_num then
				if math.fmod(worry, monkey.test_num) == 0 then
					table.insert(monkeys[monkey.true_dst + 1].items, worry)
				else
					table.insert(monkeys[monkey.false_dst + 1].items, worry)
				end
			end
			monkey.inspections = monkey.inspections + #monkey.items
			monkey.items = {}
		end
	end

	table.sort(monkeys, function(a,b) return a.inspections > b.inspections end)

	result = monkeys[1].inspections * monkeys[2].inspections

	if expected and result ~= expected then
		log.error('expected %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result

	local monkeys = loadMonkeys(filename)

	-- Get the product of all the "divisible by" numbers,
	-- and modulo the current worry level by that product
	-- This way, we are still able to check for divisibility with any of the divisors,
	-- but the values we keep don't "explode".
	-- This works because the only operations we perform on the values are multiplication and addition,
	-- which preserve congruence.
	local product = 1
	for _, monkey in ipairs(monkeys) do
		product = product * monkey.test_num
	end

	for _ = 1, 10000 do
		for _, monkey in ipairs(monkeys) do
			for _, item in ipairs(monkey.items) do
				local worry = item % product
				if monkey.op == '*' then
					worry = worry * (monkey.op_num or item)
				elseif monkey.op == '+' then
					worry = worry + (monkey.op_num or item)
				end
				-- if worry / monkey.test_num == worry // monkey.test_num then
				if math.fmod(worry, monkey.test_num) == 0 then
					table.insert(monkeys[monkey.true_dst + 1].items, worry)
				else
					table.insert(monkeys[monkey.false_dst + 1].items, worry)
				end
			end
			monkey.inspections = monkey.inspections + #monkey.items
			monkey.items = {}
		end
	end

	table.sort(monkeys, function(a,b) return a.inspections > b.inspections end)

	-- for i = 1, #monkeys do
	-- 	print(i, monkeys[i].inspections)
	-- end
	result = monkeys[1].inspections * monkeys[2].inspections

	if expected and result ~= expected then
		log.error('expected %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('11-input-test.txt', 10605))
log.report('part one      %d\n', partOne('11-input.txt', 50172))
log.report('part two test %d\n', partTwo('11-input-test.txt', 2713310158))
log.report('part two      %d\n', partTwo('11-input.txt', 11614682178))

--[[
$ time lua54 11.lua
Lua 5.4
part one test 10605
part one      50172
part two test 2713310158
part two      11614682178

real	0m0.277s
user	0m0.277s
sys	0m0.000s
]]