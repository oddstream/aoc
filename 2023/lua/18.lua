-- https://adventofcode.com/2023/day/

local log = require 'log'

-- input contains 666 lines
-- largest n appears to be 12

local cornerMap = {
	['UR'] = '┌',
	['LD'] = '┌',

	['RU'] = '┘',
	['DL'] = '┘',

	['DR'] = '└',
	['LU'] = '└',

	['RD'] = '┐',
	['UL'] = '┐',
}

---@param filename string
---@param part integer
---@param expected? integer
---@return integer
local function partOneAndTwo(filename, part, expected)
	local result = 0

	local function key(x,y)
		return tostring(x) .. ',' .. tostring(y)
	end

	local dirs = {U={x=0, y=-1}, R={x=1, y=0}, L={x=-1, y=0}, D={x=0, y=1}}
	local walls = {}
	local x, y = 1, 1
	local prevdir = 'U' -- TODO KLUDGE
	local minx, miny = math.huge, math.huge
	local maxx, maxy = -math.huge, -math.huge
	for line in io.lines(filename) do
		local dir, n, hex = line:match'([UDLR]) (%d+) %(%#(%x+)%)'
		if part == 2 then
			local hexdist = tonumber(hex:sub(1,5), 16)
			local hexdir = hex:sub(6,6)
			if hexdir == '0' then
				hexdir = 'R'
			elseif hexdir == '1' then
				hexdir = 'D'
			elseif hexdir == '2' then
				hexdir = 'L'
			elseif hexdir == '3' then
				hexdir = 'U'
			end
			n = hexdist
			dir = hexdir
		end
		n = tonumber(n)
		local ch
		if dir == 'R' or dir == 'L' then
			ch = '─'
		else
			ch = '│'
		end
		local corner = cornerMap[prevdir .. dir]
		for i = 1, n do
			if i == 1 and corner then
				walls[key(x,y)] = corner
			else
				walls[key(x,y)] = ch -- and/or a color
			end
			x = x + dirs[dir].x
			y = y + dirs[dir].y

			if x < minx then minx = x end
			if y < miny then miny = y end
			if x > maxx then maxx = x end
			if y > maxy then maxy = y end
		end
		prevdir = dir
	end

	-- for yy = miny, maxy do
	-- 	for xx = minx, maxx do
	-- 		local ch = walls[key(xx,yy)]
	-- 		if ch then
	-- 			io.write(ch)
	-- 		else
	-- 			io.write('.')
	-- 		end
	-- 	end
	-- 	io.write('\n')
	-- end
	print('min x y', minx, miny, 'max x y', maxx, maxy)

	for _, _ in pairs(walls) do result = result + 1 end
	print('number of walls before filling', result)

	-- could do BFS floodfill or shoelace formula,
	-- but instead build the map with corners and do a day 10
	for yy = miny, maxy do
		print(miny, yy, maxy)
		local inside = false
		for xx = minx, maxx do
			local wk = key(xx,yy)
			local ch = walls[wk] -- could be nil
			if ch == '│' or ch == '└' or ch == '┘' then
				inside = not inside
			else
				if inside and not ch then
					--walls[wk] = '#'
					result = result + 1
				end
			end
		end
	end

	-- for yy = miny, maxy do
	-- 	for xx = minx, maxx do
	-- 		local ch = walls[key(xx,yy)]
	-- 		if ch then
	-- 			io.write(ch)
	-- 		else
	-- 			io.write('.')
	-- 		end
	-- 	end
	-- 	io.write('\n')
	-- end

	-- for _, _ in pairs(walls) do result = result + 1 end

	if expected and result ~= expected then
		log.error('expected %d\n', expected)
	end
	return result
end

