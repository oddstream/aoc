-- https://adventofcode.com/2019/day/20

local function key(y,x)
	return y*1000+x
end

local function unkey(k)
	return math.floor(k / 1000), k % 1000
end

local function key3(y,x,level)
	return string.format('%d,%d,%d', y, x, level)
end

local directions = {
	{y=-1, x=0},	-- up
	{y=1, x=0},		-- down
	{y=0, x=-1},	-- left
	{y=0, x=1},		-- right
}

local maze = {}
local pos2label = {}	-- positions will be unique, labels will not
local portals = {}
local start, finish = {}, {}

for line in io.lines('20-input.txt') do
	local row = {}
	for ch in line:gmatch'.' do
		row[#row+1] = ch
	end
	maze[#maze+1] = row
end

for y1=1, #maze do
	for x1=1, #maze[y1] do
		local ch1 = maze[y1][x1]
		if ch1:match'%u' then
			for _, dir in ipairs(directions) do
				local y2, x2 = y1 + dir.y, x1 + dir.x
				if y2 > 0 and x2 > 0 and y2 <= #maze and x2 <= #maze[y2] then
					local ch2 = maze[y2][x2]
					if ch2:match'%u' then
						local y3, x3 = y2 + dir.y, x2 + dir.x
						if y3 > 0 and x3 > 0 and y3 <= #maze and x3 <= #maze[y3] then
							local ch3 = maze[y3][x3]
							if ch3 == '.' then
								if ch1 == 'A' and ch2 == 'A' then
									start.y = y3
									start.x = x3
								elseif ch1 == 'Z' and ch2 == 'Z' then
									finish.y = y3
									finish.x = x3
								else
									-- .FB and BF. are not the same
									-- reading order is always down or right
									local label
									if dir.y == -1 or dir.x == -1 then
										label = ch2 .. ch1
									else
										label = ch1 .. ch2
									end
									pos2label[key(y3, x3)] = label
									-- print(y3, x3, label)
								end
							end
						end
					end
				end
			end
		end
	end
end

-- print(start.y, start.x, finish.y, finish.x)

for k1,v1 in pairs(pos2label) do
	for k2, v2 in pairs(pos2label) do
		if k1 ~= k2 then
			if v1 == v2 then
				portals[k1] = k2
				portals[k2] = k1
			end
		end
	end
end


-- for k,v in pairs(portals) do print(k,v) end

local queue = {
	[1] = {y=start.y, x=start.x, steps=0},
}
local seen = {
	[key(start.y,start.x)] = true,
}

while #queue > 0 do
	local pos = table.remove(queue, 1)
	if pos.y == finish.y and pos.x == finish.x then
		print('Part One', pos.steps)
		break
	end
	local k = key(pos.y, pos.x)
	if portals[k] ~= nil then
		-- assert(portals[key(pos.y,pos.x)]~=key(pos.y,pos.x))
		local ny, nx = unkey(portals[k])
		-- assert(maze[ny][nx]=='.')
		local nk = key(ny,nx)
		if not seen[nk] then
			seen[nk] = true
			table.insert(queue, {y=ny, x=nx, steps=pos.steps+1})
		end
	end
	for _, dir in ipairs(directions) do
		local ny, nx = pos.y + dir.y, pos.x + dir.x
		local nk = key(ny,nx)
		if maze[ny][nx] == '.' then
			if not seen[nk] then
				seen[nk] = true
				table.insert(queue, {y=ny, x=nx, steps=pos.steps+1,level=pos.level})
			end
		end
	end
end

local function portalType(y, x)
	assert(maze[y][x]=='.')
	if maze[y][x-3] == nil or maze[y][x+3] == nil then
		return 'outer'
	end
	if maze[y-3] == nil or maze[y+3] == nil then
		return 'outer'
	end
	-- if maze[y-3][x] == nil or maze[y+3][x] == nil then
	-- 	return 'outer'
	-- end
	return 'inner'
end

local function levelChange(y, x)
	if portalType(y, x) == 'outer' then
		return -1
	else
		return 1
	end
end

queue = {
	[1] = {y=start.y, x=start.x, steps=0, level=0},
}
seen = {
	[key3(start.y,start.x,0)] = true,
}

while #queue > 0 do
	local pos = table.remove(queue, 1)
	-- assert(pos.level>=0)
	if pos.y == finish.y and pos.x == finish.x and pos.level == 0 then
		print('Part Two', pos.steps)
		break
	end
	local k = key(pos.y,pos.x)
	if portals[k] ~= nil then
		if pos.level == 0 and portalType(pos.y, pos.x) == 'outer' then
			-- "when at the outermost level, only the outer labels AA and ZZ function
			-- all other outer labeled tiles are effectively walls."
			-- this rule gets invoked 20 times
		else
		-- assert(portals[key(pos.y,pos.x)]~=key(pos.y,pos.x))
			local ny, nx = unkey(portals[k])
			-- assert(maze[ny][nx]=='.')
			local nlevel = pos.level+levelChange(pos.y,pos.x)
			local nk3 = key3(ny,nx,nlevel)
			if not seen[nk3] then
				seen[nk3] = true
				table.insert(queue, {y=ny, x=nx, steps=pos.steps+1, level=nlevel})
			end
		end
	end
	for _, dir in ipairs(directions) do
		local ny, nx = pos.y + dir.y, pos.x + dir.x
		if maze[ny][nx] == '.' then
			local nk3 = key3(ny,nx,pos.level)
			if not seen[nk3] then
				seen[nk3] = true
				table.insert(queue, {y=ny, x=nx, steps=pos.steps+1, level=pos.level})
			end
		end
	end
end

--[[
$ time luajit 20.lua
Part One	696
Part Two	7538

real	0m0.473s
user	0m0.429s
sys	0m0.044s
]]