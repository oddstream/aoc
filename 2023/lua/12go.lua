-- 12go.lua
-- all credit to https://old.reddit.com/r/adventofcode/comments/18ge41g/2023_day_12_solutions/kd3rclt/

local function split(str, pat)
	local t = {}  -- NOTE: use {n = 0} in Lua-5.0
	local fpat = "(.-)" .. pat
	local last_end = 1
	local s, e, cap = str:find(fpat, 1)
	while s do
	   if s ~= 1 or cap ~= "" then
		  table.insert(t, cap)
	   end
	   last_end = e+1
	   s, e, cap = str:find(fpat, last_end)
	end
	if last_end <= #str then
	   cap = str:sub(last_end)
	   table.insert(t, cap)
	end
	return t
end

---@param ts string[]
---@param fn function
---@return integer[]
local function slicesMap(ts, fn)
	local us = {}
	for _, v in ipairs(ts) do
		us[#us+1] = fn(v)
	end
	return us
end

---@param s string
---@param c integer[]
---@return integer
local function countPossible(s, c)

	---pack three numbers into a single number so it can be used as a map key
	---@param num1 integer
	---@param num2 integer
	---@param num3 integer
	---@return integer
	local function key(num1, num2, num3)
		-- num1 min=0, max=30
		-- num2 min=0, max=15
		-- num3 min=0, max=1 (natch, it's a flag)
		-- return ('%d,%d,%d'):format(one,two,three)
		-- return tostring(one) .. ',' .. tostring(two) .. ',' .. tostring(three)
		return num1 * 32 * 32 + num2 * 32 + num3
	end

	---unpack a single number (a map key) into three numbers
	---@param num integer
	---@return integer
	---@return integer
	---@return integer
	local function unkey(num)
		-- local t = split(num, ',')
		-- return tonumber(t[1]), tonumber(t[2]), tonumber(t[3])
		local num3 = num % 32
		num = (num - num3) / 32
		local num2 = num % 32
		local num1 = (num - num2) / 32
		return num1, num2, num3
	end

	local pos = 0
	-- state is a tuple of 3 values; an index (into c), a count, a flag
	-- KLUDGE ci is zero-based and +1'd before being used as an index
	local cstates = {[key(0,0,0)] = 1}
	for sc in s:gmatch'.' do
		local nstates = {}

		local function incr(k, val)
			local v = nstates[k]
			if v == nil then
				nstates[k] = val
			else
				nstates[k] = v + val
			end
		end

		for state, num in pairs(cstates) do
			-- ci index into array c
			-- cc count
			-- expdot expecting dot flag? always 0 or 1
			local ci, cc, expdot = unkey(state)
			if (sc == '#' or sc == '?') and (ci < #c) and (expdot == 0) then -- 1,2,3 len(3)
				-- we are still looking for broken springs (#)
				if (sc == '?') and (cc == 0) then
					-- we are not in a run of broken springs, so ? can be working
					incr(key(ci, cc, expdot), num)
				end
				cc = cc + 1
				if cc == c[ci+1] then	-- KLUDGE
					-- we've found the full next contiguous section of broken springs
					ci, cc, expdot = ci+1, 0, 1 -- we only want a working spring next
				end
				incr(key(ci, cc, expdot), num)
			elseif (sc == '.' or sc == '?') and (cc == 0) then
				-- we are not in a contiguous run of broken springs
				expdot = 0
				incr(key(ci, cc, expdot), num)
			end
		end
		-- print(sc, next(nstates))
		cstates = nstates
	end
	-- sum states that reached the end of the pattern
	for k, v in pairs(cstates) do
		local one, _, _ = unkey(k)
		if one == #c then
			pos = pos + v
		end
	end
	return pos
end

local part = 2
local input = '12-input.txt'

local paths = 0
for line in io.lines(input) do
	local b, a = line:match'([%?%.%#]+) ([%d,]+)'
	if part == 2 then
		b = string.rep(b, 5, '?')
		a = string.rep(a, 5, ',')
	end
	local c = slicesMap(split(a, ','), tonumber)
	local p = countPossible(b, c)
	paths = paths + p
end
print(paths)

--[[
$ time luajit 12go.lua
11607695322318

real	0m0.254s
user	0m0.253s
sys	0m0.001s
]]