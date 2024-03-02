local inputFilename = "06-input.txt"

local function unique(s)
	local char_set = {}
	for i = 1, #s do
		local c = s:sub(i,i)
		if char_set[c] then
			return false
		end
		char_set[c] = true
	end
	return true
end

for line in io.lines(inputFilename) do
	for i = 1, #line - 14 do
		local quad = line:sub(i, i+13)
		if unique(quad) then
			print(quad .. " unique " .. tostring(i + 13))
			break
		end
	end
end

-- part 1 1779
-- part 2 2635