---@param filename string
---@param part integer
---@param expected? integer
---@return integer
local function shoelace(filename, part, expected)
	local result

	-- https://rosettacode.org/wiki/Shoelace_formula_for_polygonal_area#Lua

	---LuaJIT fails with PANIC: unprotected error in call to Lua API (not enough memory)
	local function shoeArea1(ps)
		local function det2(i,j)
		  return ps[i][1]*ps[j][2]-ps[j][1]*ps[i][2]
		end
		local sum = #ps>2 and det2(#ps,1) or 0
		for i=1,#ps-1 do sum = sum + det2(i,i+1)end
		return math.abs(0.5 * sum)
	end

	---Fails with 'too many results to unpack' on 5.4 and LuaJIT
	local function shoeArea2(ps)
		local function ssum(acc, p1, p2, ...)
			if not p2 or not p1 then
				return math.abs(0.5 * acc)
			else
				return ssum(acc + p1[1]*p2[2]-p1[2]*p2[1], p2, ...)
			end
		end
		return ssum(0, ps[#ps], table.unpack(ps))
	end

	local dirs = {U={x=0, y=-1}, R={x=1, y=0}, L={x=-1, y=0}, D={x=0, y=1}}
	local walls = {}
	local x, y = 1,1
	for line in io.lines(filename) do
		local dir, n, hex = line:match'([UDLR]) (%d+) %(%#(%x+)%)'
		if part == 2 then
			local hexdist = tonumber(hex:sub(1,5), 16)
			local hexdir = hex:sub(6,6)
			if hexdir == '0' then
				hexdir = 'R'
			elseif hexdir == '1' then
				hexdir = 'D'
			elseif hexdir == '2' then
				hexdir = 'L'
			elseif hexdir == '3' then
				hexdir = 'U'
			end
			n = hexdist
			dir = hexdir
		end
		n = tonumber(n)
		for _ = 1, n do
			walls[#walls+1] = {x, y}
			x = x + dirs[dir].x
			y = y + dirs[dir].y
		end
	end

	print('loaded', #walls, 'walls')

	-- I can't seem to wrap my head around why the perimeter needs to be divided by 2.
	-- Could you please explain why?
	-- Because half of the thick line is already included in the (computed) interior area.
	result = shoeArea1(walls) + (#walls / 2) + 1

	if expected and result ~= expected then
		log.error('expected %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', shoelace('18-test.txt', 1, 62))
log.report('part one      %d\n', shoelace('18-input.txt', 1, 46359))
log.report('part two test %d\n', shoelace('18-test.txt', 2, 952408144115))
log.report('part two      %d\n', shoelace('18-input.txt', 2, 59574883048274))
-- log.report('part one test %d\n', partOneAndTwo('18-test.txt', 1, 62))
-- log.report('part one      %d\n', partOneAndTwo('18-input.txt', 1, 46359))
-- log.report('part two test %d\n', partOneAndTwo('18-test.txt', 2, 952408144115))
-- log.report('part two      %d\n', partOneAndTwo('18-input.txt', 2, 0))

--[[
$ time lua51 18.lua
Lua 5.1
loaded	38	walls
part one test 62
loaded	3576	walls
part one      46359
loaded	6405262	walls
part two test 952408144115
/home/gilbert/lua-5.1.5/src/lua: 18.lua:189: table overflow
stack traceback:
	18.lua:189: in function 'shoelace'
	18.lua:212: in main chunk
	[C]: ?

real	3m55.115s
user	3m49.643s
sys	0m5.287s

$ time lua54 18.lua
Lua 5.4
loaded	38	walls
part one test 62
loaded	3576	walls
part one      46359
loaded	6405262	walls
part two test 952408144115
loaded	156391032	walls
part two      59574883048274

real	0m41.495s
user	0m36.481s
sys	0m5.004s

$ time luajit 18.lua
Lua 5.1
loaded	38	walls
part one test 62
loaded	3576	walls
part one      46359
loaded	6405262	walls
part two test 952408144115
PANIC: unprotected error in call to Lua API (not enough memory)

real	0m5.048s
user	0m4.415s
sys	0m0.632s
]]