local inputFilename = "04-input.txt"

local hits1, hits2 = 0, 0
for line in io.lines(inputFilename) do
	local pair = {}
	local n = 1
	for num in line:gmatch("[0-9]+") do
		pair[n] = tonumber(num)
		n = n + 1
	end
	if pair[1] >= pair[3] and pair[2] <= pair[4] then
		hits1 = hits1 + 1
	elseif pair[3] >= pair[1] and pair[4] <= pair[2] then
		hits1 = hits1 + 1
	end

	if pair[2] < pair[3] then
		-- first is before second
	elseif pair[1] > pair[4] then
		-- first is after second
	else
		hits2 = hits2 + 1
	end
end

print(hits1, "hits") -- 2, 444
print(hits2, "hits") -- 4, 801
