local inputFilename = "08-input.txt"

local forest = {}
local visible = {}

do
	local row = 1
	for line in io.lines(inputFilename) do
		forest[row] = {}
		visible[row] = {}
		for col = 1, #line do
			table.insert(forest[row], tonumber(line:sub(col,col)))
			table.insert(visible[row], false)
		end
		row = row + 1
	end
end

-- for row, _ in ipairs(forest) do
-- 	for col, _ in ipairs(forest[row]) do
-- 		io.write(forest[row][col])
-- 	end
-- 	io.write('\n')
-- end
if #forest == #forest[1] then
	print('forest is square', #forest)
else
	print('forest is NOT square', #forest == #forest[1])
end

local function look_east()
	for row = 1, #forest do
		-- local seen = {}
		local tallest = -1
		for col = 1, #forest do
			local height = forest[row][col]
			if height > tallest then
				-- table.insert(seen, height)
				visible[row][col] = true
			end
			tallest = math.max(tallest, height)
		end
		-- io.write(string.format('row %d:', row))
		-- for _, v in ipairs(seen) do
		-- 	io.write(string.format(' %d', v))
		-- end
		-- io.write('\n')
	end
end

local function look_west()
	for row = 1, #forest do
		-- local seen = {}
		local tallest = -1
		for col = #forest, 1, -1 do
			local height = forest[row][col]
			if height > tallest then
				-- table.insert(seen, height)
				visible[row][col] = true
			end
			tallest = math.max(tallest, height)
		end
		-- io.write(string.format('row %d:', row))
		-- for _, v in ipairs(seen) do
		-- 	io.write(string.format(' %d', v))
		-- end
		-- io.write('\n')
	end

end

local function look_south()
	for col = 1, #forest do
		-- local seen = {}
		local tallest = -1
		for row = 1, #forest do
			local height = forest[row][col]
			if height > tallest then
				-- table.insert(seen, height)
				visible[row][col] = true
			end
			tallest = math.max(tallest, height)
		end
		-- io.write(string.format('col %d:', col))
		-- for _, v in ipairs(seen) do
		-- 	io.write(string.format(' %d', v))
		-- end
		-- io.write('\n')
	end
end

local function look_north()
	for col = 1, #forest do
		-- local seen = {}
		local tallest = -1
		for row = #forest, 1, -1 do
			local height = forest[row][col]
			if height > tallest then
				-- table.insert(seen, height)
				visible[row][col] = true
			end
			tallest = math.max(tallest, height)
		end
		-- io.write(string.format('col %d:', col))
		-- for _, v in ipairs(seen) do
		-- 	io.write(string.format(' %d', v))
		-- end
		-- io.write('\n')
	end
end

look_east()
look_west()
look_south()
look_north()

local total = 0
for row = 1, #forest do
	for col = 1, #forest do
		if visible[row][col] == true then
			total = total + 1
		end
	end
end
print(total)	-- 21, 1816

local function view_dist_north(row, col)
	local count = 0
	local height = forest[row][col]
	for c = col - 1, 1, -1 do
		count = count + 1
		if forest[row][c] >= height then
			break
		end
	end
	return count
end

local function view_dist_south(row, col)
	local count = 0
	local height = forest[row][col]
	for c = col + 1, #forest do
		count = count + 1
		if forest[row][c] >= height then
			break
		end
	end
	return count
end

local function view_dist_east(row, col)
	local count = 0
	local height = forest[row][col]
	for r = row + 1, #forest do
		count = count + 1
		if forest[r][col] >= height then
			break
		end
	end
	return count
end

local function view_dist_west(row, col)
	local count = 0
	local height = forest[row][col]
	for r = row - 1, 1, -1 do
		count = count + 1
		if forest[r][col] >= height then
			break
		end
	end
	return count
end

-- print(forest[4][3])
-- print(view_dist_north(4, 3))
-- print(view_dist_west(4, 3))
-- print(view_dist_east(4, 3))
-- print(view_dist_south(4, 3))

local max_score = 0
local max_row, max_col = 0, 0
for row = 2, #forest - 1 do
	for col = 2, #forest - 1 do
		local n = view_dist_north(row, col); assert(n~=0)
		local s = view_dist_south(row, col); assert(s~=0)
		local e = view_dist_east(row, col); assert(e~=0)
		local w = view_dist_west(row, col); assert(w~=0)
		local score = n * s * e * w
		if score > max_score then
			max_score = score
			max_row = row
			max_col = col
		end
	end
end

print(max_score, "at", max_row, max_col) -- 383520
