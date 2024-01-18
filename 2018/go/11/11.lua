local input = 2866
local max, best = -math.huge
local power = {}
local lastpower = {}
for x = 1, 300 do
	local rack = x + 10
	power[x] = {}
	lastpower[x] = {}
	for y = 1, 300 do
		power[x][y] = (math.floor((rack*y + input)*rack/100)%10)-5
		lastpower[x][y] = 0
	end
end
for size = 1, 300 do
	for i = 1, 300-size+1 do
		for j = 1, 300-size+1 do
			local sum = lastpower[i][j]
			for x = i, i+size-1 do
				sum = sum + power[x][j+size-1]
			end
			for y = j, j+size-2 do
				sum = sum + power[i+size-1][y]
			end
			lastpower[i][j] = sum
			if sum > max then max, best = sum, {i, j, size} end
		end
	end
	-- print(size, max, table.concat(best, ","))
end
print(table.concat(best, ","))