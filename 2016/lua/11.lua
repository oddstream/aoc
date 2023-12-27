-- https://adventofcode.com/2016/day/11 Radioisotope Thermoelectric Generators
--
-- https://eddmann.com/posts/advent-of-code-2016-day-11-radioisotope-thermoelectric-generators/

local log = require 'log'

-- the element types DO NOT have unique first letters
-- hydrogen, lithium
-- plutonium, promethium, ruthenium, strontium, thulium
-- the device types have unique first letters (generator, microchip)
-- 4 objects in test, 14 objects in input

-- simple list of {f=floor:integer e=element:string d=device:string}
-- immutable (except just after cloning)
-- piggyback key, moves, floor

---this may not work for part 2, revert to a string of 1..4
---integer key: 30s, string key: 48s (part 1, LuaJIT)
---@param state table
---@return integer
local function key(state)
	local k = state.floor
	for _, item in ipairs(state) do
		k = k * 10 + item.f
	end
	-- local k = tostring(state.floor)
	-- for _, item in ipairs(state) do
	-- 	k = k .. tostring(item.f)
	-- end
	return k
end

---@param filename string
---@return table
local function loadInput(filename)
	local state = {}

	local floor = 1
	for line in io.lines(filename) do
		line = line:gsub('-compatible', '')	-- remove '-compatible'
		for _, device in ipairs({'generator', 'microchip'}) do
			repeat
				local element = line:match('(%l+) ' .. device) -- eg 'lithium generator'
				if element ~= nil then
					line, _ = line:gsub(element .. ' ' .. device, '') -- remove 'lithium generator'
					table.insert(state, {f=floor, e=element, d=device})
					-- print(floor, element, device)
				end
			until element == nil
		end
		floor = floor + 1
	end
	-- "you and the elevator will start on the first floor"
	state.floor = 1
	state.key = key(state)
	return state
end

---@param state table
---@return table
local function clone(state)
	local newstate = {}
	for i, v in ipairs(state) do
		newstate[i] = {f=v.f, e=v.e, d=v.d}	-- fed
	end
	newstate.floor = state.floor
	return newstate
end

---@param state table
local function display(state)
	print('state ', state.key)
	for floor = 1, 4 do
		io.write(floor)
		if floor == state.floor then
			io.write('*')
		else
			io.write(' ')
		end
		for _, v in ipairs(state) do
			if v.f == floor then
				io.write(v.e, ' ', v.d, ' ')
			end
		end
		io.write('\n')
	end
	io.write('\n')
end

---@param state table
---@return boolean
local function valid(state)
	-- if state.floor < 1 or state.floor > 4 then
	-- 	return false
	-- end
	-- "if a chip is ever left in the same area as another RTG,
	-- and it's not connected to its own RTG, the chip will be fried."

	-- "An RTG powering a microchip is still dangerous to other microchips."

	for _, a in ipairs(state) do
		-- if lift/player is on same floor
		-- them items are not "left"
		-- so any combination is ok?
		if a.f ~= state.floor and a.d == 'microchip' then
			-- we found an unattended microchip
			local found_another = false
			for _, b in ipairs(state) do
				if b.f == a.f and b.e ~= a.e and b.d == 'generator' then
					found_another = true
					break
				end
			end
			if found_another then
				-- there is a hostile generator on the same floor as the microchip
				local found_own = false
				for _, b in ipairs(state) do
					if b.f == a.f and b.e == a.e and b.d == 'generator' then
						found_own = true
						break
					end
				end
				if not found_own then
					-- the microchip does not have a protecting generator
					return false
				end
			end
		end
	end
	return true
end

---@param num integer
---@return boolean
local function is_all_4s(num)
	while num > 0 do
		local remainder = num % 10
		if remainder ~= 4 then
			return false
		end
		num = math.floor(num / 10)
	end
	return true
end

---complete when rows 1 .. 3 are empty
---@param state table
---@return boolean
local function complete(state)
	-- return is_all_4s(state.key)
	for _, item in ipairs(state) do
		if item.f ~= 4 then
			return false
		end
	end
	return true
end

---yield all combinations of pairs of numbers from start to stop, inclusive
---eg 0 2 -> 0 1, 0 2, 1 2
---@param start integer
---@param stop integer
---@return function (iterator)
local function combinations(start, stop)
	local function combin()
		for i = start, stop do
			for j = i+1, stop do
				coroutine.yield(i, j)
			end
		end
	end
	local co = coroutine.create(function() combin() end)
	return function()
		local success, i, j = coroutine.resume(co)
		if success then return i, j end
	end
end

-- for i, j in combinations(1, 2) do
-- 	print(i, j)
-- end

