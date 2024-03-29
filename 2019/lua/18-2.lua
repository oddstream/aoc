-- https://adventofcode.com/2019/day/18

local log = require 'log'
local mk = require 'multikey'	-- http://siffiejoe.github.io/lua-multikey/

---@param y integer
---@param x integer
---@return integer
local function poskey(y,x)
	return y * 100 + x
end

---@comment putstate uses values so doesn't need copystate
---@param val integer
local function putstate(t, state, val)
	mk.put(t,
		state.pos[1].y, state.pos[1].x,
		state.pos[2].y, state.pos[2].x,
		state.pos[3].y, state.pos[3].x,
		state.pos[4].y, state.pos[4].x,
		state.keys, state.active, val)
end

---@return integer?
local function getstate(t, state)
	return mk.get(t,
		state.pos[1].y, state.pos[1].x,
		state.pos[2].y, state.pos[2].x,
		state.pos[3].y, state.pos[3].x,
		state.pos[4].y, state.pos[4].x,
		state.keys, state.active)
end

local function copystate(src)
	local dst = {pos={{},{},{},{}}, keys=src.keys, active=src.active}
	for i=1,4 do
		dst.pos[i].y = src.pos[i].y
		dst.pos[i].x = src.pos[i].x
	end
	return dst
end

---@param s string
---@return string
local function sortstring(s)
	local tmp = {}
	for i=1,#s do
		tmp[i] = s:sub(i,i)
	end
	table.sort(tmp)
	return table.concat(tmp)
end

---@type table<integer, string>
local maze = {}		-- map[poskey(y,x)] = char

local start = {pos={{},{},{},{}}, keys='', active=1}

---@type string
local goal = ''		-- string of keys in starting maze

---@type integer
local y = 1			-- only used when building maze from input
local robots = 0	-- number of robots

for line in io.lines('18-input2.txt') do
	for x = 1, #line do
		local ch = line:sub(x,x)
		if ch == '@' then
			start.pos[robots+1] = {y=y,x=x}
			robots = robots + 1
			ch = '.'
		end
		if ch:match'%l' then
			goal = goal .. ch
		end
		maze[poskey(y,x)] = ch
	end
	y = y + 1
end
goal = sortstring(goal)

assert(robots==4)

log.trace('goal %s\n', goal)
for i = 1,4 do
	log.trace('start %d,%d\n', start.pos[i].y, start.pos[i].x)
end

--[[
do
	for y=1,82 do
		for x=1,82 do
			local ch = maze[poskey(y,x)]
			if ch then
				io.write(ch)
			end
		end
		io.write('\n')
	end
end
]]

local dist = {}
local q = {}

for i=1,4 do
	start.active = i
	putstate(dist, start, 0)
	table.insert(q, copystate(start))
end

local directions = {
	{y=1, x=0},
	{y=0, x=1},
	{y=0, x=-1},
	{y=-1, x=0},
}

while #q > 0 do
	-- if #q % 1000 == 0 then log.report('queue %d\n', #q) end	-- 28930

	local state = table.remove(q, 1)

	if state.keys == goal then
		log.report('Part Two %d\n', getstate(dist, state))	-- 1540
		return
	end

	for _, dir in ipairs(directions) do
		local nextState = copystate(state)
		-- assert(nextState.active==state.active)
		-- assert(nextState.keys==state.keys)
		nextState.pos[state.active].y = state.pos[state.active].y + dir.y
		nextState.pos[state.active].x = state.pos[state.active].x + dir.x
		local nextTile = maze[poskey(nextState.pos[state.active].y, nextState.pos[state.active].x)]
		if nextTile == '#' then
			goto nextdirection
		end
		if nextTile:match'%u' then
			if state.keys:find(nextTile:lower(), 1, true) == nil then
				goto nextdirection
			end
		end
		if nextTile:match'%l' then
			-- log.info('found key %s\n', nextTile)
			if nextState.keys:find(nextTile, 1, true) == nil then
				nextState.keys = sortstring(nextState.keys .. nextTile)
				-- log.info('added %s, now %s\n', nextTile, nextState.keys)
			end
		end
		for i=1,4 do
			if (i ~= state.active) and (nextState.keys == state.keys) then
				goto nextrobot
			end
			nextState.active = i
			if getstate(dist, nextState) == nil then
				putstate(dist, nextState, getstate(dist, state) + 1)
				table.insert(q, copystate(nextState))
			end
::nextrobot::
		end

::nextdirection::
	end
end

--[[
$ time lua54 18-2.lua
Part Two 1540
#queue 28930

real	72m14.453s
user	62m49.849s
sys	0m2.673s


$ time luajit 18-2.lua
#queue 24000
PANIC: unprotected error in call to Lua API (not enough memory)

real	1m47.719s
user	1m47.015s
sys	0m0.684s

]]