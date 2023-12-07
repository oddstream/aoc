-- https://adventofcode.com/2016/day/11 Radioisotope Thermoelectric Generators
--
-- https://eddmann.com/posts/advent-of-code-2016-day-11-radioisotope-thermoelectric-generators/

local log = require 'log'

if not _G.table.contains then
function _G.table.contains(tab, val)
	for index, value in ipairs(tab) do
		if value == val then
			return true, index
		end
	end
	return false, 0
end
end

if not _G.table.equals then
---@param o1 any|table First object to compare
---@param o2 any|table Second object to compare
function _G.table.equals(o1, o2)
    if o1 == o2 then return true end
    local o1Type = type(o1)
    local o2Type = type(o2)
    if o1Type ~= o2Type then return false end
    if o1Type ~= 'table' then return false end

    local keySet = {}

    for key1, value1 in pairs(o1) do
        local value2 = o2[key1]
        if value2 == nil or _G.table.equals(value1, value2) == false then
            return false
        end
        keySet[key1] = true
    end

    for key2, _ in pairs(o2) do
        if not keySet[key2] then return false end
    end
    return true
end
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

local function deepcopy(orig)
    local orig_type = type(orig)
    local copy
    if orig_type == 'table' then
        copy = {}
        for orig_key, orig_value in next, orig, nil do
            copy[deepcopy(orig_key)] = deepcopy(orig_value)
        end
        setmetatable(copy, deepcopy(getmetatable(orig)))
    else -- number, string, boolean, etc
        copy = orig
    end
    return copy
end

-- the element types have unique first letters
-- the device types have unique first letters (g, m)

local function readInput(filename)
	local floors = {}

	local floor = 1
	for line in io.lines(filename) do
		-- Lua does not have POSIX regualr expressions, so we have to match n' cut ...
		table.insert(floors, {})
		line, _ = line:gsub('-compatible', '')
		for _, device in ipairs({'generator', 'microchip'}) do
			repeat
				local element = line:match('(%l+) ' .. device)
				if element ~= nil then
					line, _ = line:gsub(element .. ' ' .. device, '')
					table.insert(floors[floor], element:sub(1,1) .. device:sub(1,1))
					print(floor, element, device)
				end
			until element == nil
		end
		floor = floor + 1
	end
	return floors
end

---@return boolean
local function isValid(state)
	for floor = 1, #state do
		-- create a list of all microchips on this floor
		local microchips = {}
		for _, item in ipairs(state[floor]) do
			if item:sub(2,2) == 'm' then
				table.insert(microchips, item)
			end
		end
		if #microchips > 0 then
			-- create a list of all generators on this floor
			local generators = {}
			for _, item in ipairs(state[floor]) do
				if item:sub(2,2) == 'g' then
					table.insert(generators, item)
				end
			end
			if #generators > 0 then
				-- for each microchip on the floor ...
				for _, microchip in ipairs(microchips) do
					-- is there a corresponding generator on this floor?
					local gen = microchip:sub(1,1) .. 'g'
					if not table.contains(generators, gen) then
						print('no', gen, 'on floor', floor)
						return false
					end
				end
			end
		end
	end
	return true
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	local initial_state = readInput(filename)
	initial_state.visited = false -- piggy back

	print('initial', isValid(initial_state))

	local final_state = {{},{},{},{}}
	for i = 1, #initial_state - 1 do
		for _, v in ipairs(initial_state[i]) do
			table.insert(final_state[4], v)
		end
	end
	final_state.visited = false -- piggy back

	print('final', isValid(final_state))

	local t2 = deepcopy(initial_state)
	local t3 = shallowcopy(initial_state)
	print(_G.table.equals(t2, initial_state))
	print(_G.table.equals(t2, t3))

	if expected ~= nil and result ~= expected then
		log.error('part one should be %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result = 0

	for line in io.lines(filename) do
		log.trace(line)
	end

	if expected ~= nil and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('11-test.txt', 0))
-- log.report('part one      %d\n', partOne('11-input.txt', 0))
-- log.report('part two test %d\n', partTwo('11-test.txt', 0))
-- log.report('part two      %d\n', partTwo('11-input.txt', 0))