---generate new states from input state
---@param state table
---@return function iterator
local function permutations(state)

	-- "Its capacity rating means it can carry at most yourself and two RTGs or microchips in any combination."
	-- "the elevator will only function if it contains at least one RTG or microchip"

--[[
	If you can move two items upstairs, don't bother bringing one item upstairs.
	If you can move one item downstairs, don't bother bringing two items downstairs.
	(But still try to move items downstairs even if you can move items upstairs)

	ALL PAIRS ARE INTERCHANGEABLE - The following two states are EQUIVALENT:
	(HGen@floor0, HChip@floor1, LGen@floor2, LChip@floor2),
	(LGen@floor0, LChip@floor1, HGen@floor2, HChip@floor2)
	prune any state EQUIVALENT TO (not just exactly equal to) a state you have already seen!
]]
	local function permute()
		local atfloor = {} -- indexes to items at this floor
		for i, item in ipairs(state) do
			if item.f == state.floor then
				table.insert(atfloor, i)
			end
		end
		if #atfloor == 0 then
			print'#atfloor zero'
			return	-- current floor should always have some items
		end
		-- consider moving items up and down one floor
		for _, newf in ipairs({state.floor + 1, state.floor - 1}) do
			if newf >= 1 and newf <= 4 then
				assert(math.abs(state.floor - newf)==1)
				-- try all combinations of two items on this floor
				for i1, i2 in combinations(0, #atfloor) do
					local n1, n2
					if i1 ~= 0 then
						n1 = atfloor[i1]
					end
					n2 = atfloor[i2]
					local newstate = clone(state)
					newstate.floor = newf
					if i1 ~= 0 then
						newstate[n1].f = newf
					end
					newstate[n2].f = newf
					newstate.key = key(newstate)
					coroutine.yield(newstate)
				end
		end
		end
	end
	local co = coroutine.create(function() permute() end)
	return function()
		local success, newstate = coroutine.resume(co)
		if success then return newstate end
	end
end

---@param start table
---@return integer
local function bfs(start)
	assert(not complete(start))
	local seen = {}
	seen[start.key] = true
	start.moves = 0
	local q = {start}
	while #q > 0 do
		-- try sorting so states nearest complete get processed first
		-- time with sort: 29 minutes (wrong answer, 31 instead of 35)
		-- time without sort: 33 seconds (right answer)
		-- returns true when the first element must come before the second in the final order
		-- table.sort(q, function(a,b) return a.key > b.key end)
		local state = table.remove(q, 1)
		-- io.write(state.key, ' ')
		-- if state.key == 43444 then
			-- display(state)
		-- end
		for newstate in permutations(state) do
			--[[
			if valid(newstate) and not seen[newstate.key] then
				seen[newstate.key] = true
				newstate.moves = state.moves + 1
				if complete(newstate) then
					return newstate.moves
				end
				q[#q + 1] = newstate
			end
			]]
			if not seen[newstate.key] then
				--  disregard all previously seen states, valid and invalid
				seen[newstate.key] = true
				if valid(newstate) then
					newstate.moves = state.moves + 1
					-- we can return immediately as soon as we find a completed state
					-- no need to push it onto the q
					if complete(newstate) then
						return newstate.moves
					end
					-- push this state onto the q so we can run permutation off it
					q[#q + 1] = newstate
				end
			end
		end
	end
	return -1
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result

	local initial_state = loadInput(filename)
	display(initial_state)

	assert(valid(initial_state))
	assert(not complete(initial_state))

	result = bfs(initial_state)

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

	local initial_state = loadInput(filename)
	table.insert(initial_state, {f=1, e='elerium', d='generator'})
	table.insert(initial_state, {f=1, e='elerium', d='microchip'})
	table.insert(initial_state, {f=1, e='dilithium', d='generator'})
	table.insert(initial_state, {f=1, e='dilithium', d='microchip'})
	initial_state.key = key(initial_state)

	display(initial_state)

	assert(valid(initial_state))
	assert(not complete(initial_state))
	print('initial key', initial_state.key)

	result = bfs(initial_state)

	if expected ~= nil and result ~= expected then
		log.error('expecting %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
-- log.report('part one test %d\n', partOne('11-test.txt', 11))
log.report('part one      %d\n', partOne('11-input.txt', 31)) -- not 20, 21, 22, 23, 24
-- log.report('part two test %d\n', partTwo('11-test.txt', 0))
log.report('part two      %d\n', partTwo('11-input.txt', 55))

--[[
	LuaJIT won't run part 2
	(PANIC: unprotected error in call to Lua API (not enough memory))
	even with keys as strings instead of numbers
	The total memory of LuaJIT objects (tables, strings) is limited to about 2GB.
	https://stackoverflow.com/questions/35155444/why-is-luajits-memory-limited-to-1-2-gb-on-64-bit-platforms
]]
