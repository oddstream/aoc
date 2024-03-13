local log = require 'log'

local f = assert(io.open("03-input.txt", "r"))
local content = f:read("*a")
f:close()

local function key(x, y)
	return ("%d,%d"):format(x, y)
end

local function partOne()
	local sx, sy = 1000, 1000
	local visited = {[key(sx,sy)] = true}
	local count = 1

	local function visit(n, x, y)
		local c = content:sub(n,n)
		if c == '<' then
			x = x - 1
		elseif c == '>' then
			x = x + 1
		elseif c == '^' then
			y = y - 1
		elseif c == 'v' then
			y = y + 1
		else
			log.error('unknown input \'%s\'\n', c)
		end
		local k = key(x,y)
		if not visited[k] then
			visited[k] = true
			count = count + 1
		end
		return x, y
	end

	for i = 1, #content do
		sx, sy = visit(i, sx, sy)
	end
	log.report('part one %d\n', count)
end

local function partTwo()
	local sx, sy = 1000, 1000
	local rx, ry = 1000, 1000
	local visited = {[key(sx,sy)] = true}
	local count = 1

	local function visit(n, x, y)
		local c = content:sub(n,n)
		if c == '<' then
			x = x - 1
		elseif c == '>' then
			x = x + 1
		elseif c == '^' then
			y = y - 1
		elseif c == 'v' then
			y = y + 1
		else
			log.error('unknown input \'%s\'\n', c)
		end
		local k = key(x,y)
		if not visited[k] then
			visited[k] = true
			count = count + 1
		end
		return x, y
	end

	for i = 1, #content, 2 do
		sx, sy = visit(i, sx, sy)
		rx, ry = visit(i+1, rx, ry)
	end
	log.report('part two %d\n', count)
end

log.report('%s\n', _VERSION)
partOne() -- 2081
partTwo() -- 2341