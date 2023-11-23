local inputFilename = "05-input.txt"

-- if not _G.string.trim then
-- 	function _G.string.trim(str)
-- 		return str:gsub("%s+", "")
-- 	end
-- end

local piles = {}

local function display_piles ()
	for k, pile in ipairs(piles) do
		io.write(string.format("%d. ", k))
		for _, value in ipairs(pile) do
			io.write(string.format("%s ", value))
		end
		io.write("\n")
	end
end

local function peek(col)
	return piles[col][#piles[col]]
end

local function pop(col)
	return table.remove(piles[col])
end

local function push(col, c)
	table.insert(piles[col], c)
end

local function get_result()
	local str = ""
	for _, pile in ipairs(piles) do
		str = str .. pile[#pile]
	end
	return str
end

local f, err = io.open(inputFilename, "r")
if f == nil then
	print(err)
	return
end

for line in f:lines("*line") do
	if line == "" then
		break
	end
	-- print(line)
	local col = 1
	local pos = 1
	while pos < #line do
		local tok = line:sub(pos, pos + 3)
		local a, b = tok:find("%u+")
		if a ~= nil then
			tok = tok:sub(a,b)
			-- print(string.format("%d '%s'", col, tok))
			if piles[col] == nil then
				piles[col] = {}
			end
			table.insert(piles[col], 1, tok)
		end

		pos = pos + 4
		col = col + 1
	end
end
-- print("done reading header")
-- display_piles()

for line in f:lines("*line") do
	local num, src, dst = line:match("move (%d*) from (%d*) to (%d*)")
	num = tonumber(num)
	src = tonumber(src)
	dst = tonumber(dst)
	-- print(num, src, dst)
	-- part 1
	-- while num > 0 do
	-- 	local c = pop(src)
	-- 	push(dst, c)
	-- 	num = num - 1
	-- end
	-- part 2
	local ridx = #piles[src] - num + 1
	while num > 0 do
		local c = table.remove(piles[src], ridx)
		table.insert(piles[dst], c)
		num = num - 1
	end
end

f:close()	-- io.close(f)

-- display_piles()

print(get_result())	-- CWMTGHBDW (part 1), SSCGWJCRB (part 2)
