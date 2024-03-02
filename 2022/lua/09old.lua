local inputFilename = "09-input.txt"

local log = require 'log'

function _G.table.copy(t)
	local u = { }
	for k, v in pairs(t) do u[k] = v end
	return setmetatable(u, getmetatable(t))
end

local function key(x, y)
	return ("%d,%d"):format(x, y)
end

local directions = {
	L = {x = -1, y = 0},
	R = {x = 1, y = 0},
	U = {x = 0, y = -1},
	D = {x = 0, y = 1},
}

local hvisited, tvisited = {}, {}

local hpos = {x=30, y=30}	-- enough to stop any coord going -ve
local tpos = table.copy(hpos)
local hcount, tcount = 0, 0

for line in io.lines(inputFilename) do
	local dir, num = line:match("(%u) (%d*)")
	num = tonumber(num)
	while num > 0 do
		local hprev = table.copy(hpos)
		hpos.x = hpos.x + directions[dir].x
		hpos.y = hpos.y + directions[dir].y
		num = num - 1
		if math.abs(tpos.x - hpos.x) >= 2 or math.abs(tpos.y - hpos.y) >= 2 then
			tpos = table.copy(hprev)
		end
		local hkey = key(hpos.x, hpos.y)
		if not hvisited[hkey] then
			hvisited[hkey] = true
			hcount = hcount + 1
		end
		local tkey = key(tpos.x, tpos.y)
		if not tvisited[tkey] then
			tvisited[tkey] = true
			tcount = tcount + 1
		end
	end
end

log.report("%d visited by head\n", hcount)
log.report("%d visited by tail\n", tcount)	-- 13, 5619