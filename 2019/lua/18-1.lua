-- https://adventofcode.com/2019/day/18

local log = require 'log'
local mk = require 'multikey'	-- http://siffiejoe.github.io/lua-multikey/

---@param y integer
---@param x integer
---@return integer
local function poskey(y,x)
	return y * 100 + x
end

local function putstate(t, state, val)
	mk.put(t, state.y, state.x, state.keys, val)
end

local function getstate(t, state)
	return mk.get(t, state.y, state.x, state.keys)
end

local function copystate(src)
	return {y=src.y, x=src.x, keys=src.keys}
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

---@class State
local start = {}	-- type State = struct{y, x: int, keys: string}

---@type string
local goal = ''		-- string of keys in starting maze

---@type integer
local y = 1			-- only used when building maze from input

for line in io.lines('18-input1.txt') do
	for x = 1, #line do
		local ch = line:sub(x,x)
		if ch == '@' then
			start = {y=y,x=x, keys=''}
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

log.trace('start %d,%d\ngoal %s\n', start.y, start.x, goal)

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

---@type table<string, integer>
local dist = {}
putstate(dist, start, 0)

---@type State[]
local q = {
	start
}

local directions = {
	{y=1, x=0},
	{y=0, x=1},
	{y=0, x=-1},
	{y=-1, x=0},
}

while #q > 0 do
	local state = table.remove(q, 1)

	if state.keys == goal then
		log.report('Part One %d\n', getstate(dist, state))	-- 4520
		return
	end

	for _, dir in ipairs(directions) do
		local nextState = {y=state.y + dir.y, x=state.x + dir.x, keys=state.keys}
		local nextTile = maze[poskey(nextState.y, nextState.x)]
		-- local nextTileb = nextTile:byte()
		if nextTile == '#' then
			goto continue
		end
		-- if nextTileb >= 65 and nextTileb <= 90 then
		if nextTile:match'%u' then
			if state.keys:find(nextTile:lower(), 1, true) == nil then
				goto continue
			end
		end
		-- if nextTileb >= 97 and nextTileb <= 122 then
		if nextTile:match'%l' then
			-- log.info('found key %s\n', nextTile)
			if nextState.keys:find(nextTile, 1, true) == nil then
				nextState.keys = sortstring(nextState.keys .. nextTile)
				-- log.info('added %s, keys now %s\n', nextTile, nextState.keys)
			end
		end
		if getstate(dist, nextState) == nil then
			putstate(dist, nextState, getstate(dist, state) + 1)
			q[#q+1] = nextState
		end
::continue::
	end
end

--[[
$ time lua54 18.lua
start 41,41
goal abcdefghijklmnopqrstuvwxyz
Part One 4520

real	2m21.676s
user	2m21.319s
sys	0m0.352s

$ time luajit 18.lua
start 41,41
goal abcdefghijklmnopqrstuvwxyz
Part One 4520

real	0m24.983s
user	0m24.706s
sys	0m0.276s
]]