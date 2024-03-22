-- https://adventofcode.com/2019/day/16

local log = require 'log'

if _VERSION ~= 'Lua 5.4' then
	log.error('Lua 5.4 required because of integer division\n')
	return
end

local f = assert(io.open('16-input.txt', "r"))
local input = f:read("*all")
f:close()

local signal = {}
for i = 1, #input do
	signal[#signal+1] = tonumber(input:sub(i,i))
end

-- print(#input)
-- print(#signal)

local basePattern = {0,1,0,-1}

local function coef(i, j)
	local n = (j // i) % 4		-- * / // % have same prec
	-- assert(n>=0 and n<=3)
	return basePattern[n+1]
end

local scratch = {}
for _ = 1, 100 do
	for i = 1, #signal do
		local sum = 0
		for j = 1, #signal do
			-- sum = sum + coef(i,j) * signal[j]
			sum = sum + basePattern[((j // i) % 4)+1] * signal[j]
		end
		scratch[i] = math.abs(sum) % 10
	end
	signal, scratch = scratch, signal

end

log.report('Part One %s\n', table.concat(signal, '', 1, 8))

-- Duplicate the content by 10,000, len will be 6,500,000
local h = string.rep(input, 10000)

-- Extract a substring from position 1 to 7
local i = string.sub(h, 1, 7)
-- print(i)
i = string.sub(h, tonumber(i)+1, -1)
-- print(i)

for _ = 1, 100 do
	local total = 0
	local e = 0
	local len = string.len(i)
	local stringBuilder = {}

	while e < len do
		if e == 0 then
			for digit in string.gmatch(i, "%d") do
				total = total + tonumber(digit)
			end
		elseif e > 0 then
			total = total - tonumber(string.sub(i, e, e))
		end
		stringBuilder[#stringBuilder+1] = tostring(total):sub(-1)	-- get least sig digit
		e = e + 1
	end
	i = table.concat(stringBuilder)
end

log.report('Part Two %s\n', string.sub(i, 1, 8))

--[[
$ time lua54 16.lua
Part One 68764632
Part Two 52825021

real	0m21.056s
user	0m21.033s
sys	0m0.016s
]]

