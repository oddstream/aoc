-- https://adventofcode.com/2022/day/17

-- ProggyVector

-- https://www.reddit.com/r/adventofcode/comments/znykq2/comment/j0l2dv4/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
-- https://github.com/MrSimbax/advent-of-code-2022/blob/main/day_17.lua
-- https://jactl.io/blog/2023/04/22/advent-of-code-2022-day17.html

local log = require 'log'

--[[
local Point, PointMeta = {}, {}
PointMeta.__add = function(a,b) return Point:new(a.x+b.x, a.y+b.y) end
function Point:new(x,y)
	return setmetatable({x=x or 0, y=y or 0}, PointMeta)
end
]]

local jets
do
	local f = assert(io.open('17-test1.txt', "r"))
	jets = f:read("*all")
	f:close()
end

log.trace('%d jets\n', #jets)

-- local WIDTH = 7
-- local DOWN = {x=0,y=-1}

local jetiter = function()
	local i = 1
	local n = #jets
	while true do
		local ch = jets:sub(i,i)
		if ch == '<' then
			coroutine.yield({x=-1,y=0})
		elseif ch == '>' then
			coroutine.yield({x=1,y=0})
		else
			log.error('unknown jet %s\n', ch)
		end
		i = i % n + 1
	end
end

local rocks = {
	-- coords origin is top left hand corner (0,0)
	-- x increases rightwards
	-- y decreases downwards
	{shape='hline',  width=4, height=1, {x=0,y=0}, {x=1,y= 0}, {x=2,y= 0}, {x=3,y= 0}},
	{shape='plus',   width=3, height=3, {x=1,y=0}, {x=0,y=-1}, {x=1,y=-1}, {x=2,y=-1}, {x=1,y=-2}},
	{shape='corner', width=3, height=3, {x=2,y=0}, {x=2,y=-1}, {x=0,y=-2}, {x=1,y=-2}, {x=2,y=-2}},
	{shape='vline',  width=1, height=4, {x=0,y=0}, {x=0,y=-1}, {x=0,y=-2}, {x=0,y=-3}},
	{shape='square', width=2, height=2, {x=0,y=0}, {x=0,y=-1}, {x=1,y= 0}, {x=1,y=-1}},
}

local rocksiter = function()
	local i = 1
	local n = #rocks
	while true do
		coroutine.yield(i, rocks[i])
		i = i % n + 1
	end
end

-- local function ingrid(pos, rock)
-- 	return 1 <= pos.x and pos.x + rock.width - 1 <= WIDTH and 1 <= pos.y - rock.height + 1
-- end

-- local function isblocked(grid, pos, rock)
-- 	for _, offset in ipairs(rock) do
-- 	end
-- end

for _, r in pairs(rocks) do
	print(#r, r[1].x, r[1].y, r.width, r.height)
end

local nextrock = coroutine.wrap(rocksiter)
for _=1,10 do
	local irock, rock = nextrock()
	print(irock, rock.shape)
end

local nextjet = coroutine.wrap(jetiter)
for _=1,10 do
	local dir = nextjet()
	print(dir.x, dir.y)
end
