-- https://adventofcode.com/2023/day/15 Lens Library

local log = require 'log'

---@param s string
---@return integer
local function calcHash(s)
	local hash = 0
	for i = 1, #s do
		local b = s:byte(i)
		hash = hash + b
		hash = hash * 17
		-- hash = math.fmod(hash, 256)
		hash = hash % 256
	end
	return hash
end

---@param list table
---@param label string
---@return integer
local function findLabel(list, label)
	local idx = 0
	for k,v in ipairs(list) do
		if v.lbl == label then
			idx = k
			break
		end
	end
	return idx
end

---@param filename string
---@param expected? integer
---@return integer
local function partOne(filename, expected)
	local result = 0

	local f = assert(io.open(filename, "r"))
	local content = f:read("*all")
	f:close()

	-- remove any trailing \n
	while content:sub(#content, #content) =='\n' do
		content = content:sub(1, #content-1)
	end
	content = content .. ','	-- parsing kludge

	for str in content:gmatch'([^,]+),' do
		result = result + calcHash(str)
	end

	if expected and result ~= expected then
		log.error('part one should be %d\n', expected)
	end
	return result
end

---@param filename string
---@param expected? integer
---@return integer
local function partTwo(filename, expected)
	local result = 0

	local f = assert(io.open(filename, "r"))
	local content = f:read("*all")
	f:close()

	-- remove any trailing \n
	-- while content:sub(#content, #content) =='\n' do
	-- 	content = content:sub(1, #content-1)
	-- end
	-- content = content .. ','	-- parsing kludge

	local boxes = {}
	for _ = 0, 255 do
		table.insert(boxes, {})
	end

	for label, _, focalLength in content:gmatch'(%a+)([=-])(%d*)' do

		-- The result of running the HASH algorithm on the LABEL
		-- indicates the correct box for that step
		local boxNumber = 1 + calcHash(label)

		if focalLength == '' then
			local idx = findLabel(boxes[boxNumber], label)
			if idx ~= 0 then
				table.remove(boxes[boxNumber], idx)
			end
		else
			local idx = findLabel(boxes[boxNumber], label)
			if idx ~= 0 then
				-- lens found in box, 'replace' it
				-- assert(boxes[boxNumber][idx].lbl==label)
				boxes[boxNumber][idx].fl = focalLength
			else
				-- add lens to box
				table.insert(boxes[boxNumber], {lbl=label, fl=focalLength})
			end
		end
		-- print(label, focal_length)
	end

	for ibox,box in ipairs(boxes) do
		for ilens, lens in ipairs(box) do
			local fp = ibox * ilens * lens.fl
			-- print(lens.lbl, ibox, ilens, lens.fl, fp)
			result = result + fp
		end
	end

	if expected and result ~= expected then
		log.error('part two should be %d\n', expected)
	end
	return result
end

log.report('%s\n', _VERSION)
log.report('part one test %d\n', partOne('15-test.txt', 1320))
log.report('part one      %d\n', partOne('15-input.txt', 514025))
log.report('part two test %d\n', partTwo('15-test.txt', 145))
log.report('part two      %d\n', partTwo('15-input.txt', 244461))

--[[
$ time luajit 15.lua
Lua 5.1
part one test 1320
part one      514025
part two test 145
part two      244461

real    0m0.022s
user    0m0.020s
sys     0m0.002s
]]