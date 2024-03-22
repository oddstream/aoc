-- https://adventofcode.com/2019/day/15

local log = require 'log'

local POSITION, IMMEDIATE, RELATIVE = '0', '1', '2'
local NORTH, SOUTH, WEST, EAST = 1, 2, 3, 4
local WALL, STEP, OXYGEN = 0, 1, 2

local toutesDirections = {
	NORTH = {y=-1, x=0},
	SOUTH = {y=1, x=0},
	WEST = {y=0, x=-1},
	EAST = {y=0, x=1},
}

---@param input string?
---@return string[]
local function load(input)

	if input == nil then
		local f = assert(io.open('15-input.txt', 'r'))
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

---@param master string[]
---@return string[]
local function clone(master)
	local program = {}
	for i, n in ipairs(master) do
		program[i] = n
	end
	-- The computer's available memory should be much larger than the initial program.
	-- Memory beyond the initial program starts with the value 0
	-- and can be read or written like any other memory.
	-- BUT Lua doesn't seem to care about any extra memory
	return program
end

local function key(y, x)
	return string.format('%d,%d', y, x)
end

local function unkey(k)
	local y, x = k:match('([0-9-]+),([0-9-]+)')
	return tonumber(y), tonumber(x)
end

local function run(bot)
	---return the next data item, advance instruction pointer
	---@return integer
	local function read()
		local v = bot.data[bot.ip]	-- the instruction pointer is already 1-based
		bot.ip = bot.ip + 1
		return v
	end

	local function store(loc, val)
		bot.data[loc+1] = val	-- +1 because Lua arrays are 1-based
	end

	---@comment tied in knots over the 203 problem, and this makes no sense
	---@param mode string
	---@param parameter integer
	---@return integer
	local function getAddress(mode, parameter)
		if mode == POSITION then
			return parameter
		elseif mode == RELATIVE then
			return bot.relbase + parameter
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
		return bot.data[address+1]
	end

	while true do
		local header = string.format('%05d', read())

		local opcode = header:sub(4,5)

		if opcode == '99' then
			bot.halted = true
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
			store(getAddress(modeA, parameterA), bot.input)
			-- print('INPUT', bot.input)
			goto continue
		end

		if opcode == '04' then
			bot.output = valueA
			if valueA == -1 then print('OUTPUT', valueA) end
			return bot
		end

		if opcode == '09' then
			bot.relbase = bot.relbase + valueA
			goto continue
		end

		-- opcodes with 2 parameters

		local parameterB = read()
		local valueB = getValue(modeB, parameterB)

		if opcode == '05' then
			if valueA ~= 0 then
				bot.ip = valueB + 1
			end
			goto continue
		end
		if opcode == '06' then
			if valueA == 0 then
				bot.ip = valueB + 1
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

local function walk(map, bots)

	local function neighbour(direction, coord, bot, futures)
		if map[coord] ~= nil then
			return
		end
		map[coord] = 'r' -- reserved?
		futures[#futures+1] = coord
		local newbot = {
			data = clone(bot.data),
			ip = bot.ip,
			relbase = bot.relbase,
			input = direction,
			output = bot.output,
		}
		bots[coord] = run(newbot)
	end

	local futures = {key(0,0)}	-- list of coord
	local distance = 0
	while #futures > 0 do
		local current = futures
		futures = {}
		for _, coord in pairs(current) do
			local bot = bots[coord]
			if bot.output == OXYGEN then
				map[coord] = 'O'
				log.report('part 1 oxygen at %s distance %d\n', coord, distance)
			elseif bot.output == WALL then
				map[coord] = '#'
			elseif bot.output == STEP or bot.output == -1 then	-- mystery how this this -1 gets here, from an un-run bot
				map[coord] = '.'
				local y, x = unkey(coord)
				neighbour(NORTH, key(y-1, x), bot, futures)
				neighbour(SOUTH, key(y+1, x), bot, futures)
				neighbour(WEST, key(y, x-1), bot, futures)
				neighbour(EAST, key(y, x+1), bot, futures)
			else
				log.error('unknown bot output %d\n', bot.output)
				return
			end
		end
		distance = distance + 1
	end
end

---@comments there are cleverer ways of doing this, but this was quick to write, easy to read and fast enough
---@param map table
---@return integer
local function fill(map)

	local minutes = 0

	while true do
		local lst = {}
		for k,v in pairs(map) do
			if v == 'O' then
				local y, x = unkey(k)
				for _, vv in pairs(toutesDirections) do
					local yy, xx = y + vv.y, x + vv.x
					local kk = key(yy,xx)
					if map[kk] == '.' then
						lst[#lst+1] = kk
					end
				end
			end
		end
		if #lst == 0 then
			break
		end
		for _, k in ipairs(lst) do
			map[k] = 'O'
		end
		minutes = minutes + 1
	end

	return minutes
end

--[[
local function display(map)
	for y = -2, 20 do
		for x = -20, 25 do
			local k = key(y, x)
			if map[k] == nil then
				io.write(' ')
			else
				io.write(map[k])
			end
		end
		io.write('\n')
	end
end
]]

local master = load()
local bots = {} -- indexed by coord key
local map = {}	-- indexed by coord key, fill with '.' or '#' or 'O'
bots[key(0,0)] = {
	data = clone(master),
	ip = 1,
	relbase = 0,
	output = -1,
}
walk(map, bots)	-- 220
local result = fill(map)
log.report('part 2 %d\n', result)	-- 334

--[[
$ time luajit 15.lua
part 1 oxygen at 12,16 distance 220
part 2 334

real	0m0.126s
user	0m0.106s
sys	0m0.020s
]]