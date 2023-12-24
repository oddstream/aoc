-- https://adventofcode.com/2022/day/13 Distress Signal

local log = require 'log'

---thank you https://github.com/MrSimbax/advent-of-code-2022/blob/main/day_13.lua
---@param a table
---@param b table
---@return boolean?
local function lt (a, b)
    local ta = type(a)
    local tb = type(b)
    if ta == "number" and tb == "number" then
        return a < b or a == b and nil
    elseif ta == "number" and tb == "table" then
        return lt({a}, b)
    elseif ta == "table" and tb == "number" then
        return lt(a, {b})
    elseif ta == "table" and tb == "table" then
        for i = 1, #a do
            local r = lt(a[i], b[i])
            if r ~= nil then
                return r
            end
        end
        return #a < #b and true or nil
    elseif tb == "nil" then
        return false
    end
end

---don't do this at home
---@param s string
---@return table
local function listify(s)
	s = s:gsub('%[', '{')
	s = s:gsub('%]', '}')
	local fn, err = load('return ' .. s)
	if err then print(err) end
	return fn()
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	local index = 1
	local f = assert(io.open(filename, 'r'))
	while true do
		local left = f:read'line'
		if not left then break end
		local right = f:read'line'

		left = listify(left)
		right = listify(right)

		if lt(left, right) then
			result = result + index
		end

		local blank = f:read'l'
		if not blank then break end

		index = index + 1
	end
	f:close()

	if expected and result ~= expected then
		log.error('expecting %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result

	local packets = {{{2}}, {{6}}}
	local f = assert(io.open(filename, 'r'))
	while true do
		local packet = f:read'line'
		if not packet then break end
		if #packet > 0 then
			packets[#packets+1] = listify(packet)
		end
	end
	f:close()

	table.sort(packets, function(a,b) return lt(a,b) end)

	local two, six
	for i, packet in ipairs(packets) do
		if type(packet) == 'table' and #packet == 1 then
			if type(packet[1]) == 'table' and #packet[1] == 1 then
				if packet[1][1] == 2 then
					two = i
				elseif packet[1][1] == 6 then
					six = i
				end
			end
		end
	end

	result = two * six

	if expected and result ~= expected then
		log.error('expecting %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('13-input-test.txt', 13))
log.report('part one      %d\n', partOne('13-input.txt', 4894))
log.report('part two test %d\n', partTwo('13-input-test.txt', 140))
log.report('part two      %d\n', partTwo('13-input.txt', 24180))

--[[
$ time luajit 13.lua
Lua 5.1
part one test 13
part one      4894
part two test 140
part two      24180

real	0m0.041s
user	0m0.027s
sys	0m0.012s

$ time lua54 13.lua
Lua 5.4
part one test 13
part one      4894
part two test 140
part two      24180

real	0m0.035s
user	0m0.031s
sys	0m0.004s
]]