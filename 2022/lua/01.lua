-- ownerproof-3207314-1699891550-7148ed351f78

local inputFilename = "01-input.txt"

local input, error = io.open(inputFilename, "r")
if input == nil then
	print(error)
	return
end
io.close(input)

local elf = {id=1, total=0}
local elves = {}

for line in io.lines(inputFilename) do
	if line == "" then
		table.insert(elves, elf)
		elf = {id = elf.id + 1, total = 0}
	else
		elf.total = elf.total + tonumber(line)
	end
end
if elf.total > 0 then
	table.insert(elves, elf)
end

io.write(string.format('data for %d elves loaded\n\n', #elves))

table.sort(elves, function(a,b) return a.total > b.total end)

local total = 0
print('rank', 'id', 'total')
for i = 1, 3 do
	print(i, elves[i].id, elves[i].total)
	total = total + elves[i].total
end
io.write(string.format('\ntotal %d\n', total))

--[[
data for 224 elves loaded

num     id      total
1       26      70764
2       179     67568
3       39      65573

total   203905
]]