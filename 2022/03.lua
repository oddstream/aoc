local inputFilename = "03-input.txt"

local priority = {}

-- Lowercase item types a through z have priorities 1 through 26.
for i = ("a"):byte(1,1), ("z"):byte(1,1) do
	priority[string.char(i)] = i - ("a"):byte(1,1) + 1
end
assert(priority['a']==1)
-- Uppercase item types A through Z have priorities 27 through 52.
for i = ("A"):byte(1,1), ("Z"):byte(1,1) do
	priority[string.char(i)] = i - ("A"):byte(1,1) + 27
end
assert(priority['A']==27)

-- part 1
local sum = 0
for line in io.lines(inputFilename) do
	local first = line:sub(1, #line / 2)
	local second = line:sub(#line / 2 + 1, #line)
	assert(first .. second == line)
	for i = 1, #first do
		local ch = first:sub(i, i)
		if second:find(ch) ~= nil then
			sum = sum + priority[ch]
			break
		end
	end
end

print(sum)	-- 157, 8105

-- part 2
local function common_char(l1, l2, l3)
	for c1 in l1:gmatch"." do
		for c2 in l2:gmatch"." do
			if c1 == c2 then
				for c3 in l3:gmatch"." do
					if c3 == c2 then
						return c3
					end
				end
			end
		end
	end
end

sum = 0
local l = 1
local lines = {}
for line in io.lines(inputFilename) do
	lines[l] = line
	if l == 3 then
		sum = sum + priority[common_char(lines[1], lines[2], lines[3])]
		l = 1
		lines = {}
	else
		l = l + 1
	end
end

print(sum) -- 70, 2363